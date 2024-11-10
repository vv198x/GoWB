package tasks

import (
	"context"
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/requests"
	"log/slog"
	"math"
	"net/http"
	"time"
)

const minBid = 150
const uriCPM = "https://advert-api.wb.ru/adv/v0/cpm"

func Bidding(ctx context.Context) error {
	if err := UpdateNames(ctx); err != nil {
		return fmt.Errorf("UpdateNames err: %v", err)
	}
	if err := CheckPositions(ctx); err != nil {
		return fmt.Errorf("CheckPositions err: %v", err)
	}

	ids, err := repository.Do().GetAllIds(ctx)
	if err != nil {
		return fmt.Errorf("GetAllIds err: %v", err)
	}
	for _, id := range ids {
		if err = BiddingById(ctx, int64(id)); err != nil {
			slog.Error("BiddingById err: %v", err)
		}
		time.Sleep(time.Duration(config.Get().RetriesTime) * time.Millisecond)
	}

	return err
}

func BiddingById(ctx context.Context, id int64) error {
	bidderInfo, err := repository.Do().GetBidderInfoByAdID(ctx, id)
	if err != nil {
		return fmt.Errorf("GetBidderInfoByAdID err: %v", err)
	}
	if bidderInfo.Paused || bidderInfo.Type == models.TYPE_AUTO || bidderInfo.MaxPosition == 0 {
		return nil
	}

	var nextBid int
	step := config.Get().BidderStep

	//заглушка если не даёт настоящие ставки берём прошлую
	slog.Debug("bidderInfo: ", bidderInfo)
	if bidderInfo.OldCpm < minBid {
		bidderInfo.OldCpm = minBid
	}
	bidderInfo.CurrentBid = bidderInfo.OldCpm

	//Если долеко то можно увеличить шаг
	positionDiff := bidderInfo.Position - bidderInfo.MaxPosition
	if math.Abs(float64(positionDiff)) > 3 {
		step = step * 5
	} else if math.Abs(float64(positionDiff)) > 1 {
		step = step * 3
	}

	//Биддер
	if bidderInfo.Position > bidderInfo.MaxPosition {
		nextBid = bidderInfo.CurrentBid + step
		if nextBid > bidderInfo.MaxBid {
			nextBid = bidderInfo.MaxBid
		}
	} else {
		nextBid = bidderInfo.CurrentBid - step
		if nextBid < minBid {
			nextBid = minBid
		}
	}

	//Если ровное число добавить шаг
	if nextBid%100 == 0 {
		nextBid += config.Get().BidderStep
	}

	//Сечас если не поменялась ставка значит максимальная
	if nextBid == bidderInfo.OldCpm {
		//Нужно если WB показывает всегда правильные ставки
		slog.Debug("nextBid = OldCpm", bidderInfo)
		nextBid += config.Get().BidderStep
	}

	//Получаю айди АРК по айди поиска и синхронизирую ставку
	autoId, err := repository.Do().GetAutoId(ctx, id)
	if err != nil {
		return fmt.Errorf("GetAutoId err: %v", err)
	}
	if autoId > 0 {
		if err2 := SetCPM(ctx, autoId, models.TYPE_AUTO, nextBid, bidderInfo); err2 != nil {
			slog.Error("SetCPM AutoId err: %v", err2)
		}
		time.Sleep(time.Duration(config.Get().RetriesTime) * time.Millisecond)
	}

	if err = SetCPM(ctx, id, models.TYPE_SHEARCH, nextBid, bidderInfo); err != nil {
		slog.Error("SetCPM err: %v", err)
	}
	slog.Debug("bidder AUTO and SHEARCH: ", autoId, id)

	return repository.Do().SaveCpm(ctx, &models.Cpm{
		AdID:        id,
		NewCpm:      nextBid,
		OldCpm:      bidderInfo.CurrentBid,
		OldPosition: bidderInfo.Position,
	})
}

func SetCPM(ctx context.Context, id int64, typeAd int, nextBid int, bidderInfo models.BidderInfo) error {
	reqBody := fmt.Sprintf(`{"advertId": %d, "type": %d, "cpm": %d, "param": %d, "instrument": 6}`,
		id,
		typeAd,
		nextBid,
		bidderInfo.Subject,
	)

	_, err := requests.New(http.MethodPost, uriCPM, []byte(reqBody)).DoWithRetries(ctx)
	if err != nil {
		return fmt.Errorf("request ad status error: %v", err)
	}
	return nil
}

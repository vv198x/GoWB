package tasks

import (
	"context"
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/requests"
	"log/slog"
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
	if nextBid == bidderInfo.OldCpm {
		slog.Error("nextBid == bidderInfo.OldCpm", id)
	}

	reqBody := fmt.Sprintf(`{"advertId": %d, "type": %d, "cpm": %d, "param": %d, "instrument": 6}`,
		id,
		models.TYPE_SHEARCH,
		nextBid,
		bidderInfo.Subject,
	)

	_, err = requests.New(http.MethodPost, uriCPM, []byte(reqBody)).DoWithRetries(ctx)
	if err != nil {
		return fmt.Errorf("request ad status error: %v", err)
	}

	return repository.Do().SaveCpm(ctx, &models.Cpm{
		AdID:   id,
		NewCpm: nextBid,
		OldCpm: bidderInfo.CurrentBid,
	})
}

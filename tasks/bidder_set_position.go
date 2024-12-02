package tasks

import (
	"context"
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/repository"
	"log/slog"
	"time"
)

const factorStep = 3
const factorBid = 65

func ChoiceMaxPosition(ctx context.Context) error {
	ids, err := repository.Do().GetAllIds(ctx)
	if err != nil {
		return fmt.Errorf("GetAllIds err: %v", err)
	}
	for _, id := range ids {
		if err = ChoiceAndSetPositionById(ctx, int64(id)); err != nil {
			slog.Error("ChoiceAndSetPositionById err: %v", err)
		}
		time.Sleep(time.Duration(config.Get().RetriesTime) * time.Millisecond)
	}

	return err
}

func ChoiceAndSetPositionById(ctx context.Context, id int64) error {
	bidderInfo, err := repository.Do().GetBidderInfoByAdID(ctx, id)
	if err != nil {
		return fmt.Errorf("GetBidderInfoByAdID err: %v", err)
	}
	if bidderInfo.Paused || bidderInfo.Type == models.TYPE_AUTO || bidderInfo.MaxPosition <= 2 {
		return nil
	}
	avgBid, err := repository.Do().GetAverageBid(ctx, id, 1)
	if err != nil {
		return fmt.Errorf("GetAverageBid err: %v", err)
	}
	if avgBid == 0 {
		slog.Debug("avgBid == 0")
		return nil
	}

	step := config.Get().BidderStep
	maxPosition := bidderInfo.MaxPosition
	// Отпустить позицию если почти максимальная ставка
	if avgBid > (bidderInfo.MaxBid - step*factorStep) {
		maxPosition += 1
	}

	//Увеличить позицию если средняя ставка меньше 65% (factorBid)
	if avgBid < (bidderInfo.MaxBid * factorBid / 100) {
		maxPosition -= 1
	}

	//Установить если отличется
	if maxPosition != bidderInfo.MaxPosition {
		err = repository.Do().UpdateMaxPosition(ctx, id, maxPosition)
		if err != nil {
			return fmt.Errorf("UpdateMaxPosition err: %w", err)
		}
		slog.Debug("Position updated", "adID", id, "newMaxPosition", maxPosition)
	}
	return err
}

package tasks

import (
	"context"
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/requests"
	"net/http"
	"time"
)

const uriReFill = "https://advert-api.wb.ru/adv/v1/budget/deposit"

func ReFillBalance(ctx context.Context) error {
	adIds, err := repository.Do().GetReFillIds(ctx)
	if err != nil {
		return fmt.Errorf("request refill error: %v", err)
	}
	//запускаю с таймаутом для WB
	for _, id := range adIds {
		if err = reFill(ctx, id); err != nil {
			return fmt.Errorf("ReFill err: %v", err)
		}
		time.Sleep(time.Duration(config.Get().RetriesTime) * time.Millisecond)
	}
	return err
}

func reFill(ctx context.Context, adId int) error {
	finalURL := fmt.Sprintf("%s?id=%d", uriReFill, adId)
	reqBody := fmt.Sprintf(`{"sum": %d,  "type": 1, "return": false}`, config.Get().Amount)

	_, err := requests.New(http.MethodPost, finalURL, []byte(reqBody)).DoWithRetries(ctx)
	if err != nil {
		return fmt.Errorf("request ad status error: %v", err)
	}
	// Записать историю
	if err = repository.Do().AddHistoryAmount(ctx, &models.History{
		AdID:   adId,
		Date:   time.Now(),
		Amount: float64(config.Get().Amount),
	}); err != nil {
		return fmt.Errorf("write db err ", err)
	}

	return err
}

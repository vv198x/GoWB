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

const uriBalance = "https://advert-api.wb.ru/adv/v1/budget"

func UpdateBalance(ctx context.Context) error {
	adIds, err := repository.Do().GetAllIds(ctx)
	if err != nil {
		return fmt.Errorf("db err: %v", err)
	}

	//запускаю с таймаутом для WB
	for _, id := range adIds {
		if err = GetAdBalance(ctx, id); err != nil {
			return fmt.Errorf("getAdBalance err: %v", err)
		}
		time.Sleep(time.Duration(config.Get().RetriesTime) * time.Millisecond)
	}

	return err
}

func GetAdBalance(ctx context.Context, adId int) error {
	var total int
	finalURL := fmt.Sprintf("%s?id=%d", uriBalance, adId)
	data, err := requests.New(http.MethodGet, finalURL, nil).DoWithRetries(ctx)
	if err != nil {
		return fmt.Errorf("request ad status error: %v", err)
	}

	// Не стал создавать модель
	_, err = fmt.Sscanf(string(data), `{"cash":0,"netting":0,"total":%d}`, &total)
	if err != nil {
		return fmt.Errorf("failed to scan JSON string: %v", err)
	}

	if err = repository.Do().SaveOrUpdateBalance(ctx, &models.Balance{
		AdID:    adId,
		Balance: float64(total),
	}); err != nil {
		return fmt.Errorf("write db err ", err, err)
	}
	return err
}

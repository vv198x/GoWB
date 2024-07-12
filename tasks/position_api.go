package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/requests"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

const uriSearch = "https://search.wb.ru/exactmatch/ru/common/v5/search?ab_testing=false&appType=32&curr=rub&dest=-364001&page=1&resultset=catalog&sort=popular&spp=32&suppressSpellcheck=false"

func CheckPositions(ctx context.Context) error {
	requests, err := repository.Do().GetAllRequests(ctx)
	if err != nil {
		return fmt.Errorf("db error getting requests: %v", err)
	}

	for _, request := range requests {

		sleepDuration := time.Duration(config.Get().RetriesTime) * time.Millisecond
		// Удвоил количество попыток
		for i := 0; i < config.Get().Retries*2; i++ {
			time.Sleep(sleepDuration)
			//удвоил время между попытками
			sleepDuration *= 2
			err = GetPosition(ctx, request)
			if err == nil {
				break
			}
			if i == config.Get().Retries*2-1 {
				slog.Debug("maximum number of retries reached", "request", request, "error", err)
			}
		}

	}

	return nil
}

func GetPosition(ctx context.Context, query models.BidderRequest) error {
	finalURL := fmt.Sprintf(`%s&query=%s`, uriSearch, url.QueryEscape(query.Request))

	var wbSearch models.WbSearch

	if data, err := requests.New(http.MethodGet, finalURL, nil).DoUnauthorizedRequest(ctx); err == nil {
		if err = json.Unmarshal(data, &wbSearch); err != nil {
			return fmt.Errorf("json error %V", err)
		}
	} else {
		return fmt.Errorf("request GetPosition error %V", err)
	}

	if len(wbSearch.Data.Products) < 50 {
		return fmt.Errorf("less than 50 products")
	}

	//Записать все продукты под брендом

	for i, product := range wbSearch.Data.Products {
		if product.Brand == config.Get().BrandName {

			if err := repository.Do().SaveOrUpdatePosition(ctx, &models.Position{
				SKU:       int64(product.Id),
				RequestID: query.ID,
				Organic:   product.Log.Position,
				Position:  i + 1,
			}); err != nil {
				slog.Error("db error saving position: %v", err)
			}

		}

	}

	return nil
}

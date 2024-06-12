package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/requests"
	"log/slog"
	"net/http"
	"time"
)

const uriCount = "https://advert-api.wb.ru/adv/v1/promotion/count"
const uriStatus = "https://advert-api.wb.ru/adv/v1/promotion/adverts"

func GetAdList(ctx context.Context) error {
	var listAd models.AdCount
	var advertIds []int
	// Реквест с повтором
	if data, err := requests.New(http.MethodGet, uriCount, nil).DoWithRetries(ctx); err == nil {
		//Байты в json
		if err = json.Unmarshal(data, &listAd); err != nil {
			return fmt.Errorf("json error %V", err)
		}
		slog.Debug("send ad count")
	} else {
		return fmt.Errorf("Request ad count error %V", err)
	}

	//Список рабочих компаний
	for _, advert := range listAd.Adverts {
		if advert.Status == models.AD_PAUSE || advert.Status == models.AD_RUN {
			for _, advertList := range advert.AdvertList {
				advertIds = append(advertIds, advertList.AdvertId)
				//Записываю или обновляю ид, статусы
				if err := repository.Do().SaveOrUpdate(ctx, &models.AdCampaign{
					AdID:   advertList.AdvertId,
					Status: advert.Status,
				}); err != nil {
					return fmt.Errorf("write db err ", err, err)
				}
			}
		}

	}
	if len(advertIds) == 0 {
		return fmt.Errorf("advertIds not found")
	}
	slog.Debug("count ads found: ", len(advertIds))

	return nil
}

func UpdateNames(ctx context.Context) error {
	//Ожидание если база пустая
	time.Sleep(5 * time.Minute)

	adIds, err := repository.Do().GetAllIds(ctx)
	if err != nil {
		return fmt.Errorf("db err: %v", err)
	}
	jsonData, err := json.Marshal(adIds)
	if err != nil {
		return fmt.Errorf("JSON marshal error: %v", err)
	}

	data, err := requests.New(http.MethodPost, uriStatus, jsonData).DoWithRetries(ctx)
	if err != nil {
		return fmt.Errorf("request ad status error: %v", err)
	}

	// Не стал создавать модель
	var mapJ []map[string]interface{}
	if err := json.Unmarshal(data, &mapJ); err != nil {
		return fmt.Errorf("JSON unmarshal error: %v", err)
	}

	for _, item := range mapJ {
		//собираю из мапы данные
		if name, ok := item["name"].(string); ok {
			if id, ok := item["advertId"].(float64); ok {
				if typeAd, ok := item["type"].(float64); ok {
					//обновляю в базе
					if err := repository.Do().SaveOrUpdate(ctx, &models.AdCampaign{
						AdID: int(id),
						Type: int(typeAd),
						Name: name,
					}); err != nil {
						return fmt.Errorf("db error saving or updating campaign: %v", err)
					}
				}
			}
		}
	}

	return nil
}

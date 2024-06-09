package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/requests"
	"log/slog"
	"net/http"
)

const uriCount = "https://advert-api.wb.ru/adv/v1/promotion/count"
const uriStatus = "https://advert-api.wb.ru/adv/v1/promotion/adverts"

func GetAdList() error {
	var listAd models.AdCount
	var advertIds []int
	// Реквест с повтором
	if data, err := requests.New(http.MethodGet, uriCount, nil).DoWithRetries(); err == nil {
		//Байты в json
		if err = json.Unmarshal(data, &listAd); err != nil {
			return fmt.Errorf("json error %V", err)
		}
		slog.Debug("send ad count")
	} else {
		return fmt.Errorf("Request ad count error %V", err)
	}

	// Создать список рабочих компаний пока в массив интов
	for _, advert := range listAd.Adverts {
		if advert.Status == models.AD_PAUSE || advert.Status == models.AD_RUN {
			for _, advertList := range advert.AdvertList {
				advertIds = append(advertIds, advertList.AdvertId)
			}
		}

	}
	if len(advertIds) == 0 {
		return fmt.Errorf("advertIds not found")
	}
	slog.Debug("count ads found: ", len(advertIds))

	return nil
}

func GetAdStatus(adIds []int) error {
	jsonData, err := json.Marshal(adIds)
	if err != nil {
		return fmt.Errorf("JSON marshal error: %v", err)
	}

	data, err := requests.New(http.MethodPost, uriStatus, jsonData).DoWithRetries()
	if err != nil {
		return fmt.Errorf("request ad status error: %v", err)
	}

	// Не стал создавать модель
	var mapJ []map[string]interface{}
	if err := json.Unmarshal(data, &mapJ); err != nil {
		return fmt.Errorf("JSON unmarshal error: %v", err)
	}

	for _, item := range mapJ {
		if name, ok := item["name"].(string); ok {
			fmt.Println("Name:", name)
		} else {
			fmt.Println("Error: name is not a string")
		}
	}

	return nil
}

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

func GetAdStatus() error {
	var listAd models.AdCount
	var advertIds []int
	// Реквест с повтором
	if data, err := requests.New(http.MethodGet, uriCount, nil).DoWithRetries(); err == nil {
		//Байты в json
		if err = json.Unmarshal(data, &listAd); err != nil {
			slog.Error("json error", err)
		}
		slog.Debug("send ad count")
	} else {
		slog.Error("Request ad count error", err)
	}
	// Создать список рабочих компаний пока в массив интов
	for _, advert := range listAd.Adverts {
		if advert.Status == models.AD_PAUSE || advert.Status == models.AD_RUN {
			for _, advertList := range advert.AdvertList {
				advertIds = append(advertIds, advertList.AdvertId)
			}
		}

	}
	fmt.Println(advertIds)
	return nil
}

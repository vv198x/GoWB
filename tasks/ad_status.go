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
		slog.Debug("send ad count", listAd)
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
					return fmt.Errorf("write db err ", err)
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

	var campaigns []models.WBCampaign
	if err := json.Unmarshal(data, &campaigns); err != nil {
		return fmt.Errorf("JSON unmarshal error: %v", err)
	}

	for _, campaign := range campaigns {
		var sku int64
		var cpm, subject int
		switch {
		case campaign.Type == models.TYPE_AUTO:
			if len(campaign.AutoParams.Nms) > 0 {
				sku = campaign.AutoParams.Nms[0]
			}
			if len(campaign.AutoParams.NmCPM) > 0 {
				cpm = campaign.AutoParams.NmCPM[0].Cpm
			}
			subject = campaign.AutoParams.Subject.Id

		case campaign.Type == models.TYPE_SHEARCH:
			if len(campaign.UnitedParams) > 0 {
				if len(campaign.UnitedParams[0].Nms) > 0 {
					sku = campaign.UnitedParams[0].Nms[0]
				}
				cpm = campaign.UnitedParams[0].SearchCPM
				subject = campaign.UnitedParams[0].Subject.Id

			}
		}

		if err := repository.Do().SaveOrUpdate(ctx, &models.AdCampaign{
			AdID:       campaign.AdvertId,
			Type:       campaign.Type,
			Name:       campaign.Name,
			SKU:        sku,
			CurrentBid: cpm,
			Subject:    subject,
		}); err != nil {
			return fmt.Errorf("db error saving or updating campaign: %v", err)
		}
	}

	return nil
}

package tasks

import (
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/models"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/requests"
	"net/http"
	"time"
)

const uriReFill = "https://advert-api.wb.ru/adv/v1/budget/deposit"

func ReFillBalance() error {
	adIds, err := repository.Do().GetReFillIds()
	if err != nil {
		return fmt.Errorf("request refill error: %v", err)
	}
	//запускаю с таймаутом для WB
	for _, id := range adIds {
		if err = reFill(id); err != nil {
			return fmt.Errorf("ReFill err: %v", err)
		}
		time.Sleep(time.Duration(config.Get().RetriesTime) * time.Millisecond)
	}
	return err
}

func reFill(adId int) error {
	finalURL := fmt.Sprintf("%s?id=%d", uriReFill, adId)
	reqBody := fmt.Sprintf(`{"sum": %d,  "type": 1, "return": false}`, config.Get().Amount)

	_, err := requests.New(http.MethodPost, finalURL, []byte(reqBody)).DoWithRetries()
	if err != nil {
		return fmt.Errorf("request ad status error: %v", err)
	}
	// Записать историю
	if err = repository.Do().AddHistoryAmount(&models.History{
		AdID:   adId,
		Date:   time.Now(),
		Amount: float64(config.Get().Amount),
	}); err != nil {
		return fmt.Errorf("write db err ", err, err)
	}
	fmt.Println(finalURL, reqBody)

	return err
}

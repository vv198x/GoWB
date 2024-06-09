package tasks

import (
	"fmt"
	"github.com/vv198x/GoWB/requests"
	"log/slog"
	"net/http"
)

const uriBalance = "https://advert-api.wb.ru/adv/v1/budget"

func GetAdBalance(adId int) error {
	var total int
	finalURL := fmt.Sprintf("%s?id=%d", uriBalance, adId)
	data, err := requests.New(http.MethodGet, finalURL, nil).DoWithRetries()
	if err != nil {
		return fmt.Errorf("request ad status error: %v", err)
	}

	// Не стал создавать модель
	_, err = fmt.Sscanf(string(data), `{"cash":0,"netting":0,"total":%d}`, &total)
	if err != nil {
		return fmt.Errorf("failed to scan JSON string: %v", err)
	}
	slog.Debug("id balance: ", adId, total)
	return nil
}

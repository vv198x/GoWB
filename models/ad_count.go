package models

import "time"

// Для ответа вб
type AdCount struct {
	Adverts []struct {
		Type       int `json:"type"`
		Status     int `json:"status"`
		Count      int `json:"count"`
		AdvertList []struct {
			AdvertId   int       `json:"advertId"`
			ChangeTime time.Time `json:"changeTime"`
		} `json:"advert_list"`
	} `json:"adverts"`
}

package models

import "time"

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

const (
	AD_PAUSE = 11
	AD_RUN   = 9
)

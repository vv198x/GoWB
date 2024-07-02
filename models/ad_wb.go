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

type WBCampaign struct {
	EndTime          time.Time `json:"endTime"`
	CreateTime       time.Time `json:"createTime"`
	ChangeTime       time.Time `json:"changeTime"`
	StartTime        time.Time `json:"startTime"`
	SearchPluseState bool      `json:"searchPluseState"`
	AutoParams       struct {
		Cpm     int `json:"cpm"`
		Subject struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"subject"`
		Sets []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"sets"`
		Nms    []int64 `json:"nms"`
		Active struct {
			Carousel bool `json:"carousel"`
			Recom    bool `json:"recom"`
			Booster  bool `json:"booster"`
		} `json:"active"`
		NmCPM []struct {
			Nm  int `json:"nm"`
			Cpm int `json:"cpm"`
		} `json:"nmCPM"`
	} `json:"autoParams"`
	UnitedParams []struct {
		CatalogCPM int `json:"catalogCPM"`
		SearchCPM  int `json:"searchCPM"`
		Subject    struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"subject"`
		Menus []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"menus"`
		Nms []int64 `json:"nms"`
	} `json:"unitedParams"`
	Name        string `json:"name"`
	DailyBudget int    `json:"dailyBudget"`
	AdvertId    int    `json:"advertId"`
	Status      int    `json:"status"`
	Type        int    `json:"type"`
	PaymentType string `json:"paymentType"`
}

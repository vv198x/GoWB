package repository

import "github.com/vv198x/GoWB/models"

type AdCampaign interface {
	SaveOrUpdate(campaign *models.AdCampaign) error
	GetAllIds() ([]int, error)
	SaveOrUpdateBalance(Balance *models.Balance) error
}

var R AdCampaign

func Do() AdCampaign {
	return R
}

package repository

import (
	"context"
	"github.com/vv198x/GoWB/models"
)

type AdCampaign interface {
	SaveOrUpdate(context.Context, *models.AdCampaign) error
	GetAllIds(context.Context) ([]int, error)
	SaveOrUpdateBalance(context.Context, *models.Balance) error
	GetReFillIds(context.Context) ([]int, error)
	AddHistoryAmount(context.Context, *models.History) error
}

var R AdCampaign

func Do() AdCampaign {
	return R
}

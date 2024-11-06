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
	GetAllAds(ctx context.Context) ([]models.AdCampaign, error)
	GetAllRequests(ctx context.Context) ([]models.BidderRequest, error)
	SaveOrUpdatePosition(ctx context.Context, position *models.Position) error
	GetBidderInfoByAdID(ctx context.Context, adID int64) (models.BidderInfo, error)
	SaveCpm(ctx context.Context, cpm *models.Cpm) error
	GetAutoId(ctx context.Context, adID int64) (int64, error)
}

var R AdCampaign

func Do() AdCampaign {
	return R
}

package repository

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/vv198x/GoWB/models"
	"time"
)

func (repo *AdCampaignRepository) GetAllRequests(ctx context.Context) ([]models.BidderRequest, error) {
	var requests []models.BidderRequest
	err := repo.DB.Model(&requests).
		Context(ctx).
		Select()
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func (repo *AdCampaignRepository) SaveOrUpdatePosition(ctx context.Context, position *models.Position) error {
	existingPosition := &models.Position{}
	err := repo.DB.Model(existingPosition).
		Context(ctx).
		Where("sku = ? AND request_id = ?", position.SKU, position.RequestID).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return err
	}

	if err == pg.ErrNoRows {
		// Создать новую запись
		_, err := repo.DB.Model(position).
			Context(ctx).
			Insert()
		return err
	}

	position.UpdatedAt = time.Now()

	// Обновить только непустые столбцы
	columns := []string{"updated_at"}
	if position.Organic != 0 {
		columns = append(columns, "organic")
	}
	if position.Position != 0 {
		columns = append(columns, "position")
	}

	_, err = repo.DB.Model(position).
		Context(ctx).
		Column(columns...).
		Where("sku = ? AND request_id = ?", position.SKU, position.RequestID).
		Update()
	return err
}

func (repo *AdCampaignRepository) GetBidderInfoByAdID(ctx context.Context, adID int64) (models.BidderInfo, error) {
	var bidderInfo models.BidderInfo

	subquery := repo.DB.Model((*models.Cpm)(nil)).
		Context(ctx).
		ColumnExpr("old_cpm").
		Where("ad_id = ?", adID).
		Order("created_at DESC").
		Limit(1)

	if err := repo.DB.Model((*models.BidderList)(nil)).
		Context(ctx).
		ColumnExpr("bidder_list.request_id, ad_campaign.current_bid, bidder_list.max_bid, bidder_list.paused, position.position, (?) AS old_cpm", subquery).
		Join("JOIN positions AS position ON position.request_id = bidder_list.request_id").
		Where("bidder_list.ad_id = ?", adID).
		Select(&bidderInfo); err != nil {
		return bidderInfo, err
	}

	return bidderInfo, nil
}

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

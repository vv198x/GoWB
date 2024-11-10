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
	query := `
SELECT
    ac.current_bid,
    ac.type,
    ac.subject,
    bl.max_bid,
    bl.max_position,
    bl.paused,
    p.position,
    (
        SELECT
            c.new_cpm
        FROM
            cpms c
        WHERE
            c.ad_id = bl.ad_id
        ORDER BY
            c.created_at DESC
        LIMIT 1
    ) AS old_cpm
FROM
    bidder_lists bl
        JOIN
    ad_campaigns ac ON ac.ad_id = bl.ad_id
        JOIN
    positions p ON p.request_id = bl.request_id AND p.sku = ac.sku
WHERE
    bl.ad_id = ?;
    `
	_, err := repo.DB.QueryContext(ctx, &bidderInfo, query, adID)

	return bidderInfo, err
}

func (repo *AdCampaignRepository) SaveCpm(ctx context.Context, cpm *models.Cpm) error {
	_, err := repo.DB.Model(cpm).
		Context(ctx).
		Insert()
	return err
}

func (repo *AdCampaignRepository) GetAutoId(ctx context.Context, adID int64) (int64, error) {
	var id int64
	subQuery := repo.DB.Model((*models.AdCampaign)(nil)).
		Column("sku").
		Where("ad_id = ?", adID)

	err := repo.DB.Model((*models.AdCampaign)(nil)).
		Context(ctx).
		Column("ad_id").
		Where("type = ?", models.TYPE_AUTO).
		Where("sku = (?)", subQuery).
		Select(&id)

	if err != nil {
		if err == pg.ErrNoRows {
			return 0, nil // Возвращаем 0, если запись не найдена
		}
		return 0, err
	}
	return id, nil
}
func (repo *AdCampaignRepository) GetAverageBid(ctx context.Context, adID int64, hours int) (int, error) {
	var averageBid float64

	err := repo.DB.Model((*models.Cpm)(nil)).
		Context(ctx).
		ColumnExpr("COALESCE(AVG(new_cpm), 0) AS avg_bid").
		Where("ad_id = ?", adID).
		Where("created_at > NOW() - INTERVAL '? hours'", hours).
		Select(&averageBid)

	if err != nil {
		return 0, err
	}

	return int(averageBid), nil
}

func (repo *AdCampaignRepository) UpdateMaxPosition(ctx context.Context, adID int64, newMaxPosition int) error {
	_, err := repo.DB.Model(&models.BidderList{}).
		Context(ctx).
		Set("max_position = ?", newMaxPosition).
		Set("updated_at = now()").
		Where("ad_id = ?", adID).
		Update()

	return err
}

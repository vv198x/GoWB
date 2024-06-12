package repository

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/vv198x/GoWB/models"
	"time"
)

type AdCampaignRepository struct {
	DB *pg.DB
}

func (repo *AdCampaignRepository) SaveOrUpdate(ctx context.Context, campaign *models.AdCampaign) error {
	existingCampaign := &models.AdCampaign{}
	err := repo.DB.Model(existingCampaign).
		Context(ctx).
		Where("ad_id = ?", campaign.AdID).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return err
	}

	if err == pg.ErrNoRows {
		// Рекламная кампания не найдена, создаем новую запись
		_, err := repo.DB.Model(campaign).
			Context(ctx).
			Insert()
		return err
	}

	// Рекламная кампания найдена, обновляем запись
	campaign.CreatedAt = existingCampaign.CreatedAt // Не изменяем дату создания
	campaign.UpdatedAt = time.Now()                 // Обновляем дату обновления

	//не обновляем если пустые
	columns := []string{"updated_at"}
	if campaign.Name != "" {
		columns = append(columns, "name")
	}
	// budget не обновляем
	if campaign.Type != 0 {
		columns = append(columns, "type")
	}
	if campaign.Status != 0 {
		columns = append(columns, "status")
	}

	_, err = repo.DB.Model(campaign).
		Context(ctx).
		Column(columns...).
		Where("ad_id = ?", campaign.AdID).
		Update()
	return err
}

func (repo *AdCampaignRepository) SaveOrUpdateBalance(ctx context.Context, balance *models.Balance) error {
	// Устанавливаем текущее время для updated_at перед выполнением запроса
	balance.UpdatedAt = time.Now()
	_, err := repo.DB.Model(balance).
		Context(ctx).
		OnConflict("(ad_id) DO UPDATE").
		Set("balance = EXCLUDED.balance, updated_at = EXCLUDED.updated_at").
		Insert()
	return err
}

func (repo *AdCampaignRepository) GetAllIds(ctx context.Context) ([]int, error) {
	var ids []int
	err := repo.DB.Model(&models.AdCampaign{}).
		Context(ctx).
		Column("ad_id").
		Select(&ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// Оставил тут логику
// Меньше 500, дневной бюджет не привышен, не помечен "непополнять"
func (repo *AdCampaignRepository) GetReFillIds(ctx context.Context) ([]int, error) {
	var ids []int
	query := `
        SELECT 
            ac.ad_id
        FROM 
            ad_campaigns ac
        LEFT JOIN 
            balances b ON ac.ad_id = b.ad_id
        WHERE 
            COALESCE(b.balance, 0) < 500
            AND ac.do_not_refill = FALSE
            AND ac.budget > (
                SELECT 
                    COALESCE(SUM(amount), 0)
                FROM 
                    histories
                WHERE 
                    ad_id = ac.ad_id
                    AND date = CURRENT_DATE
            );
    `
	_, err := repo.DB.QueryContext(ctx, &ids, query)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (repo *AdCampaignRepository) AddHistoryAmount(ctx context.Context, entry *models.History) error {
	_, err := repo.DB.Model(entry).
		Context(ctx).
		Insert()
	return err
}

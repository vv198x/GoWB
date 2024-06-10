package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/vv198x/GoWB/models"
	"time"
)

type AdCampaignRepository struct {
	DB *pg.DB
}

func (repo *AdCampaignRepository) SaveOrUpdate(campaign *models.AdCampaign) error {
	existingCampaign := &models.AdCampaign{}
	err := repo.DB.Model(existingCampaign).
		Where("ad_id = ?", campaign.AdID).
		Select()
	if err != nil && err != pg.ErrNoRows {
		return err
	}

	if err == pg.ErrNoRows {
		// Рекламная кампания не найдена, создаем новую запись
		_, err := repo.DB.Model(campaign).Insert()
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
		Column(columns...).
		Where("ad_id = ?", campaign.AdID).
		Update()
	return err
}

func (repo *AdCampaignRepository) SaveOrUpdateBalance(balance *models.Balance) error {
	// Устанавливаем текущее время для updated_at перед выполнением запроса
	balance.UpdatedAt = time.Now()
	_, err := repo.DB.Model(balance).
		OnConflict("(ad_id) DO UPDATE").
		Set("balance = EXCLUDED.balance, updated_at = EXCLUDED.updated_at").
		Insert()
	return err
}

func (repo *AdCampaignRepository) GetAllIds() ([]int, error) {
	var ids []int
	err := repo.DB.Model(&models.AdCampaign{}).
		Column("ad_id").
		Select(&ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// Оставил тут логику
// Меньше 500, дневной бюджет не привышен, не помечен "непополнять"
func (repo *AdCampaignRepository) GetReFillIds() ([]int, error) {
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
	_, err := repo.DB.Query(&ids, query)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (repo *AdCampaignRepository) AddHistoryAmount(entry *models.History) error {
	_, err := repo.DB.Model(entry).Insert()
	return err
}
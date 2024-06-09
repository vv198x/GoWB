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
	err := repo.DB.Model(existingCampaign).Where("ad_id = ?", campaign.AdID).Select()
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

	_, err = repo.DB.Model(campaign).Where("ad_id = ?", campaign.AdID).Update()
	return err
}

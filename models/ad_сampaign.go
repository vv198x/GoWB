package models

import (
	"time"
)

type AdCampaign struct {
	AdID        int       `pg:"ad_id,pk"`
	Name        string    `pg:"name"`
	Budget      float64   `pg:"budget"`
	Status      int       `pg:"status"`
	Type        int       `pg:"type"`
	DoNotRefill bool      `pg:"do_not_refill"`
	SKU         int64     `pg:"sku"`
	CurrentBid  int       `pg:"current_bid"`
	CreatedAt   time.Time `pg:"created_at,default:now()"`
	UpdatedAt   time.Time `pg:"updated_at"`
}

// отдельно от таблицы
type Balance struct {
	AdID      int       `pg:"ad_id,pk"`
	Balance   float64   `pg:"balance"`
	UpdatedAt time.Time `pg:"updated_at"`
}

// status
const (
	AD_PAUSE     = 11
	AD_RUN       = 9
	TYPE_SHEARCH = 9
	TYPE_AUTO    = 8
)

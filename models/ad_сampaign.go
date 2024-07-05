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
	Subject     int       `pg:"subject"`
	CreatedAt   time.Time `pg:"created_at,default:now()"`
	UpdatedAt   time.Time `pg:"updated_at"`
}

type Balance struct {
	AdID      int       `pg:"ad_id,pk"`
	Balance   float64   `pg:"balance"`
	UpdatedAt time.Time `pg:"updated_at"`
}

const (
	AD_PAUSE     = 11
	AD_RUN       = 9
	TYPE_SHEARCH = 9
	TYPE_AUTO    = 8
)

type History struct {
	ID        int       `pg:",pk"`
	AdID      int       `pg:",notnull"`
	Date      time.Time `pg:",notnull"`
	Amount    float64   `pg:",notnull,default:0"`
	CreatedAt time.Time `pg:",default:now()"`
}

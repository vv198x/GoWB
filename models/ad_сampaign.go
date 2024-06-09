package models

import (
	"time"
)

type AdCampaign struct {
	AdID      int       `pg:"ad_id,pk"`
	Name      string    `pg:"name"`
	Budget    float64   `pg:"budget"`
	Status    int       `pg:"status"`
	Type      int       `pg:"type"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at"`
}

// status
const (
	AD_PAUSE = 11
	AD_RUN   = 9
)

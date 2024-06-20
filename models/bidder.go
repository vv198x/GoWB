package models

import "time"

type BidderRequest struct {
	ID      int    `pg:"id,pk"`
	Request string `pg:"request"`
}

type BidderList struct {
	AdID       int64     `pg:"ad_id,pk"`
	RequestID  int       `pg:"request_id,pk"`
	CurrentBid int       `pg:"current_bid"`
	MaxBid     int       `pg:"max_bid"`
	Paused     bool      `pg:"paused,default:false"`
	CreatedAt  time.Time `pg:"created_at,default:now()"`
	UpdatedAt  time.Time `pg:"updated_at"`
}

type Position struct {
	SKU       int64     `pg:"sku,pk"`
	RequestID int       `pg:"request_id,pk"`
	Organic   int       `pg:"organic"`
	Position  int       `pg:"position"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}

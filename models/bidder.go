package models

import "time"

type BidderRequest struct {
	ID      int    `pg:"id,pk"`
	Request string `pg:"request"`
}

type BidderList struct {
	AdID      int64     `pg:"ad_id,pk"`
	RequestID int       `pg:"request_id,pk"`
	MaxBid    int       `pg:"max_bid"`
	Paused    bool      `pg:"paused,default:false"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}

type Position struct {
	SKU       int64     `pg:"sku"`
	RequestID int       `pg:"request_id"`
	Organic   int       `pg:"organic"`
	Position  int       `pg:"position"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}

type Cpm struct {
	ID        int       `pg:"id,pk"`
	AdID      int64     `pg:"ad_id"`
	OldCpm    int       `pg:"old_cpm"`
	NewCpm    int       `pg:"new_cpm"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
}

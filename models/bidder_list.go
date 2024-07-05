package models

type BidderInfo struct {
	CurrentBid  int  `pg:"current_bid"`
	MaxBid      int  `pg:"max_bid"`
	MaxPosition int  `pg:"max_position"`
	Paused      bool `pg:"paused"`
	Position    int  `pg:"position"`
	OldCpm      int  `pg:"old_cpm"`
}

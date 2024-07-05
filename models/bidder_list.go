package models

type BidderInfo struct {
	CurrentBid  int  `pg:"current_bid"`
	Type        int  `pg:"type"`
	Subject     int  `pg:"subject"`
	MaxBid      int  `pg:"max_bid"`
	MaxPosition int  `pg:"max_position"`
	Paused      bool `pg:"paused"`
	Position    int  `pg:"position"`
	OldCpm      int  `pg:"old_cpm"`
}

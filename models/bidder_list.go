package models

type BidderInfo struct {
	CurrentBid  int
	MaxBid      int
	MaxPosition int
	Paused      bool
	Position    int
	OldCpm      int
}

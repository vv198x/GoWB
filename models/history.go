package models

import "time"

type History struct {
	ID        int       `pg:",pk"`
	AdID      int       `pg:",notnull"`
	Date      time.Time `pg:",notnull"`
	Amount    float64   `pg:",notnull,default:0"`
	CreatedAt time.Time `pg:",default:now()"`
}

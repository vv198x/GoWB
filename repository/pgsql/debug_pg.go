package pgsql

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/vv198x/GoWB/config"
	"log/slog"
	"strings"
)

var DB *pg.DB

type DBLogger struct{}

func (d DBLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d DBLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	b, _ := q.FormattedQuery()
	slog.Debug(string(b))
	return nil
}

func DebugPG() {
	if strings.ToLower(config.Get().LogLevel) == "debug" {
		DB.AddQueryHook(DBLogger{})
	}
}

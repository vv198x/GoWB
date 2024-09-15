package main

import (
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/logger"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/repository/pgsql"
	migrations "github.com/vv198x/GoWB/repository/pgsql/migration"
	"github.com/vv198x/GoWB/scheduler"
	"github.com/vv198x/GoWB/tasks"
	"log/slog"
	"time"
)

func main() {
	slog.SetDefault(logger.NewLogger(config.Get().LogLevel))

	//коннект pg c TSL и повтором
	if err := pgsql.ConnPG(); err != nil {
		slog.Error("Dont connect pgsql")
		panic("Dont connect pgsql")
	}
	//логирование запросов если логлевел дебаг
	//pgsql.DebugPG()
	//миграция
	migrations.Start()
	//инициализация репозитория
	repository.R = &repository.AdCampaignRepository{DB: pgsql.DB}

	//планировщик
	go scheduler.Add(tasks.AutoReFill, 20*time.Minute)
	go scheduler.Add(tasks.Bidding, 3*time.Minute)

	select {}
}

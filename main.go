package main

import (
	"context"
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/logger"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/repository/pgsql"
	migrations "github.com/vv198x/GoWB/repository/pgsql/migration"
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
	pgsql.DebugPG()
	//миграция
	migrations.Start()
	//инициализация репозитория
	repository.R = &repository.AdCampaignRepository{DB: pgsql.DB}

	//go scheduler.Add(tasks.AutoReFill, 1*time.Hour)
	//go scheduler.Add(tasks.UpdateNames, 24*time.Hour)

	//select {}
	for i := 0; i < 20; i++ {
		fmt.Println(tasks.GetPosition(context.Background(),
			"l карнитин жидкий"))
		time.Sleep(500 * time.Millisecond)
	}

}

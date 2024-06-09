package main

import (
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/logger"
	"github.com/vv198x/GoWB/repository"
	"github.com/vv198x/GoWB/repository/pgsql"
	migrations "github.com/vv198x/GoWB/repository/pgsql/migration"
	"github.com/vv198x/GoWB/tasks"
	"log/slog"
)

func main() {
	slog.SetDefault(logger.NewLogger(config.Get().LogLevel))

	//коннект pg c TSL и повтором
	if err := pgsql.ConnPG(); err != nil {
		slog.Error("Dont connect pgsql")
		panic("Dont connect pgsql")
	}
	//Логирование если логлевел дебаг
	pgsql.DebugPG()
	//Миграция
	migrations.Start()
	//Инициализация репозитория
	repository.R = &repository.AdCampaignRepository{DB: pgsql.DB}

	//записать все id и статусы
	fmt.Println(tasks.GetAdList())
	//обновить имена и тип
	//fmt.Println(tasks.UpdateNames())
	//обновить бюджеты
	fmt.Println(tasks.UpdateBalance())

	//TODO вынести баланс в отдельную таблицу

}

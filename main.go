package main

import (
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/logger"
	"github.com/vv198x/GoWB/repository/pgsql"
	migrations "github.com/vv198x/GoWB/repository/pgsql/migration"
	"log/slog"
)

func main() {
	slog.SetDefault(logger.NewLogger(config.Get().LogLevel))

	//коннект TSL и повтором
	if err := pgsql.ConnPG(); err != nil {
		slog.Error("Dont connect pgsql")
		panic("Dont connect pgsql")
	}
	//Логирование если логлевел дебаг
	pgsql.DebugPG()
	//Миграция
	migrations.Start()

}

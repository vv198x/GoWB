package migrations

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/repository/pgsql"
	"log/slog"
)

/*
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - set_version [version] - sets db version without running migrations.
*/

func Start() {
	migrationLevel := config.Get().Migration
	if migrationLevel == "init" || migrationLevel == "up" {
		var found bool
		_, err := pgsql.DB.Query(pg.Scan(&found), `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'gopg_migrations')`)
		if err != nil {
			slog.Error("migration err", err)
		}

		// если не найдена таблица gopg_migrations
		if !found {
			migrate("init")
		}

		// для первого раза и по дефолту
		migrate("up")
	} else {
		if migrationLevel != "" {
			migrate(migrationLevel)
		}

	}
}

func migrate(t string) {
	oldVersion, newVersion, err := migrations.Run(pgsql.DB, t)
	if err != nil {
		slog.Error("migration err", err)
	}
	slog.Info(fmt.Sprintf("Migrated from version %d to %d\n", oldVersion, newVersion))
}

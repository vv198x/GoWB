package pgsql

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/vv198x/GoWB/config"
	"log/slog"
	"math/rand"
	"time"
)

const retries = 3
const delay = 15 * time.Second
const maxRandomDelay = 45 * time.Second

func ConnPG() error {
	opt := pg.Options{
		Addr:        config.Get().AddrPg,
		Database:    config.Get().DbPg,
		User:        config.Get().UserPg,
		Password:    config.Get().PassPg,
		PoolSize:    40,
		PoolTimeout: time.Minute,
	}

	// TLS pg from internet for me
	if config.IsHomeHost() {
		opt.TLSConfig = TlsConn()
		slog.Info("Connect from TLS")
	}

	//Коннект с повтором
	return connRetry(opt)
}

func connRetry(opt pg.Options) error {
	for i := 0; i < retries; i++ {
		DB = pg.Connect(&opt)

		// ping pg
		if _, err := DB.Exec("SELECT 2+2"); err == nil {
			return nil

			// retry
		} else {
			random := rand.New(rand.NewSource(time.Now().UnixNano()))
			sleepTime := delay + time.Duration(random.Int63n(int64(maxRandomDelay)))
			slog.Debug(fmt.Sprintf(
				"Failed to connect to PG after trying %d times, address: %s, error: %s",
				i+1, opt.Addr, err))
			time.Sleep(sleepTime)
			return err
		}
	}

	return fmt.Errorf("Failed to connect to PG after retries")
}

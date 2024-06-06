package main

import (
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/logger"
	"log/slog"
)

func main() {
	slog.SetDefault(logger.NewLogger(config.Get().Logger))

}

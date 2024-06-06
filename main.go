package main

import (
	"github.com/vv198x/GoWB/logger"
	"log/slog"
)

func main() {
	slog.SetDefault(logger.NewLogger(slog.LevelDebug))

	slog.Error("1")

}

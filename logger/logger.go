package logger

import (
	"github.com/vv198x/GoWB/config"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const logExt = ".log"

func NewLogger(logLevel string) *slog.Logger {
	var level slog.Level
	if err := os.MkdirAll(config.Get().LogDir, 0755); err != nil {
		slog.Error("create log directory err: %v", err)
		return nil
	}

	file, err := os.OpenFile(getLogFilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		slog.Error("write log err: %v", err)
		return nil
	}

	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case "info":
		level = slog.LevelInfo
	case "debug":
		level = slog.LevelDebug

	}

	return slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	}))
}

func getLogFilePath() string {
	return filepath.Join(
		config.Get().LogDir,
		time.Now().Format("06.01.02")+logExt,
	)
}

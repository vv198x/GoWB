package logger

import (
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

const logDir = "log"
const logExt = ".log"

func NewLogger(logLevel slog.Level) *slog.Logger {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		slog.Error("create log directory err: %v", err)
		return nil
	}

	file, err := os.OpenFile(getLogFilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		slog.Error("write log err: %v", err)
		return nil
	}

	return slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}))
}

func getLogFilePath() string {
	return filepath.Join(
		logDir,
		time.Now().Format("06.01.02")+logExt,
	)
}

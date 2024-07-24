package scheduler

import (
	"context"
	"github.com/vv198x/GoWB/config"
	"log/slog"
	"time"
)

type Task func(ctx context.Context) error

// 20 минут тайм аут контекста, 5 минут между повторами
func Add(task Task, interval time.Duration) {
	slog.Info("scheduler start task")
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		// Контекст прокинут до реквестов и базы
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
		defer cancel()

		retryJob(ctx,
			task,
			config.Get().Retries,
			time.Duration(config.Get().RetriesTimeMinutes)*time.Minute)

		<-ticker.C
	}
}

func retryJob(ctx context.Context, task Task, maxRetries int, retryInterval time.Duration) {
	for i := 0; i < maxRetries; i++ {
		err := task(ctx)
		if err == nil {
			return
		}
		// последняя попытка, выводим ошибку
		if i == maxRetries-1 {
			slog.Error("task failed after maximum retries:", err)
		} else {
			slog.Debug("task failed, retrying in %s...", retryInterval)

			// Ожидание или отмена по контексту.
			select {
			case <-ctx.Done():
				slog.Error("task context cancelled:", ctx.Err())
				return
			case <-time.After(retryInterval):
			}
		}
	}
}

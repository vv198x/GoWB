package main

import (
	"fmt"
	"github.com/vv198x/GoWB/config"
	"github.com/vv198x/GoWB/logger"
	"github.com/vv198x/GoWB/tasks"
	"log/slog"
)

func main() {
	slog.SetDefault(logger.NewLogger(config.Get().LogLevel))
	fmt.Println(tasks.GetAdStatus([]int{17196078}))
	fmt.Println(tasks.GetAdBalance(17196078))

}

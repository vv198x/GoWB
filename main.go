package main

import (
	"context"
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
	//логирование запросов если логлевел дебаг
	pgsql.DebugPG()
	//миграция
	migrations.Start()
	//инициализация репозитория
	repository.R = &repository.AdCampaignRepository{DB: pgsql.DB}

	/*


				go scheduler.Add(tasks.AutoReFill, 30*time.Minute)
				go scheduler.Add(tasks.UpdateNames, 24*time.Hour)

				select {}

			Высчитать шаг
			проверить старую ставку и текущую если меньше на шаг значит не увеличилась
			•	Если текущая позиция выше 10 (т.е. >10):
			o	  Увеличьте ставку на заранее определенный шаг, чтобы улучшить позицию.
			o	  Проверьте, не превышает ли новая ставка максимальную допустимую ставку. Если превышает, установите максимальную ставку.
			•	Если текущая позиция ниже 10 (т.е. <=10):
			o	  Уменьшите ставку на заранее определенный шаг, чтобы снизить позицию.
			•  Логирование и отслеживание:
			•	  Записывайте каждое изменение ставки и текущую позицию для анализа и оптимизации стратегии в будущем.

		17182684
		3468
		{
		  "advertId": 17182684,
		  "type": 9,
		  "cpm": 1350,
		  "param": 3468,
		  "instrument": 4
		}

		арк
		17655713
		125 р
		1103


		{
		  "advertId": 17655713,
		  "type": 8,
		  "cpm": 150,
		  "param": 1103,
		  "instrument": 4
		}
	*/

	fmt.Println(tasks.UpdateNames(context.Background()))
	fmt.Println(tasks.CheckPositions(context.Background()))
	fmt.Println(repository.Do().GetBidderInfoByAdID(context.Background(), 17182684))

}

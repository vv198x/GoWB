package tasks

import (
	"context"
	"fmt"
)

func AutoReFill(ctx context.Context) error {
	if err := GetAdList(ctx); err != nil {
		return fmt.Errorf("Get AdList err ", err)
	}

	if err := UpdateBalance(ctx); err != nil {
		return fmt.Errorf("UpdateBalance err ", err)
	}

	if err := ReFillBalance(ctx); err != nil {
		return fmt.Errorf("ReFillBalance err ", err)
	}
	//TODO - Запустить компании не помеченные DoNotRefill = false
	return UpdateBalance(ctx)
}

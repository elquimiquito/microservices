package cron

import (
	"Wildberries_services/repository/proxy"
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			employees := proxy.GetBirthdays()
			fmt.Println(employees)
		}
		time.Sleep(time.Hour * 24)
	}
}

func Init(ctx context.Context) {
	go task(ctx)
}

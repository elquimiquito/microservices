package main

import (
	cron "Wildberries_services/cron/internal"
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cron.Init(ctx)
	time.Sleep(time.Hour * 240)
	cancel()
}

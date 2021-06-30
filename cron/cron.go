package cron

import (
	"os"
	"os/signal"

	"github.com/cosmos-gaminghub/explorer-backend/client"
	"github.com/cosmos-gaminghub/explorer-backend/conf"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/robfig/cron"
)

// Start starts to create cron jobs which fetches chosen asset list information and
// store them in database every hour and every 24 hours.
func Start() error {
	logger.Info("Starting cron jobs...")

	cron := cron.New()

	// Every hour
	cron.AddFunc("0 */20 * * * *", func() {
		client.SaveMarketChartRange(conf.Get().Hub.Coin, 20)
		logger.Info("successfully saved asset information list 1H")
	})

	go cron.Start()

	// Allow graceful closing of the governance loop
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	return nil
}

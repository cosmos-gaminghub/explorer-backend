package main

import (
	"github.com/cosmos-gaminghub/explorer-backend/cron"
	"github.com/cosmos-gaminghub/explorer-backend/exporter"
)

func main() {
	exporter.Start()
	cron.Start()
}

package cron

// Start starts to create cron jobs which fetches chosen asset list information and
// store them in database every hour and every 24 hours.
// func Start() error {
// 	log.Println("Starting cron jobs...")

// 	cron := cron.New()

// 	// Every hour
// 	cron.AddFunc("0 0 * * * *", func() {
// 		AssetInfoList1H()
// 		log.Println("successfully saved asset information list 1H")
// 	})

// 	// Every 24 hours at 1:00 AM UTC timezone
// 	cron.AddFunc("0 0 1 * * *", func() {
// 		AssetInfoList24H()
// 		log.Println("successfully saved asset information list 24H")
// 	})

// 	go cron.Start()

// 	// Allow graceful closing of the governance loop
// 	signalCh := make(chan os.Signal, 1)
// 	signal.Notify(signalCh, os.Interrupt)
// 	<-signalCh

// 	return nil
// }

// AssetInfoList1H fetches asset information list and save them
// func AssetInfoList1H() ([]*schema.StatAssetInfoList1H, error) {
// 	assetInfoList := make([]*schema.StatAssetInfoList1H, 0)

// 	for _, assetName := range models.ChosenAssetNames {
// 		asset, err := c.client.Asset(assetName)
// 		if err != nil {
// 			log.Printf("failed to request client Asset: %s", err)
// 		}

// 		tempAssetInfo := &schema.StatAssetInfoList1H{
// 			Asset:        asset.Asset,
// 			MappedAsset:  asset.MappedAsset,
// 			Name:         asset.Name,
// 			AssetImage:   asset.AssetImg,
// 			Price:        asset.Price,
// 			QuoteUnit:    asset.QuoteUnit,
// 			ChangeRange:  asset.ChangeRange,
// 			Supply:       asset.Supply,
// 			Marketcap:    asset.Price * asset.Supply,
// 			Owner:        asset.Owner,
// 			Transactions: asset.Transactions,
// 			Holders:      asset.Holders,
// 		}

// 		assetInfoList = append(assetInfoList, tempAssetInfo)
// 	}

// 	err := c.db.SaveAssetInfoList1H(assetInfoList)
// 	if err != nil {
// 		log.Printf("failed to save AssetInfoList: %s", err)
// 	}

// 	return assetInfoList, nil
// }

// // AssetInfoList24H fetches asset information list and save them
// func AssetInfoList24H() ([]*schema.StatAssetInfoList24H, error) {
// 	assetInfoList := make([]*schema.StatAssetInfoList24H, 0)

// 	for _, assetName := range models.ChosenAssetNames {
// 		asset, err := c.client.Asset(assetName)
// 		if err != nil {
// 			log.Printf("failed to request client Asset: %s", err)
// 		}

// 		tempAssetInfo := &schema.StatAssetInfoList24H{
// 			Asset:        asset.Asset,
// 			MappedAsset:  asset.MappedAsset,
// 			Name:         asset.Name,
// 			AssetImage:   asset.AssetImg,
// 			Price:        asset.Price,
// 			QuoteUnit:    asset.QuoteUnit,
// 			ChangeRange:  asset.ChangeRange,
// 			Supply:       asset.Supply,
// 			Marketcap:    asset.Price * asset.Supply,
// 			Owner:        asset.Owner,
// 			Transactions: asset.Transactions,
// 			Holders:      asset.Holders,
// 		}

// 		assetInfoList = append(assetInfoList, tempAssetInfo)
// 	}

// 	err := c.db.SaveAssetInfoList24H(assetInfoList)
// 	if err != nil {
// 		log.Printf("failed to save AssetInfoList: %s", err)
// 	}

// 	return assetInfoList, nil
// }

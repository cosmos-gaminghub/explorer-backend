package exporter

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/client"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/pkg/errors"
)

var (
	// Version is this application's version.
	Version = "Development"

	// Commit is this application's commit hash.
	Commit = ""
)

// Exporter wraps the required params to export blockchain
type Exporter struct {
}

// Start starts to synchronize Binance Chain data.
func Start() error {
	fmt.Println("Starting Chain Exporter...")

	go func() {
		for {
			fmt.Println("start - sync blockchain")
			err := sync()
			if err != nil {
				fmt.Sprintf("error - sync blockchain: %v\n", err)
			}
			fmt.Println("finish - sync blockchain")
			time.Sleep(time.Second)
		}
	}()

	for {
		select {}
	}
}

// sync compares block height between the height saved in your database and
// the latest block height on the active chain and calls process to start ingesting data.
func sync() error {
	// Query latest block height saved in database
	block, err := document.Block{}.QueryLatestBlockFromDB()
	if err != nil {
		logger.Error("Block not found")
	}

	dbHeight, err := strconv.ParseInt(block.Height, 10, 64)
	if dbHeight == -1 {
		log.Fatal(errors.Wrap(err, "failed to query the latest block height saved in database"))
	}

	fmt.Println(block)

	latestBlockHeight, err := client.GetLatestBlockHeight()
	if err != nil {
		logger.Error("Not found block from lcd", logger.String("err", err.Error()))
	}
	// Synchronizing blocks from the scratch will return 0 and will ingest accordingly.
	// Skip the first block since it has no pre-commits
	if dbHeight == 0 {
		dbHeight = 1
	}

	// Ingest all blocks up to the latest height
	for i := dbHeight + 1; i <= latestBlockHeight; i++ {
		err = process(i)
		if err != nil {
			return err
		}
		fmt.Printf("synced block %d/%d \n", i, latestBlockHeight)
	}

	return nil
}

// process ingests chain data, such as block, transaction, validator set information
// and save them in database
func process(height int64) error {
	block, err := client.GetBlock(height)
	if err != nil {
		return fmt.Errorf("failed to query block using rpc client: %s", err)
	}
	orm.Save("block", block)

	// resultTxs, err := ex.getTxs(block)
	// if err != nil {
	// 	return fmt.Errorf("failed to get transactions: %s", err)
	// }

	// valSet, err := ex.client.GetValidatorSet(block.Block.LastCommit.Height())
	// if err != nil {
	// 	return fmt.Errorf("failed to query validator set using rpc client: %s", err)
	// }

	// vals, err := ex.client.GetValidators()
	// if err != nil {
	// 	return fmt.Errorf("failed to query validators using rpc client: %s", err)
	// }

	// // TODO: Reward Fees Calculation
	// resultBlock, err := ex.getBlock(block)
	// if err != nil {
	// 	return fmt.Errorf("failed to get block: %s", err)
	// }

	// resultValidators, err := ex.getValidators(vals)
	// if err != nil {
	// 	return fmt.Errorf("failed to get validators: %s", err)
	// }

	// resultPreCommits, err := ex.getPreCommits(block.Block.LastCommit, valSet)
	// if err != nil {
	// 	return fmt.Errorf("failed to get precommits: %s", err)
	// }

	// err = ex.db.InsertExportedData(resultBlock, resultTxs, resultValidators, resultPreCommits)
	// if err != nil {
	// 	return fmt.Errorf("failed to insert exporterd data: %s", err)
	// }

	return nil
}

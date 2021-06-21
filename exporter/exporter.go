package exporter

import (
	"fmt"
	"log"
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/client"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
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

	go func() {
		for {
			fmt.Println("start - sync proposal")
			err := syncProposal()
			if err != nil {
				fmt.Sprintf("error - sync proposal blockchain: %v\n", err)
			}
			fmt.Println("finish - sync proposal blockchain")
			time.Sleep(3600 * time.Second)
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

	dbHeight := block.Height
	fmt.Println(dbHeight)
	if dbHeight == -1 {
		log.Fatal(errors.Wrap(err, "failed to query the latest block height saved in database"))
	}

	latestBlockHeight, err := client.GetLatestBlockHeight()
	if err != nil {
		logger.Error("Not found block from lcd", logger.String("err", err.Error()))
	}
	// Synchronizing blocks from the scratch will return 0 and will ingest accordingly.
	// Skip the first block since it has no pre-commits
	if dbHeight == 0 {
		dbHeight = 5200790
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
		logger.Error("failed to query block using rpc client:", logger.String("err", err.Error()))
	}

	txs, err := client.GetTxs(height)
	if err != nil {
		logger.Error("failed to get transactions:", logger.String("err", err.Error()))
	}
	// lastCommitHeight, err := strconv.ParseInt(block.Block.LastCommit.Height, 10, 64)
	// valSet, err := client.GetValidatorSet(lastCommitHeight)
	// if err != nil {
	// 	return fmt.Errorf("failed to query validator set using rpc client: %s", err)
	// }

	vals, err := client.GetValidators()
	if err != nil {
		logger.Error("failed to query validators using rpc client:", logger.String("err", err.Error()))
	}

	// TODO: Reward Fees Calculation
	resultBlock, err := GetBlock(block)
	if err != nil {
		logger.Error("failed to get block:", logger.String("err", err.Error()))
	}
	orm.Save(document.CollectionNmBlock, resultBlock)

	resultTxs, err := GetTxs(txs, *resultBlock)
	if err != nil {
		logger.Error("failed to get txs:", logger.String("err", err.Error()))
	}
	for _, item := range resultTxs {
		orm.Save(document.CollectionNmCommonTx, item)
	}

	validatorSets, err := client.GetValidatorSet(height, 0)
	if err != nil {
		logger.Error("failed to get validator set:", logger.String("err", err.Error()))
	}

	SaveMissedBlock(vals, validatorSets, block)

	resultValidators, err := GetValidators(vals, validatorSets)
	if err != nil {
		logger.Error("failed to get validators:", logger.String("err", err.Error()))
	}
	for _, item := range resultValidators {
		SaveValidator(*item)
	}

	// resultPreCommits, err := GetPreCommits(block.Block.LastCommit, valSet)
	// if err != nil {
	// 	return fmt.Errorf("failed to get precommits: %s", err)
	// }

	// err = ex.db.InsertExportedData(resultBlock, resultTxs, resultValidators, resultPreCommits)
	// if err != nil {
	// 	return fmt.Errorf("failed to insert exporterd data: %s", err)
	// }

	return nil
}

func syncProposal() error {
	proposals, err := client.GetProposals()
	if err != nil {
		return nil
	}

	resultProposals, err := GetProposals(proposals)
	if err != nil {
		return fmt.Errorf("failed to get validators: %s", err)
	}
	for _, proposal := range resultProposals {
		err = processProposal(*proposal)
		if err != nil {
			return err
		}
		fmt.Printf("synced proposal %d \n", proposal.ProposalId)
	}

	return nil
}

func processProposal(proposal schema.Proposal) error {
	deposits, _ := client.GetProposalDeposits(proposal.ProposalId)
	if len(deposits.Proposals) > 0 {
		GetDeposits(deposits, &proposal)
	}

	votes, _ := client.GetProposalVotes(proposal.ProposalId, 0)
	if len(votes) > 0 {
		proposal.Vote = votes
	}

	proposer, err := client.GetProposalProposer(proposal.ProposalId)
	if err != nil {
		logger.Fatal("Get proposal error")
	}

	proposal.Proposer = proposer.Result.Proposer
	SaveProposal(proposal)

	return nil
}

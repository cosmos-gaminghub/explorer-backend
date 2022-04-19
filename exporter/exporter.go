package exporter

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/client"
	"github.com/cosmos-gaminghub/explorer-backend/conf"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/pkg/errors"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

var (
	// Version is this application's version.
	Version = "Development"

	// Commit is this application's commit hash.
	Commit = ""
)

var clientHTTP *rpchttp.HTTP

// Exporter wraps the required params to export blockchain
type Exporter struct {
	clientHTTP *rpchttp.HTTP
}

// Start starts to synchronize Chain data.
func Start() error {
	fmt.Println("Starting Chain Exporter...")

	var (
		err error
	)

	clientHTTP, err = rpchttp.New(conf.Get().Hub.RpcUrl, "/websocket")
	if err != nil {
		fmt.Println(err.Error())
	}

	// go func() {
	// 	for {
	// 		fmt.Println("start - sync blockchain")
	// 		err := sync()
	// 		if err != nil {
	// 			fmt.Printf("error - sync blockchain: %v\n", err)
	// 		}
	// 		fmt.Println("finish - sync blockchain")
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		fmt.Println("start - sync proposal")
	// 		err := syncProposal()
	// 		if err != nil {
	// 			fmt.Printf("error - sync proposal blockchain: %v\n", err)
	// 		}
	// 		fmt.Println("finish - sync proposal blockchain")
	// 		time.Sleep(3600 * time.Second)
	// 	}
	// }()

	go func() {
		for {
			fmt.Println("start - sync codes")
			err := syncCode()
			if err != nil {
				fmt.Printf("error - sync code blockchain: %v\n", err)
			}
			fmt.Println("finish - sync code blockchain")
			time.Sleep(3600 * 24 * time.Second)
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
		dbHeight = int64(conf.Get().Db.SyncFromHeight)
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

	vals, err := client.GetValidators(0)
	if err != nil {
		logger.Error("failed to query validators using rpc client:", logger.String("err", err.Error()))
	}

	// TODO: Reward Fees Calculation
	resultBlock, err := GetBlock(block)
	if err != nil {
		logger.Error("failed to get block:", logger.String("err", err.Error()))
	}
	orm.Save(document.CollectionNmBlock, resultBlock)

	SaveMissedBlock(clientHTTP, height, resultBlock)

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

	resultValidators, err := GetValidators(vals, validatorSets)
	if err != nil {
		logger.Error("failed to get validators:", logger.String("err", err.Error()))
	}
	for _, item := range resultValidators {
		SaveValidator(*item)
	}

	SaveAccountTransaction(resultValidators, resultTxs)

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
	depositResult, _ := client.GetProposalDeposits(proposal.ProposalId)
	if len(depositResult.Deposits) > 0 {
		deposits, err := GetDeposits(depositResult)
		if err != nil {
			logger.Error("failed to get deposits:", logger.String("err", err.Error()))
		}
		for _, item := range deposits {
			orm.Save(document.CollectionDeposit, item)
		}
	}

	voteResult, _ := client.GetProposalVotes(proposal.ProposalId, 0)
	if len(voteResult) > 0 {
		votes, err := GetVote(voteResult)

		if err != nil {
			logger.Error("failed to get votes:", logger.String("err", err.Error()))
		}
		for _, item := range votes {
			orm.Save(document.CollectionVote, item)
		}
	}

	proposer, err := client.GetProposalProposer(proposal.ProposalId)
	if err != nil {
		logger.Fatal("Get proposal error")
	}

	proposal.Proposer = proposer.Result.Proposer
	SaveProposal(proposal)

	return nil
}

func syncCode() error {
	codes, err := client.GetListWasmCode()
	if err != nil {
		return nil
	}

	resultCode := GetCodes(codes)
	if err != nil {
		return fmt.Errorf("failed to get codes: %s", err)
	}
	for _, code := range resultCode {
		err = processCode(code)
		if err != nil {
			return err
		}
		fmt.Printf("synced code %d \n", code.CodeId)
	}

	return nil
}

func processCode(code *schema.Code) error {
	contractResult, _ := client.GetListWasmCodeContracts(code.CodeId)

	height, err := strconv.ParseInt(contractResult.Height, 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to parse height %s for code %d:", contractResult.Height, code.CodeId))
		return err
	}
	block, err := client.GetBlock(height)
	if err != nil {
		logger.Error("failed to query block using rpc client:", logger.String("err", err.Error()))
	}
	fmt.Println(block)

	code.SetInstantiateCount(len(contractResult.ContractAddress)).
		SetCreatedAt(block.Block.Header.Time)

	_, err = SaveCode(code)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to save code %d:", code.CodeId))
		return err
	}
	return nil
}

package client

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
	"github.com/pkg/errors"
)

// GetBlock queries for a block by height. An error is returned if the query fails.
func GetBlock(height int64) (types.BlockResult, error) {
	url := fmt.Sprintf("http://108.61.162.170:1317/cosmos/base/tendermint/v1beta1/blocks/%d", height)
	resBytes, err := utils.Get(url)
	if err != nil {
		//logger.Error("get AssetTokens error", logger.String("err", err.Error()))
	}

	var result types.BlockResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		//logger.Error("Unmarshal AssetTokens error", logger.String("err", err.Error()))
	}

	return result, nil
}

// GetLatestBlockHeight returns the latest block height on the active chain.
func GetLatestBlockHeight() (int64, error) {
	resBytes, err := utils.Get("http://108.61.162.170:1317/cosmos/base/tendermint/v1beta1/blocks/latest")
	if err != nil {
		//logger.Error("get AssetTokens error", logger.String("err", err.Error()))
	}

	var result types.BlockResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		//logger.Error("Unmarshal AssetTokens error", logger.String("err", err.Error()))
	}

	latestBlockHeight, err := strconv.ParseInt(result.Block.Header.Height, 10, 64)
	if latestBlockHeight == -1 {
		log.Fatal(errors.Wrap(err, "failed to query the latest block height on the active network"))
	}

	return latestBlockHeight, nil
}

// GetTxs queries for all the transactions in a block height.
// It uses `Tx` RPC method to query for the transaction.
// func GetTxs(block *tmctypes.ResultBlock) ([]*rpc.ResultTx, error) {
// 	txs := make([]*rpc.ResultTx, len(block.Block.Txs), len(block.Block.Txs))
// 	var err error
// 	retryFlag := false

// 	for i, tmTx := range block.Block.Txs {
// 		hash := tmTx.Hash()
// 		controler <- struct{}{}
// 		wg.Add(1)
// 		go func(i int, hash []byte) {
// 			defer func() {
// 				<-controler
// 				wg.Done()
// 			}()

// 			txs[i], err = c.rpcClient.Tx(hash, true)
// 			if err != nil {
// 				retryFlag = true
// 				fmt.Println(hash)
// 				return
// 			}
// 		}(i, hash)
// 	}
// 	wg.Wait()

// 	if retryFlag {
// 		return nil, fmt.Errorf("can not get all of txs, retry get tx in block height = %d", block.Block.Height)
// 	}

// 	// tx, err := c.rpcClient.Tx(hash, true)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// txs[i] = tx

// 	return txs, nil
// }

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
// func (c Client) GetValidatorSet(height int64) (*tmctypes.ResultValidators, error) {
// 	return c.rpcClient.Validators(&height)
// }

// // GetValidators returns validators detail information in Tendemrint validators in active chain
// // An error returns if the query fails.
// func (c Client) GetValidators() ([]*types.Validator, error) {
// 	resp, err := c.apiClient.R().Get("/stake/validators")
// 	if err != nil {
// 		return nil, err
// 	}

// 	var vals []*types.Validator

// 	err = json.Unmarshal(resp.Body(), &vals)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return vals, nil
// }

// // GetTokens returns information about existing tokens in active chain.
// func (c Client) GetTokens(limit int, offset int) ([]*types.Token, error) {
// 	resp, err := c.apiClient.R().Get("/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
// 	if err != nil {
// 		return nil, err
// 	}

// 	var tokens []*types.Token
// 	err = json.Unmarshal(resp.Body(), &tokens)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tokens, nil
// }

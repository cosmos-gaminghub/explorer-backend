package client

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos-gaminghub/explorer-backend/conf"
	"github.com/cosmos-gaminghub/explorer-backend/lcd"
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// GetBlock queries for a block by height. An error is returned if the query fails.
func GetBlock(height int64) (types.BlockResult, error) {
	url := fmt.Sprintf(lcd.UrlBlock, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get AssetTokens error", logger.String("err", err.Error()))
	}

	var result types.BlockResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal AssetTokens error", logger.String("err", err.Error()))
	}

	return result, nil
}

// GetLatestBlockHeight returns the latest block height on the active chain.
func GetLatestBlockHeight() (int64, error) {
	url := fmt.Sprintf(lcd.UrlBlockLatest, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get block error", logger.String("err", err.Error()))
	}

	var result types.BlockResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal block error", logger.String("err", err.Error()))
	}
	latestBlockHeight, _ := utils.ParseInt(result.Block.Header.Height)
	return latestBlockHeight, nil
}

// GetTxs queries for all the transactions in a block height.
// It uses `Tx` RPC method to query for the transaction.
func GetTxs(height int64) (types.TxResult, error) {
	url := fmt.Sprintf(lcd.UrlTxsTxHeight, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get Tx error", logger.String("err", err.Error()))
	}

	var result types.TxResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal Tx error", logger.String("err", err.Error()))
	}

	return result, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func GetValidatorSet(height int64, offset int64) ([]types.ValidatorOfValidatorSet, error) {
	url := fmt.Sprintf(lcd.UrlValidatorSet, conf.Get().Hub.LcdUrl, height, offset)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get validator set error", logger.String("err", err.Error()))
	}

	var result types.ValidatorSet
	validators := make([]types.ValidatorOfValidatorSet, 0, 125)
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal validator set error", logger.String("err", err.Error()))
	}
	validators = append(validators, result.Validators...)

	if len(validators) == lcd.DefaultValidatorSetLimit {
		valSet, _ := GetValidatorSet(height, lcd.DefaultValidatorSetLimit)
		validators = append(validators, valSet...)
	}

	return validators, nil
}

func FormatValidatorSetPubkeyToIndex(valSets []types.ValidatorOfValidatorSet) map[string]int {
	validatorSets := make(map[string]int)
	for index, valSet := range valSets {
		validatorSets[valSet.PubKey.Key] = index
	}
	return validatorSets
}

func FormatValidatorSetPubkeyToAddress(valSets []types.ValidatorOfValidatorSet) map[string]string {
	validatorSets := make(map[string]string)
	for _, valSet := range valSets {
		validatorSets[valSet.PubKey.Key] = valSet.ConsensusAddr
	}
	return validatorSets
}

// GetValidators returns validators detail information in Tendemrint validators in active chain
// An error returns if the query fails.
func GetValidators() (types.ValidatorsResult, error) {
	url := fmt.Sprintf(lcd.UrlValidators, conf.Get().Hub.LcdUrl, types.DefaultValidatorLimit)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get validators error", logger.String("err", err.Error()))
	}

	var result types.ValidatorsResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal validators error", logger.String("err", err.Error()))
	}

	return result, nil
}

func GetAuthParams() (types.AuthParam, error) {
	url := fmt.Sprintf(lcd.UrlModuleParam, conf.Get().Hub.LcdUrl, "auth")
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get validators error", logger.String("err", err.Error()))
	}

	var result types.AuthParam
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal validators error", logger.String("err", err.Error()))
	}

	return result, nil
}

func GetBankParams() (types.BankParam, error) {
	url := fmt.Sprintf(lcd.UrlModuleParam, conf.Get().Hub.LcdUrl, "bank")
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get bank param error", logger.String("err", err.Error()))
	}

	var result types.BankParam
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal bank param error", logger.String("err", err.Error()))
	}

	return result, nil
}

func GetDistributionParams() (types.DistributionParam, error) {
	url := fmt.Sprintf(lcd.UrlModuleParam, conf.Get().Hub.LcdUrl, "distribution")
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get distribution param error", logger.String("err", err.Error()))
	}

	var result types.DistributionParam
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal distribution param error", logger.String("err", err.Error()))
	}

	return result, nil
}

func GetGovParams(govType string) (types.GovParam, error) {
	url := fmt.Sprintf(lcd.UrlModuleParam, conf.Get().Hub.LcdUrl, govType)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get gov params error", logger.String("err", err.Error()))
	}

	var result types.GovParam
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal gov params error", logger.String("err", err.Error()))
	}

	return result, nil
}

// GetTokens returns information about existing tokens in active chain.
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

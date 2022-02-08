package client

import (
	"encoding/json"
	"fmt"
	"strconv"

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
func GetTxs(height int64) (txs types.TxResult, err error) {
	url := fmt.Sprintf(lcd.UrlTxsTxHeight, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error(fmt.Sprintf("get Tx error for height %d", height), logger.String("err", err.Error()))
		return txs, err
	}

	if err := json.Unmarshal(resBytes, &txs); err != nil {
		logger.Error(fmt.Sprintf("Unmarshal Tx error for height %d", height), logger.String("err", err.Error()))
	}
	return txs, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func GetValidatorSet(height int64, offset int64) ([]types.ValidatorOfValidatorSet, error) {
	url := fmt.Sprintf(lcd.UrlValidatorSet, conf.Get().Hub.LcdUrl, height, offset)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error(fmt.Sprintf("Get validator set error for height %d", height), logger.String("err", err.Error()))
		return []types.ValidatorOfValidatorSet{}, err
	}

	var result types.ValidatorSet
	validators := make([]types.ValidatorOfValidatorSet, 0, lcd.DefaultValidatorSetLimit)
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error(fmt.Sprintf("Unmarshal validator set error for height %d", height), logger.String("err", err.Error()))
	}
	validators = append(validators, result.Validators...)

	if len(validators) == lcd.DefaultValidatorSetLimit {
		valSet, _ := GetValidatorSet(height, offset+lcd.DefaultValidatorSetLimit)
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
func GetValidators(offset int) ([]types.Validator, error) {
	url := fmt.Sprintf(lcd.UrlValidators, conf.Get().Hub.LcdUrl, offset)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get validators error", logger.String("err", err.Error()))
	}

	var result types.ValidatorsResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal validators error", logger.String("err", err.Error()))
	}

	total, _ := strconv.Atoi(result.Pagination.Total)
	validators := make([]types.Validator, 0, total)
	validators = append(validators, result.Validators...)

	if len(validators) == lcd.DefaultValidatorSetLimit {
		vals, _ := GetValidators(types.DefaultValidatorSetLimit + offset)
		validators = append(validators, vals...)
	}

	return validators, nil
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

func GetProposals() (types.ProposalResult, error) {
	url := fmt.Sprintf(lcd.UrlProposal, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get proposal error", logger.String("err", err.Error()))
	}

	var result types.ProposalResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal proposal error", logger.String("err", err.Error()))
	}

	return result, nil
}

func GetProposalDeposits(proposalId int) (types.ProposalDepositResult, error) {
	url := fmt.Sprintf(lcd.UrlProposalDeposit, conf.Get().Hub.LcdUrl, proposalId)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get proposal depposit error", logger.String("err", err.Error()))
	}

	var result types.ProposalDepositResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal proposal depposit error", logger.String("err", err.Error()))
	}

	return result, nil
}

func GetProposalProposer(proposalId int) (types.ProposalProposerResult, error) {
	url := fmt.Sprintf(lcd.UrlProposalProposer, conf.Get().Hub.LcdUrl, proposalId)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get proposal proposer error", logger.String("err", err.Error()))
	}

	var result types.ProposalProposerResult
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal proposal proposer error", logger.String("err", err.Error()))
	}

	return result, nil
}

func GetProposalVotes(proposalId int, offset int) ([]types.ProposalVote, error) {
	url := fmt.Sprintf(lcd.UrlProposalVoters, conf.Get().Hub.LcdUrl, proposalId, offset)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get proposal votes error", logger.String("err", err.Error()))
	}

	var result types.ProposalVoteResult

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal proposal votes error", logger.String("err", err.Error()))
	}
	total, _ := strconv.Atoi(result.Pagination.Total)
	votes := make([]types.ProposalVote, 0, total)
	votes = append(votes, result.Votes...)

	if len(votes) == lcd.DefaultValidatorSetLimit {
		vs, _ := GetProposalVotes(proposalId, offset+lcd.DefaultValidatorSetLimit)
		votes = append(votes, vs...)
	}

	return votes, nil
}

func GetValSigningInfo(consensusAddress string) (types.ValSigningInfo, error) {
	url := fmt.Sprintf(lcd.UrlValSigningInfo, conf.Get().Hub.LcdUrl, consensusAddress)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get val signing info error address:"+consensusAddress, logger.String("err", err.Error()))
	}

	var result types.ValSigningInfo
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error(fmt.Sprintf("Unmarshal val signing info error address: %s", consensusAddress), logger.String("err", err.Error()))
	}

	return result, nil
}

func GetProposalTally(id int) (types.PropsalTally, error) {
	url := fmt.Sprintf(lcd.UrlProposalTally, conf.Get().Hub.LcdUrl, id)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error(fmt.Sprintf("Get proposal tally error proposal: %d", id), logger.String("err", err.Error()))
	}

	var result types.PropsalTally
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error(fmt.Sprintf("Unmarshal proposal tally error proposal: %d", id), logger.String("err", err.Error()))
	}

	return result, nil
}

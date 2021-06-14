package exporter

import (
	"fmt"

	"github.com/cosmos-gaminghub/explorer-backend/client"
	"github.com/cosmos-gaminghub/explorer-backend/conf"
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getValidators parses validators information and wrap into Precommit schema struct
func GetValidators(vals types.ValidatorsResult, validatorSets []types.ValidatorOfValidatorSet) (validators []*schema.Validator, err error) {
	validatorSetsFormat := client.FormatValidatorSetPubkeyToAddress(validatorSets)
	for _, validator := range vals.Validators {

		validatorFromDb, err := document.Validator{}.QueryValidatorDetailByOperatorAddr(validator.OperatorAddress)
		if err == nil {
			return nil, fmt.Errorf("unexpected error when checking validator existence: %s", err)
		}
		if validatorFromDb != (document.Validator{}) {
			document.Validator{}.UpdateByOperatorAddress(validatorFromDb)
			continue
		}
		tokens, _ := utils.ParseInt(validator.Tokens)
		var consensusAddress string
		if val, ok := validatorSetsFormat[validator.ConsensusPubkey.Key]; ok {
			consensusAddress = val
		}
		val := &schema.Validator{
			OperatorAddr:    validator.OperatorAddress,
			ConsensusAddres: consensusAddress,
			ConsensusPubkey: validator.ConsensusPubkey.Key,
			AccountAddr:     utils.Convert(conf.KeyPrefixAccAddr, validator.OperatorAddress),
			Jailed:          validator.Jailed,
			Status:          validator.Status,
			Tokens:          tokens,
			DelegatorShares: validator.DelegatorShares,
			Description:     validator.Description,
			UnbondingHeight: validator.UnbondingHeight,
			UnbondingTime:   validator.UnbondingTime,
			Commission:      validator.Commission,
		}
		validators = append(validators, val)
	}

	return validators, nil
}

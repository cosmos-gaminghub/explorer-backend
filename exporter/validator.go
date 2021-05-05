package exporter

import (
	"fmt"

	"github.com/cosmos-gaminghub/explorer-backend/conf"
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getValidators parses validators information and wrap into Precommit schema struct
func GetValidators(vals types.ValidatorsResult) (validators []*schema.Validator, err error) {
	for _, validator := range vals.Validators {

		_, err := document.Validator{}.QueryValidatorDetailByOperatorAddr(validator.OperatorAddress)
		if err == nil {
			return nil, fmt.Errorf("unexpected error when checking validator existence: %s", err)
		}
		tokens, _ := utils.ParseInt(validator.Tokens)
		val := &schema.Validator{
			OperatorAddr:    validator.OperatorAddress,
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

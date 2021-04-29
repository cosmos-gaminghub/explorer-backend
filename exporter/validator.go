package exporter

import (
	"fmt"

	"github.com/cosmos-gaminghub/explorer-backend/client"
	"github.com/cosmos-gaminghub/explorer-backend/conf"
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getValidators parses validators information and wrap into Precommit schema struct
func GetValidators(vals types.ValidatorsResult, block types.BlockResult) (validators []*schema.Validator, err error) {
	height, _ := utils.ParseInt(block.Block.Header.Height)
	validatorSets, _ := client.GetValidatorSet(height)
	for _, val := range vals.Validators {

		consensusAddr := utils.Convert(conf.KeyPrefixConsAddr, val.OperatorAddress)
		var indexOfValSet int
		fmt.Println(validatorSets)
		fmt.Println(block.Block.LastCommit.Signatures)
		if len(validatorSets.Validators) > 0 {
			for index, valSet := range validatorSets.Validators {
				if valSet.ConsensusAddr == consensusAddr {
					indexOfValSet = index
					break
				}
			}
			signedInfo := block.Block.LastCommit.Signatures[indexOfValSet]
			if signedInfo.Signature == "null" {
				b := schema.NewMissedBlock(schema.MissedBlock{
					Height:       height,
					OperatorAddr: val.OperatorAddress,
				})
				orm.Save("missed_block", b)
			}
		}
		_, err := document.Validator{}.QueryValidatorDetailByOperatorAddr(val.OperatorAddress)
		if err == nil {
			return nil, fmt.Errorf("unexpected error when checking validator existence: %s", err)
		}

		fmt.Println(val.Description)
		val := &schema.Validator{
			OperatorAddr:    val.OperatorAddress,
			ConsensusAddr:   consensusAddr,
			AccountAddr:     utils.Convert(conf.KeyPrefixAccAddr, val.OperatorAddress),
			Jailed:          val.Jailed,
			Status:          val.Status,
			Tokens:          val.Tokens,
			DelegatorShares: val.DelegatorShares,
			Description:     val.Description,
			UnbondingHeight: val.UnbondingHeight,
			UnbondingTime:   val.UnbondingTime,
			Commission:      val.Commission,
		}
		validators = append(validators, val)
	}

	return validators, nil
}

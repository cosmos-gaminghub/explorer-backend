package exporter

import (
	"encoding/base64"
	"strconv"

	"github.com/cosmos-gaminghub/explorer-backend/client"
	"github.com/cosmos-gaminghub/explorer-backend/conf"
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"gopkg.in/mgo.v2/bson"
)

// getValidators parses validators information and wrap into Precommit schema struct
func GetValidators(vals []types.Validator, validatorSets []types.ValidatorOfValidatorSet) (validators []*schema.Validator, err error) {
	validatorSetsFormat := client.FormatValidatorSetPubkeyToAddress(validatorSets)
	for _, validator := range vals {
		tokens, _ := utils.ParseInt(validator.Tokens)
		var consensusAddress string
		if val, ok := validatorSetsFormat[validator.ConsensusPubkey.Key]; ok {
			consensusAddress = val
		}
		_, decodeByte, _ := bech32.DecodeAndConvert(consensusAddress)
		str := base64.StdEncoding.EncodeToString(decodeByte)

		if validator.Description.ImageUrl == "" && validator.Description.Identity != "" {
			validator.Description.ImageUrl = client.GetImageUrl(validator.Description.Identity)
		}

		val := &schema.Validator{
			OperatorAddr:    validator.OperatorAddress,
			ConsensusAddres: consensusAddress,
			ConsensusPubkey: validator.ConsensusPubkey.Key,
			AccountAddr:     utils.Convert(conf.Get().Db.AddresPrefix, validator.OperatorAddress),
			Jailed:          validator.Jailed,
			Status:          validator.Status,
			Tokens:          tokens,
			DelegatorShares: validator.DelegatorShares,
			Description:     validator.Description,
			UnbondingHeight: validator.UnbondingHeight,
			UnbondingTime:   validator.UnbondingTime,
			Commission:      validator.Commission,
			ProposerAddr:    str,
		}
		validators = append(validators, val)
	}

	return validators, nil
}

func SaveValidator(validator schema.Validator) (interface{}, error) {
	selector := bson.M{document.ValidatorFieldOperatorAddress: validator.OperatorAddr}

	return orm.Upsert(document.CollectionNmValidator, selector, validator)
}

func saveTotalMissedBlock(validators []document.Validator) {
	for _, validator := range validators {
		var missedBlockCount int64
		if validator.ConsensusAddress != "" {
			valSigningInfo, err := client.GetValSigningInfo(validator.ConsensusAddress)
			if err == nil {
				missedBlockCount, _ = strconv.ParseInt(valSigningInfo.Info.MissedBlocksCount, 10, 64)
			}
		}
		orm.Update(document.CollectionNmValidator, bson.M{"consensus_address": validator.ConsensusAddress}, bson.M{"$set": bson.M{"total_missed_block": missedBlockCount}})
	}
}

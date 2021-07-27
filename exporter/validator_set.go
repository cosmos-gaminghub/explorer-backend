package exporter

import (
	"github.com/cosmos-gaminghub/explorer-backend/client"
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getValidators parses validators information and wrap into Precommit schema struct
func SaveMissedBlock(vals []types.Validator, validatorSets []types.ValidatorOfValidatorSet, block types.BlockResult) {
	height, _ := utils.ParseInt(block.Block.Header.Height)
	validatorSetsFormat := client.FormatValidatorSetPubkeyToIndex(validatorSets)
	for _, validator := range vals {
		if val, ok := validatorSetsFormat[validator.ConsensusPubkey.Key]; ok {
			if len(block.Block.LastCommit.Signatures) > 0 {
				signedInfo := block.Block.LastCommit.Signatures[val]
				if signedInfo.Signature == "" {
					b := schema.NewMissedBlock(schema.MissedBlock{
						Height:       height,
						OperatorAddr: validator.OperatorAddress,
						Timestamp:    block.Block.Header.Time,
					})
					orm.Save("missed_block", b)
				}
			}
		}
	}
}

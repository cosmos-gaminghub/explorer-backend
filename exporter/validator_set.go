package exporter

import (
	"github.com/cosmos-gaminghub/explorer-backend/client"
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getValidators parses validators information and wrap into Precommit schema struct
func SaveMissedBlock(vals types.ValidatorsResult, block types.BlockResult) {
	height, _ := utils.ParseInt(block.Block.Header.Height)
	validatorSets, _ := client.GetValidatorSet(height, 0)
	validatorSetsFormat := client.FormatValidatorSet(validatorSets)
	for _, validator := range vals.Validators {
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

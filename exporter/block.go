package exporter

import (
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getBlock exports block information.
func GetBlock(block types.BlockResult) (*schema.Block, error) {
	height, _ := utils.ParseInt(block.Block.Header.Height)
	b := schema.NewBlock(schema.Block{
		Height:     height,
		Proposer:   block.Block.Header.ProposerAddress,
		Moniker:    "",
		BlockHash:  block.BlockId.Hash,
		ParentHash: block.Block.Header.LastBlockID.Hash,
		NumTxs:     int64(len(block.Block.Data.Txs)),
		Timestamp:  block.Block.Header.Time,
	})

	return b, nil
}

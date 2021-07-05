package exporter

import (
	"encoding/base64"
	"encoding/hex"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getBlock exports block information.
func GetBlock(block types.BlockResult) (*schema.Block, error) {
	height, _ := utils.ParseInt(block.Block.Header.Height)
	decodeBase64, _ := base64.StdEncoding.DecodeString(block.BlockId.Hash)
	str := hex.EncodeToString(decodeBase64)
	b := schema.NewBlock(schema.Block{
		ChainId:   block.Block.Header.ChainID,
		Height:    height,
		Proposer:  block.Block.Header.ProposerAddress,
		BlockHash: str,
		NumTxs:    int64(len(block.Block.Data.Txs)),
		Timestamp: block.Block.Header.Time,
	})
	return b, nil
}

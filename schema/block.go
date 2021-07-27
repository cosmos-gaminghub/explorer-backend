package schema

import (
	"time"
)

// Block defines the structure for block information.
type Block struct {
	ChainId   string    `bson:"chain_id"`
	Height    int64     `bson:"height,omitempty"`
	Proposer  string    `bson:"proposer,omitempty"`
	Moniker   string    `bson:"moniker,omitempty"`
	BlockHash string    `bson:"block_hash"`
	NumTxs    int64     `bson:"num_txs"`
	Timestamp time.Time `bson:"timestamp"`
}

// NewBlock returns a new Block.
func NewBlock(b Block) *Block {
	return &Block{
		ChainId:   b.ChainId,
		Height:    b.Height,
		Proposer:  b.Proposer,
		Moniker:   b.Moniker,
		BlockHash: b.BlockHash,
		NumTxs:    b.NumTxs,
		Timestamp: b.Timestamp,
	}
}

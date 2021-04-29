package schema

import (
	"time"
)

// Block defines the structure for block information.
type Block struct {
	Height     int64     `bson:"height,omitempty"`
	Proposer   string    `bson:"proposer,omitempty"`
	Moniker    string    `bson:"moniker,omitempty"`
	BlockHash  string    `bson:"block_hash,omitempty"`
	ParentHash string    `bson:"parent_hash,omitempty"`
	NumTxs     int64     `bson:"num_txs"`
	Timestamp  time.Time `bson:"timestamp"`
}

// NewBlock returns a new Block.
func NewBlock(b Block) *Block {
	return &Block{
		Height:     b.Height,
		Proposer:   b.Proposer,
		Moniker:    b.Moniker,
		BlockHash:  b.BlockHash,
		ParentHash: b.ParentHash,
		NumTxs:     b.NumTxs,
		Timestamp:  b.Timestamp,
	}
}

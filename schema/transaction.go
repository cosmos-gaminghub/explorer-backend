package schema

import (
	"time"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
)

// Transaction defines the structure for transaction information.
type Transaction struct {
	Height     int64              `bson:"height"`
	TxHash     string             `bson:"tx_hash"`
	Code       uint32             `bson:"code"` // https://docs.binance.org/exchange-integration.html#important-ensuring-transaction-finality
	Messages   []types.TxMessages `bson:"messages"`
	Signatures []string           `bson:"signautures"`
	Memo       string             `bson:"memo"`
	GasWanted  int64              `bson:"gas_wanted"`
	GasUsed    int64              `bson:"gas_used"`
	Timestamp  time.Time          `bson:"timestamp"`
	Fee        types.Fee          `bson:"fee"`
}

// NewTransaction returns a new Transaction.
func NewTransaction(t Transaction) *Transaction {
	return &Transaction{
		Height:     t.Height,
		TxHash:     t.TxHash,
		Code:       t.Code,
		Messages:   t.Messages,
		Signatures: t.Signatures,
		Memo:       t.Memo,
		GasWanted:  t.GasWanted,
		GasUsed:    t.GasUsed,
		Timestamp:  t.Timestamp,
		Fee:        t.Fee,
	}
}

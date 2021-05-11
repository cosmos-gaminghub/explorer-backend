package schema

import (
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
)

// Transaction defines the structure for transaction information.
type Transaction struct {
	Height     int64               `bson:"height"`
	TxHash     string              `bson:"txhash"`
	Code       uint32              `bson:"code"`
	Memo       string              `bson:"memo"`
	GasWanted  int64               `bson:"gas_wanted"`
	GasUsed    int64               `bson:"gas_used"`
	Timestamp  string              `bson:"timestamp"`
	Logs       []types.Log         `bson:"logs" json:"logs"`
	Signatures []string            `bson:"signatures" json:"signatures"`
	Messages   []types.BodyMessage `bson:"messages" json:"messages"`
	Fee        types.Fee           `bson:"fee" json:"fee"`
}

// NewTransaction returns a new Transaction.
func NewTransaction(t Transaction) *Transaction {
	return &Transaction{
		Height:     t.Height,
		TxHash:     t.TxHash,
		Code:       t.Code,
		Memo:       t.Memo,
		GasWanted:  t.GasWanted,
		GasUsed:    t.GasUsed,
		Timestamp:  t.Timestamp,
		Fee:        t.Fee,
		Signatures: t.Signatures,
		Messages:   t.Messages,
	}
}

package document

import (
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionAccountTransaction = "account_transaction"

	AccountTransaction_Field_Height    = "height"
	AccountTransaction_Account_Address = "account_address"
)

type AccountTransaction struct {
	Height      int64  `bson:"height"`
	AccountAddr string `bson:"account_address"`
	TxHash      string `bson:"tx_hash"`
}

func (d AccountTransaction) Name() string {
	return CollectionAccountTransaction
}

func (d AccountTransaction) PkKvPair() map[string]interface{} {
	return bson.M{}
}

func (d AccountTransaction) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{AccountTransaction_Account_Address, AccountTransaction_Field_Height},
			Background: true,
		},
	}

	return indexes
}

func (d AccountTransaction) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}

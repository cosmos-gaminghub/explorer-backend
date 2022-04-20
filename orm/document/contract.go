package document

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionContract = "contracts"

	ContractAddress = "contract_address"
)

type Contract struct {
	CodeId           int       `bson:"code_id"`
	Contract         string    `bson:"contract"`
	ContractAddress  string    `bson:"contract_address"`
	Admin            string    `bson:"admin"`
	Creator          string    `bson:"creator"`
	ExecutedCount    int       `bson:"executed_count"`
	InstantiatedAt   time.Time `bson:"instantiated_at"`
	Label            string    `bson:"label"`
	LastExecutedAt   time.Time `bson:"last_executed_at"`
	Permission       string    `bson:"permission"`
	PermittedAddress string    `bson:"permitted_address"`
	TxHash           string    `bson:"txhash"`
	Version          string    `bson:"version"`
}

func (d Contract) Name() string {
	return CollectionContract
}

func (d Contract) PkKvPair() map[string]interface{} {
	return bson.M{}
}

func (d Contract) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{ContractAddress},
			Background: true,
		},
	}

	return indexes
}

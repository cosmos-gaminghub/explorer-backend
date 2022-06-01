package document

import (
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionContract = "contracts"

	ContractAddressField = "contract_address"
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
			Key:        []string{ContractAddressField},
			Background: true,
		},
	}

	return indexes
}

func (_ Contract) FindByContractAddress(contractAddress string) (schema.Contract, error) {
	var contract schema.Contract
	condition := bson.M{
		ContractAddressField: contractAddress,
	}

	err := queryOne(CollectionContract, nil, condition, &contract)
	return contract, err
}

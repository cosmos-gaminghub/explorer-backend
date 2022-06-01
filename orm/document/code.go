package document

import (
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionCode = "codes"

	CodeIdField = "code_id"
)

type Code struct {
	CodeId           int       `bson:"code_id"`
	Contract         string    `bson:"contract"`
	DataHash         string    `bson:"data_hash"`
	CreatedAt        time.Time `bson:"created_at"`
	Creator          string    `bson:"creator"`
	InstantiateCount int       `bson:"instantiate_count"`
	Permission       string    `bson:"permission"`
	PermittedAddress string    `bson:"permitted_address"`
	TxHash           string    `bson:"txhash"`
	Version          string    `bson:"version"`
}

func (d Code) Name() string {
	return CollectionCode
}

func (d Code) PkKvPair() map[string]interface{} {
	return bson.M{}
}

func (d Code) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{CodeIdField},
			Background: true,
		},
	}

	return indexes
}

func (_ Code) FindByCodeId(codeId int) (schema.Code, error) {
	var code schema.Code
	condition := bson.M{
		CodeIdField: codeId,
	}

	err := queryOne(CollectionCode, nil, condition, &code)
	return code, err
}

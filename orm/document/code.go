package document

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionCode = "codes"

	CodeId = "code_id"
)

type Code struct {
	CodeId           int       `bson:"code_id"`
	Contract         string    `bson:"contract"`
	DataHash         string    `bson:"data_hash"`
	CreatedAt        time.Time `bson:"created_at"`
	Creator          string    `bson:"creator"`
	InstantiateCount int       `bson:"instantiate_count" json:"instantiate_count"`
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
			Key:        []string{CodeId},
			Background: true,
		},
	}

	return indexes
}

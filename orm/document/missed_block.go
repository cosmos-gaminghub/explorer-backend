package document

import (
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

const (
	CollectionMissedBlock = "missed_block"

	MissedBlock_Field_Height           = "height"
	MissedBlock_Field_Operator_Address = "operator_address"
)

type MissedBlock struct {
	Height       int64     `bson:"height"`
	OperatorAddr string    `bson:"operator_address"`
	Timestamp    time.Time `bson:"timestamp"`
}

func (d MissedBlock) Name() string {
	return CollectionMissedBlock
}

func (d MissedBlock) PkKvPair() map[string]interface{} {
	return bson.M{}
}

func (d MissedBlock) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{MissedBlock_Field_Operator_Address},
			Background: true,
		},
		{
			Key:        []string{MissedBlock_Field_Height},
			Background: true,
		},
	}

	return indexes
}

func (d MissedBlock) Batch(txs []txn.Op) error {
	return orm.Batch(txs)
}

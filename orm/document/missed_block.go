package document

import "time"

type MissedBlock struct {
	Height       int64     `bson:"height"`
	OperatorAddr string    `bson:"operator_address"`
	Timestamp    time.Time `bson:"timestamp"`
}

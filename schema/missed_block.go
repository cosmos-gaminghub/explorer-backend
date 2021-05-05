package schema

import "time"

type MissedBlock struct {
	Height       int64     `bson:"height,omitempty"`
	OperatorAddr string    `bson:"operator_address,omitempty"`
	Timestamp    time.Time `bson:"timestamp"`
}

func NewMissedBlock(b MissedBlock) *MissedBlock {
	return &MissedBlock{
		Height:       b.Height,
		OperatorAddr: b.OperatorAddr,
		Timestamp:    b.Timestamp,
	}
}

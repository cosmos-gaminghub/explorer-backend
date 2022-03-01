package schema

import "time"

type MissedBlock struct {
	Height           int64     `bson:"height,omitempty"`
	ConsensusAddress string    `bson:"consensus_address,omitempty"`
	Timestamp        time.Time `bson:"timestamp"`
}

func NewMissedBlock(b MissedBlock) *MissedBlock {
	return &MissedBlock{
		Height:           b.Height,
		ConsensusAddress: b.ConsensusAddress,
		Timestamp:        b.Timestamp,
	}
}

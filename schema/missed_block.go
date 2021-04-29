package schema

type MissedBlock struct {
	Height       int64  `bson:"height,omitempty"`
	OperatorAddr string `bson:"operator_address,omitempty"`
}

func NewMissedBlock(b MissedBlock) *MissedBlock {
	return &MissedBlock{
		Height:       b.Height,
		OperatorAddr: b.OperatorAddr,
	}
}

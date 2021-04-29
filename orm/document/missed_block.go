package document

type MissedBlock struct {
	Height       int64  `bson:"height"`
	OperatorAddr string `bson:"operator_address"`
}

package schema

type Deposit struct {
	ProposalId int      `bson:"proposal_id"`
	Depositor  string   `bson:"depositor"`
	Amount     []Amount `bson:"amount"`
}

type Amount struct {
	Denom  string `bson:"denom"`
	Amount string `bson:"amount"`
}

// NewDeposit returns a new Block.
func NewDeposit(d Deposit) *Deposit {
	return &Deposit{
		ProposalId: d.ProposalId,
		Depositor:  d.Depositor,
		Amount:     d.Amount,
	}
}

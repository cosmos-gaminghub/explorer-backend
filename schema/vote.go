package schema

type Vote struct {
	ProposalId int    `bson:"proposal_id"`
	Voter      string `bson:"voter"`
	Option     string `bson:"option"`
}

// NewBlock returns a new Block.
func NewVote(v Vote) *Vote {
	return &Vote{
		ProposalId: v.ProposalId,
		Voter:      v.Voter,
		Option:     v.Option,
	}
}

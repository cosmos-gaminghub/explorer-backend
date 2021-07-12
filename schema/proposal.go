package schema

import (
	"time"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
)

// Transaction defines the structure for transaction information.
type Proposal struct {
	ProposalId       int                            `bson:"proposal_id"`
	Proposer         string                         `bson:"proposer"`
	ProposalStatus   string                         `bson:"proposal_status"`
	Content          types.ProposalContent          `bson:"content" json:"content"`
	SubmitTime       time.Time                      `bson:"submit_time"`
	DepositEndTime   time.Time                      `bson:"deposit_end_time"`
	FinalTallyResult types.ProposalFinalTallyResult `bson:"final_tally_result" json:"final_tally_result"`
	VotingEndTime    time.Time                      `bson:"voting_end_time"`
	VotingStartTime  time.Time                      `bson:"voting_start_time"`
	TotalDeposit     []ProposalAmount               `bson:"total_deposit"`
}

type ProposalAmount struct {
	Denom  string `bson:"denom"`
	Amount string `bson:"amount"`
}

// NewTransaction returns a new Transaction.
func NewProposal(t Proposal) *Proposal {
	return &Proposal{
		ProposalId:       t.ProposalId,
		Proposer:         t.Proposer,
		ProposalStatus:   t.ProposalStatus,
		Content:          t.Content,
		SubmitTime:       t.SubmitTime,
		FinalTallyResult: t.FinalTallyResult,
		VotingEndTime:    t.VotingEndTime,
		VotingStartTime:  t.VotingStartTime,
		DepositEndTime:   t.DepositEndTime,
	}
}

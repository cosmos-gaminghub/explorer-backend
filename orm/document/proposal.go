package document

import (
	"time"
)

const (
	CollectionProposal = "proposal"

	ProposalFieldProposalId = "proposal_id"
)

// Transaction defines the structure for transaction information.
type Proposal struct {
	ProposalId       int              `bson:"proposal_id"`
	ProposalStatus   string           `bson:"proposal_status"`
	Content          Content          `bson:"content" json:"content"`
	SubmitTime       time.Time        `bson:"submit_time"`
	FinalTallyResult FinalTallyResult `bson:"final_tally_result" json:"final_tally_result"`
	VotingEndTime    time.Time        `bson:"voting_end_time"`
	VotingStartTime  time.Time        `bson:"voting_start_time"`
	Proposer         string           `bson:"proposer"`
}

type Content struct {
	Type        string `bson:"type"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Changes     []struct {
		Key      string `bson:"key"`
		Value    string `bson:"value"`
		Subspace string `bson:"subspace"`
	}
}

type FinalTallyResult struct {
	Yes        string `bson:"yes"`
	Abstain    string `bson:"abstain"`
	No         string `bson:"no"`
	NoWithVeto string `bson:"no_with_veto"`
}

package document

import (
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"gopkg.in/mgo.v2/bson"
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
}

type Content struct {
	Type        string `bson:"type"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
}

type FinalTallyResult struct {
	Yes        string `bson:"yes"`
	Abstain    string `bson:"abstain"`
	No         string `bson:"no"`
	NoWithVeto string `bson:"no_with_veto"`
}

func (_ Proposal) GetAllProposalId() (map[int]int, error) {
	var proposals []Proposal
	var selector = bson.M{ProposalFieldProposalId: 1}
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(CollectionProposal).
		SetSelector(selector).
		SetResult(&proposals)

	err := query.Exec()

	var listProposalId map[int]int
	for _, proposal := range proposals {
		listProposalId[proposal.ProposalId] = proposal.ProposalId
	}
	return listProposalId, err
}

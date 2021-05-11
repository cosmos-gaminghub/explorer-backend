package exporter

import (
	"fmt"
	"strconv"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"gopkg.in/mgo.v2/bson"
)

// const (
// 	PROPOSAL_STATUS_UNSPECIFIED    = 0
// 	PROPOSAL_STATUS_DEPOSIT_PERIOD = 1
// 	PROPOSAL_STATUS_VOTING_PERIOD  = 2
// 	PROPOSAL_STATUS_PASSED         = 3
// 	PROPOSAL_STATUS_REJECTED       = 4
// 	PROPOSAL_STATUS_FAILED         = 5
// )

func GetProposals(ps types.ProposalResult) (proposals []*schema.Proposal, err error) {
	for _, proposal := range ps.Proposals {
		proposalId, err := strconv.Atoi(proposal.ProposalId)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		pro := &schema.Proposal{
			ProposalId:       proposalId,
			Proposer:         "",
			ProposalStatus:   proposal.Status,
			Content:          proposal.Content,
			SubmitTime:       proposal.SubmitTime,
			FinalTallyResult: proposal.FinalTallyResult,
			VotingEndTime:    proposal.VotingEndTime,
			VotingStartTime:  proposal.VotingStartTime,
		}
		proposals = append(proposals, pro)
	}

	return proposals, nil
}
func SaveProposal(proposal schema.Proposal) (interface{}, error) {
	selector := bson.M{"proposal_id": proposal.ProposalId}
	return orm.Upsert("proposal", selector, proposal)
}

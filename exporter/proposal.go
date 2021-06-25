package exporter

import (
	"fmt"
	"strconv"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
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
			ProposalStatus:   proposal.Status,
			Content:          proposal.Content,
			SubmitTime:       proposal.SubmitTime,
			FinalTallyResult: proposal.FinalTallyResult,
			VotingEndTime:    proposal.VotingEndTime,
			VotingStartTime:  proposal.VotingStartTime,
		}

		amounts := []schema.ProposalAmount{}
		for _, item := range proposal.TotalDeposit {
			a := schema.ProposalAmount{
				Denom:  item.Denom,
				Amount: item.Amount,
			}
			amounts = append(amounts, a)
		}
		pro.TotalDeposit = amounts
		proposals = append(proposals, pro)
	}

	return proposals, nil
}

func GetVote(voteResult []types.ProposalVote) (votes []*schema.Vote, err error) {
	for _, vote := range voteResult {
		proposalId, err := strconv.Atoi(vote.ProposalId)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		i, _ := document.Vote{}.QueryVoteDetailByProposalId(proposalId)
		if i.ProposalId != 0 {
			continue
		}

		v := &schema.Vote{
			ProposalId: proposalId,
			Voter:      vote.Voter,
			Option:     vote.Option,
		}
		votes = append(votes, v)
	}

	return votes, nil
}

func SaveProposal(proposal schema.Proposal) (interface{}, error) {
	selector := bson.M{document.ProposalFieldProposalId: proposal.ProposalId}
	return orm.Upsert(document.CollectionProposal, selector, proposal)
}

func GetDeposits(depositResult types.ProposalDepositResult) (deposits []*schema.Deposit, err error) {
	for _, deposit := range depositResult.Deposits {
		proposalId, err := strconv.Atoi(deposit.ProposalId)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		i, _ := document.Deposit{}.QueryDepositDetailByProposalId(proposalId)
		if i.ProposalId != 0 {
			continue
		}

		d := &schema.Deposit{
			ProposalId: proposalId,
			Depositor:  deposit.Depositor,
		}

		amounts := []schema.Amount{}
		for _, depositAmount := range deposit.Amount {
			a := schema.Amount{
				Denom:  depositAmount.Denom,
				Amount: depositAmount.Amount,
			}
			amounts = append(amounts, a)
		}
		d.Amount = amounts
		deposits = append(deposits, d)
	}

	return deposits, nil
}

package document

import "gopkg.in/mgo.v2/bson"

const (
	CollectionVote = "vote"
)

type Vote struct {
	ProposalId int    `bson:"proposal_id"`
	Voter      string `bson:"voter"`
	Option     string `bson:"option"`
}

func (_ Vote) QueryVoteDetailByProposalId(proposalId int) (Vote, error) {

	vote := Vote{}

	valCondition := bson.M{
		ProposalFieldProposalId: proposalId,
	}

	err := queryOne(CollectionVote, nil, valCondition, &vote)

	return vote, err
}

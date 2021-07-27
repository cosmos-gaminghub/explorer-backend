package document

import "gopkg.in/mgo.v2/bson"

const (
	CollectionDeposit = "deposit"
)

type Deposit struct {
	ProposalId int      `bson:"proposal_id"`
	Depositor  string   `bson:"depositor"`
	Amount     []Amount `bson:"amount"`
}

type Amount struct {
	Denom  string `bson:"denom"`
	Amount string `bson:"amount"`
}

func (_ Deposit) QueryDepositDetailByProposalId(proposalId int) (Deposit, error) {

	deposit := Deposit{}

	valCondition := bson.M{
		ProposalFieldProposalId: proposalId,
	}

	err := queryOne(CollectionDeposit, nil, valCondition, &deposit)

	return deposit, err
}

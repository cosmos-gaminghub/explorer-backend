package schema

import (
	"time"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
)

// Validator defines the structure for validator information.
type Validator struct {
	OperatorAddr    string            `bson:"operator_address"`
	ConsensusAddr   string            `bson:"consensus_address"`
	AccountAddr     string            `bson:"account_address"`
	Jailed          bool              `bson:"jailed"`
	Status          string            `bson:"status"`
	Tokens          string            `bson:"tokens" json:"tokens"`
	DelegatorShares string            `bson:"delegator_shares"`
	Description     types.Description `bson:"description" json:"description"`
	UnbondingHeight string            `bson:"unbonding_height"`
	UnbondingTime   time.Time         `bson:"unbonding_time"`
	Commission      types.Commission  `bson:"commission" json:"commission"`
	ProposerAddr    string            `bson:"proposer_addr"`
	Icons           string            `bson:"icons"`
}

// NewValidator returns a new Validator.
func NewValidator(v Validator) *Validator {
	return &Validator{
		OperatorAddr:    v.OperatorAddr,
		ConsensusAddr:   v.ConsensusAddr,
		AccountAddr:     v.AccountAddr,
		Jailed:          v.Jailed,
		Status:          v.Status,
		Tokens:          v.Tokens,
		DelegatorShares: v.DelegatorShares,
		Description:     v.Description,
		UnbondingHeight: v.UnbondingHeight,
		UnbondingTime:   v.UnbondingTime,
		Commission:      v.Commission,
	}
}

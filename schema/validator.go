package schema

import (
	"time"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
)

// Validator defines the structure for validator information.
type Validator struct {
	OperatorAddr     string            `bson:"operator_address"`
	ConsensusPubkey  string            `bson:"consensus_pubkey"`
	ConsensusAddres  string            `bson:"consensus_address"`
	AccountAddr      string            `bson:"account_address"`
	Jailed           bool              `bson:"jailed"`
	Status           string            `bson:"status"`
	Tokens           int64             `bson:"tokens" json:"tokens"`
	DelegatorShares  string            `bson:"delegator_shares"`
	Description      types.Description `bson:"description" json:"description"`
	UnbondingHeight  string            `bson:"unbonding_height"`
	UnbondingTime    time.Time         `bson:"unbonding_time"`
	Commission       types.Commission  `bson:"commission" json:"commission"`
	ProposerAddr     string            `bson:"proposer_addr"`
	Icons            string            `bson:"icons"`
	TotalMissedBlock int64             `bson:"total_missed_block"`
}

// NewValidator returns a new Validator.
func NewValidator(v Validator) *Validator {
	return &Validator{
		OperatorAddr:     v.OperatorAddr,
		ConsensusPubkey:  v.ConsensusPubkey,
		ConsensusAddres:  v.ConsensusAddres,
		AccountAddr:      v.AccountAddr,
		Jailed:           v.Jailed,
		Status:           v.Status,
		Tokens:           v.Tokens,
		DelegatorShares:  v.DelegatorShares,
		Description:      v.Description,
		UnbondingHeight:  v.UnbondingHeight,
		UnbondingTime:    v.UnbondingTime,
		Commission:       v.Commission,
		TotalMissedBlock: v.TotalMissedBlock,
	}
}

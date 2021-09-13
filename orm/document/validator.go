package document

import (
	"fmt"
	"time"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmValidator = "validator"

	ValidatorFieldVotingPower      = "voting_power"
	ValidatorFieldJailed           = "jailed"
	ValidatorFieldStatus           = "status"
	ValidatorFieldOperatorAddress  = "operator_address"
	ValidatorFieldDescription      = "description"
	ValidatorFieldConsensusAddr    = "consensus_pubkey"
	ValidatorFieldProposerHashAddr = "proposer_addr"
	ValidatorFieldTokens           = "tokens"
	ValidatorFieldDelegatorShares  = "delegator_shares"
	ValidatorFieldIcon             = "icons"
	ValidatorFieldComission        = "commission"
	ValidatorStatusValUnbonded     = 0
	ValidatorStatusValUnbonding    = 1
	ValidatorStatusValBonded       = 2
)

type Validator struct {
	OperatorAddr     string            `json:"operator_address"`
	ConsensusPubkey  string            `json:"consensus_pubkey"`
	ConsensusAddres  string            `json:"consensus_address"`
	AccountAddr      string            `json:"account_address"`
	Jailed           bool              `json:"jailed"`
	Status           string            `json:"status"`
	Tokens           int64             `json:"tokens"`
	DelegatorShares  string            `json:"delegator_shares"`
	Description      types.Description `json:"description"`
	UnbondingHeight  string            `json:"unbonding_height"`
	UnbondingTime    time.Time         `json:"unbonding_time"`
	Commission       types.Commission  `json:"commission"`
	ProposerAddr     string            `json:"proposer_addr"`
	Icons            string            `json:"icons"`
	TotalMissedBlock int64             `bson:"total_missed_block"`
}

func (v Validator) GetValidatorList() ([]Validator, error) {
	var validatorsDocArr []Validator
	var selector = bson.M{"description.moniker": 1, "operator_address": 1, "consensus_pubkey": 1, "proposer_addr": 1}
	err := queryAll(CollectionNmValidator, selector, nil, "", 0, &validatorsDocArr)

	return validatorsDocArr, err
}

func (v Validator) GetValidatorByProposerAddr(addr string) (Validator, error) {

	var selector = bson.M{"description.moniker": 1, "operator_address": 1}
	err := queryOne(CollectionNmValidator, selector, bson.M{"proposer_addr": addr}, &v)

	return v, err
}

type Description struct {
	Moniker  string `bson:"moniker" json:"moniker"`
	Identity string `bson:"identity" json:"identity"`
	Website  string `bson:"website" json:"website"`
	Details  string `bson:"details" json:"details"`
}

func (d Description) String() string {
	return fmt.Sprintf(`Moniker  :%v  Identity :%v Website  :%v Details  :%v`, d.Moniker, d.Identity, d.Website, d.Details)
}

func (v Validator) Name() string {
	return CollectionNmValidator
}

func (v Validator) QueryValidatorMonikerOpAddrConsensusPubkey(addrArrAsVa []string) ([]Validator, error) {
	var validators []Validator
	var selector = bson.M{
		"description.moniker": 1,
		"operator_address":    1,
		"consensus_pubkey":    1,
		"status":              1,
		"voting_power":        1,
	}

	err := queryAll(CollectionNmValidator, selector, bson.M{"operator_address": bson.M{"$in": addrArrAsVa}}, "", 0, &validators)
	return validators, err
}

func (v Validator) QueryValidatorsMonikerOpAddrConsensusPubkey() ([]Validator, error) {
	var validators []Validator
	var selector = bson.M{
		"description.moniker": 1,
		"operator_address":    1,
		"consensus_pubkey":    1,
		"status":              1,
		"voting_power":        1,
	}

	condition := bson.M{
		ValidatorFieldStatus: ValidatorStatusValBonded,
	}

	err := queryAll(CollectionNmValidator, selector, condition, "", 0, &validators)
	return validators, err
}

func (v Validator) QueryValidatorMonikerOpAddrByHashAddr(hashAddr []string) ([]Validator, error) {
	var validators []Validator
	var selector = bson.M{"description.moniker": 1, "operator_address": 1, "proposer_addr": 1}

	err := queryAll(CollectionNmValidator, selector, bson.M{"proposer_addr": bson.M{"$in": hashAddr}}, "", 0, &validators)
	return validators, err
}

func GetValidatorByAddr(addr string) (Validator, error) {
	db := getDb()
	c := db.C(CollectionNmValidator)
	defer db.Session.Close()
	var validator Validator
	err := c.Find(bson.M{ValidatorFieldOperatorAddress: addr}).One(&validator)

	return validator, err
}

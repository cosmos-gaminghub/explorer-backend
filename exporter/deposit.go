package exporter

import (
	"fmt"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"gopkg.in/mgo.v2/bson"
)

func SaveDeposit(deposit types.ProposalDeposit, ProposalId int) error {
	selector := bson.M{"proposal_id": ProposalId}
	_, err := orm.RemoveAll("deposit", selector)
	if err != nil {
		fmt.Printf("Error when remove all deposit of proposal %d", ProposalId)
	}
	return nil
}

package exporter

import (
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"gopkg.in/mgo.v2/bson"
)

// GetContract parses contract information and wrap into Precommit schema struct
func GetContract(wc types.WasmContract) *schema.Contract {
	result := wc.Result
	contract := schema.NewContract().
		SetCode(result.CodeId).
		SetLabel(result.Label).
		SetCreator(result.Creator).
		SetContractAddress(result.ContractAddress)

	return contract
}

func SaveContract(t *schema.Contract) (interface{}, error) {
	selector := bson.M{document.ContractAddress: t.ContractAddress}
	return orm.Upsert(document.CollectionContract, selector, t)
}

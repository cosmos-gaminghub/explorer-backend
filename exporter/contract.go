package exporter

import (
	"time"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
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
		SetContractAddress(result.ContractAddress).
		SetAdmin(result.Admin)

	return contract
}

func SaveContract(t *schema.Contract) (interface{}, error) {
	selector := bson.M{document.ContractAddressField: t.ContractAddress}
	return orm.Upsert(document.CollectionContract, selector, t)
}

func SaveContractInstantiateInfo(contractAddress string, txhash string, instantiatedAt time.Time) (interface{}, error) {
	contract, err := document.Contract{}.FindByContractAddress(contractAddress)
	if err != nil {
		logger.Error("failed to get contract from db:", logger.String("err", err.Error()))
	}
	contract.SetTxhash(txhash).
		SetInstantiatedAt(instantiatedAt).
		SetLastExecutedAt(instantiatedAt)
	return SaveContract(&contract)
}

func SaveContractExecuteInfo(contractAddress string, executeAt time.Time) (interface{}, error) {
	contract, err := document.Contract{}.FindByContractAddress(contractAddress)
	if err != nil {
		logger.Error("failed to get contract from db:", logger.String("err", err.Error()))
	}
	executeCount := contract.ExecutedCount + 1
	contract.SetLastExecutedAt(executeAt).
		SetExecutedCount(executeCount)

	return SaveContract(&contract)
}

func SaveContractAdminInfo(contractAddress string, admin string) (interface{}, error) {
	contract, err := document.Contract{}.FindByContractAddress(contractAddress)
	if err != nil {
		logger.Error("failed to get contract from db:", logger.String("err", err.Error()))
	}
	contract.SetAdmin(admin)
	return SaveContract(&contract)
}

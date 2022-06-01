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

// GetCodes parses codes information and wrap into Precommit schema struct
func GetCodes(wc types.WasmCode) (codes []*schema.Code) {
	for _, result := range wc.Result {
		code := schema.NewCode().
			SetCode(result.Id).
			SetCreator(result.Creator).
			SetDataHash(result.DataHash)

		codes = append(codes, code)
	}

	return codes
}

func SaveCode(t *schema.Code) (interface{}, error) {
	selector := bson.M{document.CodeIdField: t.CodeId}
	return orm.Upsert(document.CollectionCode, selector, t)
}

func SaveCodeInstantiateCount(codeId int) (interface{}, error) {
	code, err := document.Code{}.FindByCodeId(codeId)
	if err != nil {
		logger.Error("failed to get code from db:", logger.String("err", err.Error()))
	}
	instantiateCount := code.InstantiateCount + 1
	code.SetInstantiateCount(instantiateCount)
	return SaveCode(&code)
}

func SaveCodeMigrateInfo(codeId int, txhash string, createdAt time.Time) (interface{}, error) {
	code, err := document.Code{}.FindByCodeId(codeId)
	if err != nil {
		logger.Error("failed to get code from db:", logger.String("err", err.Error()))
	}
	code.SetCreatedAt(createdAt).SetTxhash(txhash)
	return SaveCode(&code)
}

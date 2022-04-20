package exporter

import (
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
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
	selector := bson.M{document.CodeId: t.CodeId}
	return orm.Upsert(document.CollectionCode, selector, t)
}

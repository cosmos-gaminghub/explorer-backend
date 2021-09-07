package exporter

import (
	"strings"

	"github.com/cosmos-gaminghub/explorer-backend/conf"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getValidators parses validators information and wrap into Precommit schema struct
func SaveAccountTransaction(validators []*schema.Validator, transactions []*schema.Transaction) {

	for _, tx := range transactions {
		var listAccountAddress = getListAccountAddres(tx.Messages)
		for _, address := range listAccountAddress {
			orm.Save(document.CollectionAccountTransaction, &document.AccountTransaction{
				Height:      tx.Height,
				AccountAddr: address,
				TxHash:      tx.TxHash,
			})
		}
	}
}

func getListAccountAddres(messages string) []string {
	var list []string
	var addressPrefix = conf.Get().Db.AddresPrefix
	for {
		if strings.Contains(messages, addressPrefix) {
			index := strings.Index(messages, addressPrefix)
			address := utils.Convert(addressPrefix, messages[index:index+45])
			if address != "" {
				list = append(list, address)
			}
			messages = messages[index+45 : len(messages)-1]
		} else {
			break
		}
	}
	return list
}

package exporter

import (
	"regexp"

	"github.com/cosmos-gaminghub/explorer-backend/conf"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

// getValidators parses validators information and wrap into Precommit schema struct
func SaveAccountTransaction(validators []*schema.Validator, transactions []*schema.Transaction) {

	for _, tx := range transactions {
		var listAccountAddress = getListAccountAddress(tx.Messages)
		for _, address := range listAccountAddress {
			orm.Save(document.CollectionAccountTransaction, &document.AccountTransaction{
				Height:      tx.Height,
				AccountAddr: address,
				TxHash:      tx.TxHash,
			})
		}
	}
}

func getListAccountAddress(messages string) []string {
	var list []string
	var addressPrefix = conf.Get().Db.AddresPrefix
	var re = regexp.MustCompile(`(?m)\"juno.{39}(\\")`)

	for _, match := range re.FindAllString(messages, -1) {
		// address have format "address\" --> correct address = address[1:len(address)-2]
		address := utils.Convert(addressPrefix, match)
		if address != "" {
			list = append(list, address[1:len(address)-2])
		}
	}
	return list
}

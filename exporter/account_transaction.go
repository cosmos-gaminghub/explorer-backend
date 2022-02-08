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
	var length = len(addressPrefix) + 39 // length off account address
	for {
		if strings.Contains(messages, addressPrefix) {
			index := strings.Index(messages, addressPrefix)
			var nextIndex = index + length
			if nextIndex > len(messages)-1 {
				break
			}

			if string(messages[index+length]) != "\"" { //skip string if got string like ugame,...
				messages = messages[index+1 : len(messages)-1]
				continue
			}
			var messageAddress = messages[index : index+length]
			if !(strings.Contains(messageAddress, addressPrefix+"valoper")) { // check if string is gamevaloper
				address := utils.Convert(addressPrefix, messageAddress)
				if address != "" {
					list = append(list, address)
				}
			}
			messages = messages[index+length : len(messages)-1]
		} else {
			break
		}
	}
	return list
}

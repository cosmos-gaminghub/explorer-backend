package exporter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
)

// getTxs parses transactions in a block and return transactions.
func GetTxs(txs types.TxResult, block schema.Block) (transactions []*schema.Transaction, err error) {

	if len(txs.TxResponse) <= 0 {
		return []*schema.Transaction{}, nil
	}

	for index, tx := range txs.TxResponse {
		height, _ := strconv.ParseInt(tx.Height, 10, 64)
		gasWanted, _ := strconv.ParseInt(tx.GasWanted, 10, 64)
		gasUsed, _ := strconv.ParseInt(tx.GasUsed, 10, 64)
		messages, _ := json.Marshal(txs.Txs[index].Body.BodyMessage)

		t := schema.NewTransaction(schema.Transaction{
			Height:     height,
			TxHash:     tx.TxHash,
			Code:       tx.Code,
			Memo:       txs.Txs[index].Body.Memo,
			GasWanted:  gasWanted,
			GasUsed:    gasUsed,
			Timestamp:  block.Timestamp,
			Logs:       tx.Logs,
			Fee:        txs.Txs[index].AuthInfo.FeeInfo,
			Signatures: txs.Txs[index].Signatures,
			Messages:   string(messages),
			RawLog:     tx.RawLog,
		})
		transactions = append(transactions, t)

		saveWasmInfo(tx.Logs, tx.TxHash, block.Timestamp)
	}

	return transactions, nil
}

func saveWasmInfo(logs []types.Log, txhash string, time time.Time) (result interface{}, err error) {
	for _, log := range logs {
		for _, event := range log.Events {
			switch event.Type {
			case "execute":
				{
					for _, attribute := range event.Attributes {
						if attribute.Key == "execute" {
							SaveContractExecuteInfo(attribute.Value, time)
						}
					}
					break
				}
			case "update_admin":
				{
					for _, attribute := range event.Attributes {
						var contractAddress string
						var admin string
						if attribute.Key == "_contract_addr" {
							contractAddress = attribute.Value
						}

						if attribute.Key == "admin" {
							admin = attribute.Value
						}
						SaveContractAdminInfo(contractAddress, admin)
					}
					break
				}
			case "clear_admin":
				{
					for _, attribute := range event.Attributes {
						var contractAddress string
						var admin string
						if attribute.Key == "_contract_addr" {
							contractAddress = attribute.Value
						}
						SaveContractAdminInfo(contractAddress, admin)
					}
					break
				}
			case "instantiate":
				{
					for _, attribute := range event.Attributes {
						if attribute.Key == "_contract_address" {
							SaveContractInstantiateInfo(attribute.Value, txhash, time)
						}
						if attribute.Key == "code_id" {
							codeId, err := strconv.Atoi(attribute.Value)
							if err != nil {
								logger.Error(fmt.Sprintf("Failed to parse code id %s to int", attribute.Value))
								continue
							}
							SaveCodeInfo(codeId, txhash, time)
						}
					}
					break
				}
			default:
				{
					break
				}
			}
		}
	}

	return nil, nil
}

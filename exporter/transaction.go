package exporter

import (
	"strconv"

	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
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
			Messages:   txs.Txs[index].Body.BodyMessage,
		})
		transactions = append(transactions, t)
	}

	return transactions, nil
}

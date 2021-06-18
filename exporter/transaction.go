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
		// var stdTx txtypes.StdTx
		// err = ex.cdc.UnmarshalBinaryLengthPrefixed([]byte(tx.Tx), &stdTx)
		// if err != nil {
		// 	return []*schema.Transaction{}, err
		// }

		// msgsBz, err := ex.cdc.MarshalJSON(stdTx.GetMsgs())
		// if err != nil {
		// 	return []*schema.Transaction{}, err
		// }

		// sigs := make([]types.Signature, len(stdTx.Signatures), len(stdTx.Signatures))

		// for i, sig := range stdTx.Signatures {
		// 	consPubKey, err := ctypes.Bech32ifyConsPub(sig.PubKey)
		// 	if err != nil {
		// 		return []*schema.Transaction{}, err
		// 	}

		// 	sigs[i] = types.Signature{
		// 		Address:       sig.Address().String(), // hex string
		// 		AccountNumber: sig.AccountNumber,
		// 		Pubkey:        consPubKey,
		// 		Sequence:      sig.Sequence,
		// 		Signature:     base64.StdEncoding.EncodeToString(sig.Signature), // encode base64
		// 	}
		// }

		// sigsBz, err := ex.cdc.MarshalJSON(sigs)
		// if err != nil {
		// 	return []*schema.Transaction{}, err
		// }

		height, err := strconv.ParseInt(tx.Height, 10, 64)
		if err != nil {
			return []*schema.Transaction{}, err
		}

		gasWanted, err := strconv.ParseInt(tx.GasWanted, 10, 64)
		if err != nil {
			return []*schema.Transaction{}, err
		}

		gasUsed, err := strconv.ParseInt(tx.GasUsed, 10, 64)
		if err != nil {
			return []*schema.Transaction{}, err
		}
		t := schema.Transaction{
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
		}

		transactions = append(transactions, &t)
	}

	return transactions, nil
}

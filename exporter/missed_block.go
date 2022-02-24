package exporter

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos-gaminghub/explorer-backend/conf"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

// SaveMissedBlock by height
func SaveMissedBlock(clientHTTP *rpchttp.HTTP, height int64, block *schema.Block) {
	result, err := clientHTTP.BlockResults(context.Background(), &height)
	if err != nil {
		fmt.Println(err.Error())
	}
	var addressPrefix = conf.Get().Db.AddresPrefix
	for _, event := range result.BeginBlockEvents {
		if event.Type == "liveness" {
			var consensusAddress, eventHeight, operatorAddress string
			var insertHeight int64
			for _, v := range event.GetAttributes() {

				var ak = bytes.NewBuffer(v.GetKey()).String()
				if ak == "address" {
					consensusAddress = bytes.NewBuffer(v.GetValue()).String()
					operatorAddress = utils.Convert(addressPrefix+"valoper", consensusAddress)
				}

				if ak == "height" {
					eventHeight = bytes.NewBuffer(v.GetValue()).String()
					insertHeight, _ = strconv.ParseInt(eventHeight, 10, 64)
					if err != nil {
						logger.Error(fmt.Sprintf("[Missed block] failed to parse string %s to int64", eventHeight))
					}
				}
				if insertHeight > 0 {
					b := schema.NewMissedBlock(schema.MissedBlock{
						Height:       insertHeight,
						OperatorAddr: operatorAddress,
						Timestamp:    block.Timestamp,
					})
					orm.Save("missed_block", b)
				}
			}
		}
	}
}

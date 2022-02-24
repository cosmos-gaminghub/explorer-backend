package exporter

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

// SaveMissedBlock by height
func SaveMissedBlock(clientHTTP *rpchttp.HTTP, height int64, block *schema.Block) {
	result, err := clientHTTP.BlockResults(context.Background(), &height)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, event := range result.BeginBlockEvents {
		if event.Type == "liveness" {
			var consensusAddress, eventHeight string
			var insertHeight int64
			for _, v := range event.GetAttributes() {

				var ak = bytes.NewBuffer(v.GetKey()).String()
				if ak == "address" {
					consensusAddress = bytes.NewBuffer(v.GetValue()).String()
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
						Height:           insertHeight,
						ConsensusAddress: consensusAddress,
						Timestamp:        block.Timestamp,
					})
					orm.Save("missed_block", b)
				}
			}
		}
	}
}

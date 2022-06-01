package exporter

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos-gaminghub/explorer-backend/conf"
	types "github.com/cosmos-gaminghub/explorer-backend/lcd"
	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/orm/document"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

var (
	Key = "636F6E74726163745F696E666F"
)

// GetContract parses contract information and wrap into Precommit schema struct
func GetContract(wc types.WasmContract, contractState types.WasmRawContractState) *schema.Contract {
	result := wc.Result
	contract := schema.NewContract().
		SetCode(result.CodeId).
		SetLabel(result.Label).
		SetCreator(result.Creator).
		SetContractAddress(result.ContractAddress).
		SetAdmin(result.Admin)
	if contractState != (types.WasmRawContractState{}) {
		contract = contract.SetContract(contractState.Contract).SetVersion(contractState.Version)
	}

	return contract
}

func SaveContract(t *schema.Contract) (interface{}, error) {
	selector := bson.M{document.ContractAddressField: t.ContractAddress}
	return orm.Upsert(document.CollectionContract, selector, t)
}

func SaveContractInstantiateInfo(contractAddress string, txhash string, instantiatedAt time.Time) (interface{}, error) {
	contract, err := document.Contract{}.FindByContractAddress(contractAddress)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get contract %s from db:", contractAddress), logger.String("err", err.Error()))
		return nil, err
	}
	contract.SetTxhash(txhash).
		SetInstantiatedAt(instantiatedAt).
		SetLastExecutedAt(instantiatedAt)

	code, err := document.Code{}.FindByCodeId(contract.CodeId)
	if err == nil {
		if instantiatedAt.After(code.FirstContractTime) {
			code.SetFirstContractTime(instantiatedAt).
				SetContract(contract.Contract).
				SetVersion(contract.Version)
		}
		_, err = SaveCode(&code)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to save code %d verson and contract:", contract.CodeId), logger.String("err", err.Error()))
		}
	} else {
		logger.Error(fmt.Sprintf("failed to get code %d from db: ", contract.CodeId), logger.String("err", err.Error()))
	}

	return SaveContract(&contract)
}

func SaveContractExecuteInfo(contractAddress string, executeAt time.Time) (interface{}, error) {
	contract, err := document.Contract{}.FindByContractAddress(contractAddress)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get contract %s from db:", contractAddress), logger.String("err", err.Error()))
		return nil, err
	}
	executeCount := contract.ExecutedCount + 1
	contract.SetLastExecutedAt(executeAt).
		SetExecutedCount(executeCount)

	return SaveContract(&contract)
}

func SaveContractAdminInfo(contractAddress string, admin string) (interface{}, error) {
	contract, err := document.Contract{}.FindByContractAddress(contractAddress)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get contract %s from db:", contractAddress), logger.String("err", err.Error()))
		return nil, err
	}
	contract.SetAdmin(admin)
	return SaveContract(&contract)
}

type argumentDecoder struct {
	// dec is the default decoder
	dec                func(string) ([]byte, error)
	asciiF, hexF, b64F bool
}

func newArgDecoder(def func(string) ([]byte, error)) *argumentDecoder {
	return &argumentDecoder{dec: def}
}

func (a *argumentDecoder) RegisterFlags(f *flag.FlagSet, argName string) {
	f.BoolVar(&a.asciiF, "ascii", false, "ascii encoded "+argName)
	f.BoolVar(&a.hexF, "hex", false, "hex encoded  "+argName)
	f.BoolVar(&a.b64F, "b64", false, "base64 encoded "+argName)
}

func (a *argumentDecoder) DecodeString(s string) ([]byte, error) {
	found := -1
	for i, v := range []*bool{&a.asciiF, &a.hexF, &a.b64F} {
		if !*v {
			continue
		}
		if found != -1 {
			return nil, errors.New("multiple decoding flags used")
		}
		found = i
	}
	switch found {
	case 0:
		return asciiDecodeString(s)
	case 1:
		return hex.DecodeString(s)
	case 2:
		return base64.StdEncoding.DecodeString(s)
	default:
		return a.dec(s)
	}
}

func asciiDecodeString(s string) ([]byte, error) {
	return []byte(s), nil
}

func GetRawContractState(contractAddress string) (types.WasmRawContractState, error) {
	grpcConn, err := grpc.Dial(
		conf.Get().Hub.GrpcUrl,
		grpc.WithInsecure(),
	)
	defer grpcConn.Close()

	if err != nil {
		logger.Error("[Get Raw Contract State] Can not connect grpc", logger.String("err", err.Error()))
		return types.WasmRawContractState{}, err
	}

	decoder := newArgDecoder(hex.DecodeString)
	queryData, err := decoder.DecodeString(Key)
	if err != nil {
		logger.Error("[Get Raw Contract State] Can not decode key", logger.String("err", err.Error()))
		return types.WasmRawContractState{}, err
	}

	queryClient := wasmtypes.NewQueryClient(grpcConn)
	res, err := queryClient.RawContractState(
		context.Background(),
		&wasmtypes.QueryRawContractStateRequest{
			Address:   contractAddress,
			QueryData: queryData,
		},
	)
	if err != nil {
		logger.Error(fmt.Sprintf("[Get Raw Contract State] Can not get contract state %s", contractAddress), logger.String("err", err.Error()))
		return types.WasmRawContractState{}, err
	}

	var result types.WasmRawContractState
	if err := json.Unmarshal(res.Data, &result); err != nil {
		logger.Error(fmt.Sprintf("[Get Raw Contract State] Unmarshal wasm contract state %s", contractAddress), logger.String("err", err.Error()))
		return types.WasmRawContractState{}, err
	}
	return result, nil
}

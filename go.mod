module github.com/cosmos-gaminghub/explorer-backend

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/etcd-io/bbolt v1.3.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/tendermint/tendermint v0.34.10 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc // indirect
	google.golang.org/grpc v1.37.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)

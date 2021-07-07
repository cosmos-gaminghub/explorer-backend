package schema

import (
	"time"
)

// StatAssetInfoList1H defines the schema for asset statistics in an hourly basis
type StatAssetInfoList20Minute struct {
	Price     float64   `bson:"price"`
	Marketcap float64   `bson:"market_cap"`
	Volume24H float64   `bson:"volumne_24h"`
	Timestamp time.Time `bson:"timestamp"`
}

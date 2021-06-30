package schema

// StatAssetInfoList1H defines the schema for asset statistics in an hourly basis
type StatAssetInfoList20Minute struct {
	Price     float32 `bson:"price"`
	Marketcap float32 `bson:"market_cap"`
	Volume24H float32 `bson:"volumne_24h"`
}

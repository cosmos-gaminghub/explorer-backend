package schema

import "time"

var ChosenAssetNames = []string{
	"TUSDB-888",
	"USDSB-1AC",
	"BTCB-1DE",
	"IRIS-D88",
}

// StatAssetInfoList1H defines the schema for asset statistics in an hourly basis
type StatAssetInfoList1H struct {
	ID                int32     `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Slug              string    `json:"slug"`
	NumMarketPairs    string    `json:"num_market_pairs"`
	DateAdded         time.Time `json:"date_added"`
	MaxSupply         string    `json:"max_supply"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	IsActive          uint      `json:"is_active"`
	Platform          string    `json:"platform"`
	LastUpdated       time.Time `json:"last_updated"`
	Price             float64   `json:"price"`
	Currency          string    `json:"currency"`
	ChangeRange       float64   `json:"change_range"`
	Supply            float64   `json:"supply"`
	Marketcap         float64   `json:"marketcap"`
	AssetImg          string    `json:"asset_img"`
	AssetCreateTime   int64     `json:"asset_create_time"`
	Timestamp         time.Time `json:"timestamp"`
}

// StatAssetInfoList24H defines the schema for asset statistics in 24 hourly basis
type StatAssetInfoList24H struct {
	ID              int32     `json:"id"`
	TotalNum        int       `json:"total_num"`
	Name            string    `json:"name"`
	Asset           string    `json:"asset"`
	Owner           string    `json:"owner"`
	Price           float64   `json:"price"`
	Currency        string    `json:"currency"`
	ChangeRange     float64   `json:"change_range"`
	Supply          float64   `json:"supply"`
	Marketcap       float64   `json:"marketcap"`
	AssetImg        string    `json:"asset_img"`
	AssetCreateTime int64     `json:"asset_create_time"`
	Timestamp       time.Time `json:"timestamp"`
}

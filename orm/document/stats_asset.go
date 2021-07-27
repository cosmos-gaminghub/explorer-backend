package document

import (
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/orm"
)

const (
	CollectionNmStatsAsset = "stats_asset"

	StatsAsset_Field_Time = "timestamp"
)

// StatAssetInfoList1H defines the schema for asset statistics in an hourly basis
type StatAssetInfoList20Minute struct {
	Price     float64   `bson:"price"`
	Marketcap float64   `bson:"market_cap"`
	Volume24H float64   `bson:"volumne_24h"`
	Timestamp time.Time `bson:"timestamp"`
}

func (_ StatAssetInfoList20Minute) QueryLatestStatAssetFromDB() (StatAssetInfoList20Minute, error) {

	var statsAssets StatAssetInfoList20Minute

	sort := desc(StatsAsset_Field_Time)
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(CollectionNmStatsAsset).
		SetCondition(nil).
		SetSort(sort).
		SetResult(&statsAssets)

	err := query.Exec()
	if err == nil {
		return statsAssets, nil
	}

	return StatAssetInfoList20Minute{}, err
}

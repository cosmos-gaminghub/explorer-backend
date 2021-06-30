package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/conf"
	"github.com/cosmos-gaminghub/explorer-backend/orm"
	"github.com/cosmos-gaminghub/explorer-backend/schema"
)

const (
	ConcurrencyQuoteLast = "coins/%s/market_chart/range"
)

type CoinsIDMarketChart struct {
	Prices       []ChartItem `json:"prices"`
	MarketCaps   []ChartItem `json:"market_caps"`
	TotalVolumes []ChartItem `json:"total_volumes"`
}

type ChartItem [2]float32

func Get(uri string, mapQuery map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", conf.Get().Coingecko.EndPoint+"/"+conf.Get().Coingecko.Version+"/"+uri, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	for key, value := range mapQuery {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bz, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return bz, nil
	}
	return nil, nil
}

func SaveMarketChartRange(coin string, mintue int64) (err error) {
	query := make(map[string]string)
	currentTime := time.Now().Unix()
	query["from"] = strconv.FormatInt(currentTime-mintue*60, 10)
	query["to"] = strconv.FormatInt(currentTime, 10)
	query["vs_currency"] = conf.Get().Coingecko.Currency

	uri := fmt.Sprintf(ConcurrencyQuoteLast, coin)
	resBytes, err := Get(uri, query)
	if err != nil {
		log.Fatalln("Get oinmarket get currency quote lastest error")
		return err
	}

	var data CoinsIDMarketChart
	if err := json.Unmarshal(resBytes, &data); err != nil {
		log.Fatalln("Unmarshal coinmarket get currency quote lastest error")
		return err
	}
	for key, item := range data.Prices {
		t := schema.StatAssetInfoList20Minute{
			Price:     item[1],
			Marketcap: data.MarketCaps[key][1],
			Volume24H: data.TotalVolumes[key][1],
		}
		orm.Save("stats_asset", t)
	}
	return err
}
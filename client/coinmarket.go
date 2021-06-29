package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cosmos-gaminghub/explorer-backend/conf"
)

const (
	ConcurrencyQuoteLast = "/cryptocurrency/quotes/latest"
)

type CryptocurrencyLatestQuotes map[string]struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	CirculatingSupply float64     `json:"circulating_supply"`
	TotalSupply       float64     `json:"total_supply"`
	MaxSupply         float64     `json:"max_supply"`
	DateAdded         time.Time   `json:"date_added"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	CmcRank           int         `json:"cmc_rank"`
	LastUpdated       time.Time   `json:"last_updated"`
	Tags              []string    `json:"tags"`
	Platform          interface{} `json:"platform"`
	Quote             struct {
		USD Currency `json:"USD"`
	} `json:"quote"`
}

type ExchangeHistoricalListings struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	CmcRank        int       `json:"cmc_rank"`
	NumMarketPairs int       `json:"num_market_pairs"`
	Timestamp      time.Time `json:"timestamp"`
	Quote          struct {
		USD Currency `json:"USD"`
	} `json:"quote"`
}

type Currency struct {
	Price                  float64   `json:"price"`
	Volume24H              float64   `json:"volume_24h"`
	Volume24HAdjusted      float64   `json:"volume_24h_adjusted"`
	Volume7D               float64   `json:"volume_7d"`
	Volume30D              float64   `json:"volume_30d"`
	PercentChange1H        float64   `json:"percent_change_1h"`
	PercentChangeVolume24H float64   `json:"percent_change_24h"`
	PercentChangeVolume7D  float64   `json:"percent_change_7d"`
	PercentChangeVolume30D float64   `json:"percent_change_30d"`
	PercentChangeVolume60D float64   `json:"percent_change_60d"`
	PercentChangeVolume90D float64   `json:"percent_change_90d"`
	MarketCap              float64   `json:"market_cap"`
	TotalMarketCap         float64   `json:"total_market_cap"`
	LastUpdated            time.Time `json:"last_updated"`
}

type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int64  `json:"elapsed"`
	CreditCount  int64  `json:"credit_count"`
}

func Get(uri string, mapQuery map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", conf.Get().CoinMarket.EndPoint+"/"+conf.Get().CoinMarket.Version+"/"+uri, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	for key, value := range mapQuery {
		q.Add(key, value)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", conf.Get().CoinMarket.Apikey)
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

// func GetConcurrencyQuoteLastest(slug string) (price *model.Price, err error) {
// 	query := make(map[string]string)
// 	query["slug"] = "cosmos"
// 	resBytes, err := Get(ConcurrencyQuoteLast, query)
// 	if err != nil {
// 		log.Fatalln("Get oinmarket get currency quote lastest error")
// 		return price, err
// 	}

// 	resp := struct {
// 		Data   CryptocurrencyLatestQuotes `json:"data"`
// 		Status Status                     `json:"status"`
// 	}{}

// 	if err := json.Unmarshal(resBytes, &resp); err != nil {
// 		log.Fatalln("Unmarshal coinmarket get currency quote lastest error")
// 		return price, err
// 	}

// 	if resp.Status.ErrorCode != 0 {
// 		return price, errors.New(resp.Status.ErrorMessage)
// 	}

// 	for _, item := range resp.Data {
// 		fmt.Println(item.Quote.USD)
// 		price = &model.Price{
// 			Volume24h:        utils.ParseStringFromFloat64(item.Quote.USD.Volume24H),
// 			MarketCap:        utils.ParseStringFromFloat64(item.Quote.USD.MarketCap),
// 			Price:            utils.ParseStringFromFloat64(item.Quote.USD.Price),
// 			PercentChange24h: utils.ParseStringFromFloat64(item.Quote.USD.PercentChangeVolume24H),
// 		}
// 	}

// 	return price, nil
// }

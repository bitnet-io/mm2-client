package external_services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/helpers"
	"net/http"
	"sync"
	"time"

	"sort"
//	"strings"

//	"os"

	"github.com/kpango/glg"
)

type CoinpaprikaAnswer struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Rank              int       `json:"rank"`
	CirculatingSupply int64     `json:"circulating_supply"`
	TotalSupply       int64     `json:"total_supply"`
	MaxSupply         int64     `json:"max_supply"`
	BetaValue         float64   `json:"beta_value"`
	FirstDataAt       time.Time `json:"first_data_at"`
	LastUpdated       string    `json:"last_updated"`
	Quotes            struct {
		USD struct {
			Price               float64   `json:"price"`
			Volume24H           float64   `json:"volume_24h"`
			Volume24HChange24H  float64   `json:"volume_24h_change_24h"`
			MarketCap           int64     `json:"market_cap"`
			MarketCapChange24H  float64   `json:"market_cap_change_24h"`
			PercentChange15M    float64   `json:"percent_change_15m"`
			PercentChange30M    float64   `json:"percent_change_30m"`
			PercentChange1H     float64   `json:"percent_change_1h"`
			PercentChange6H     float64   `json:"percent_change_6h"`
			PercentChange12H    float64   `json:"percent_change_12h"`
			PercentChange24H    float64   `json:"percent_change_24h"`
			PercentChange7D     float64   `json:"percent_change_7d"`
			PercentChange30D    float64   `json:"percent_change_30d"`
			PercentChange1Y     float64   `json:"percent_change_1y"`
			AthPrice            float64   `json:"ath_price"`
			AthDate             time.Time `json:"ath_date"`
			PercentFromPriceAth float64   `json:"percent_from_price_ath"`
		} `json:"USD"`
	} `json:"quotes"`
}

const gPaprikaEndpoint = "https://api.coinpaprika.com/v1/tickers"

var CoinpaprikaRegistry sync.Map


func NewPaprikaRequest(page int) string {
        url := gPaprikaEndpoint
  //      all_coins := getPaprikaCoinsList()
    //    url += strings.Join(all_coins, ",")
//        url += "&x_cg_api_key=CG-8XDVCpAhU2YLcD3EGAU4bzCJ"
//      url += "&order=id_asc&price_change_percentage=24h&sparkline=true&per_page=250"
//      url += "&page=" + fmt.Sprintf("%d", page)
        return url
}


func processCoinpaprika() *[]CoinpaprikaAnswer {
	url := gPaprikaEndpoint
	glg.Infof("Processing coinpaprika request: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}



/*


                if resp.StatusCode == http.StatusOK {
                        defer resp.Body.Close()
//                      var page_answer = &[]CoingeckoAnswer{}

                             var resulta CoinpaprikaAnswer
                       bodyBytes, _ := ioutil.ReadAll(resp.Body)

//                      decodeErr := json.NewDecoder(resp.Body).Decode(page_answer)
//                      if decodeErr != nil {
//                              fmt.Printf("decodeErr: %v\n", decodeErr)
//                      }
//                      fmt.Printf("Got %v coins form Gecko\n", len(*page_answer))
//                       *answer = append(*answer, *page_answer...)
//
//                      if len(*page_answer) == 0 || len(*page_answer) < 250 {
//                              return answer
//                      }

                      decodeErr2 := json.Unmarshal(bodyBytes, &resulta)

                        if decodeErr2 != nil {
                                fmt.Printf("decodeErr2: %v\n", decodeErr2)
                        }
                        fmt.Printf("Got %+v coins from paprika\n", resulta)



 rankingsJson, _ := json.Marshal(resulta)
    err = ioutil.WriteFile("./stocks/crypto.json", rankingsJson, 0644)
    fmt.Printf("%+v", resulta)


        //Open our jsonFile
    jsonFile, err := os.Open("./stocks/crypto.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened crypto.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()
*/







	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &[]CoinpaprikaAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", decodeErr)
			return nil
		}

		return nil

	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		glg.Errorf("Http status not OK: %s", bodyBytes)
		return nil
	}
}

func StartCoinpaprikaService() {
	functorVerification := func(ticker string, answer CoinpaprikaAnswer) {
		if val, ok := config.GCFGRegistry[ticker]; ok && val.CoinpaprikaID == answer.Id {
			ticker = helpers.RetrieveMainTicker(ticker)
			CoinpaprikaRegistry.Store(ticker, answer)
		}
	}
	for {
		if resp := processCoinpaprika(); resp != nil {
			glg.Info("Coinpaprika request successfully processed")
			for _, cur := range *resp {
				functorVerification(cur.Symbol, cur)
				functorVerification(cur.Symbol+"-v2", cur)
				functorVerification(cur.Symbol+"-ERC20", cur)
				functorVerification(cur.Symbol+"-PLG20", cur)
				functorVerification(cur.Symbol+"-AVX20", cur)
				functorVerification(cur.Symbol+"-FTM20", cur)
				functorVerification(cur.Symbol+"-KRC20", cur)
				functorVerification(cur.Symbol+"-HRC20", cur)
				functorVerification(cur.Symbol+"-BEP20", cur)
				functorVerification(cur.Symbol+"-QRC20", cur)
			}
		} else {
//			glg.Error("Something went wrong when processing coinpaprika request")
		}
		time.Sleep(constants.GPricesLoopTime)
	}
}

func CoinpaprikaRetrieveUSDValIfSupported(coin string) (string, string, string) {
	coin = helpers.RetrieveMainTicker(coin)
	val, ok := CoinpaprikaRegistry.Load(coin)
	valStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if ok {
		resp := val.(CoinpaprikaAnswer)
		valStr = fmt.Sprintf("%f", resp.Quotes.USD.Price)
		dateStr = resp.LastUpdated
	}
	return valStr, dateStr, "coinpaprika"
}

func CoinpaprikaRetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	basePrice, baseDate, _ := CoinpaprikaRetrieveUSDValIfSupported(base)
	relPrice, relDate, _ := CoinpaprikaRetrieveUSDValIfSupported(rel)
	price := helpers.BigFloatDivide(basePrice, relPrice, 8)
	date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if helpers.RFC3339ToTimestamp(baseDate) <= helpers.RFC3339ToTimestamp(relDate) {
		date = baseDate
	} else {
		date = relDate
	}
	return price, true, date, "coinpaprika"
}

func CoinpaprikaTotalVolume(coin string) (string, string, string) {
	coin = helpers.RetrieveMainTicker(coin)
	val, ok := CoinpaprikaRegistry.Load(coin)
	totalVolumeStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if ok {
		resp := val.(CoinpaprikaAnswer)
		totalVolumeStr = fmt.Sprintf("%f", resp.Quotes.USD.Volume24H)
		dateStr = resp.LastUpdated
	}
	return totalVolumeStr, dateStr, "coinpaprika"
}

func CoinpaprikaGetChange24h(coin string) (string, string, string) {
	changePercent24h := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if _, cfgOk := config.GCFGRegistry[coin]; cfgOk {
		val, ok := CoinpaprikaRegistry.Load(coin)
		if !ok {
			val, ok = CoinpaprikaRegistry.Load(helpers.RetrieveMainTicker(coin))
		}
		if ok {
			resp := val.(CoinpaprikaAnswer)
			changePercent24h = fmt.Sprintf("%f", resp.Quotes.USD.PercentChange24H)
			dateStr = resp.LastUpdated
		}
		return changePercent24h, dateStr, "coinpaprika"
	}
	return changePercent24h, dateStr, "unknown"
}

func getPaprikaCoinsList() []string {
	coins := []string{}
	for _, cur := range config.GCFGRegistry {
		if cur.CoinpaprikaID != "test-coin" && cur.CoinpaprikaID != "" {
			coins = append(coins, cur.CoinpaprikaID)
		}
	}
	coins = helpers.UniqueStrings(coins)
	sort.Strings(coins)
	return coins
}

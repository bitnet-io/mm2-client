package external_services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/helpers"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
	"github.com/kpango/glg"


	"os"


)








type CoingeckoSparkLineData struct {
	Price *[]float64 `json:"price,omitempty"`
}

/*
type CoingeckoAnswer struct {
type CoingeckoAnswer struct {
	Id                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	Image                        string    `json:"image"`
	CurrentPrice                 float64   `json:"current_price"`
	MarketCap                    float64   `json:"market_cap"`
	MarketCapRank                *int      `json:"market_cap_rank"`
	FullyDilutedValuation        *float64  `json:"fully_diluted_valuation"`
	TotalVolume                  float64   `json:"total_volume"`
	High24H                      *float64  `json:"high_24h"`
	Low24H                       *float64  `json:"low_24h"`
	PriceChange24H               *float64  `json:"price_change_24h"`
	PriceChangePercentage24H     *float64  `json:"price_change_percentage_24h"`
	MarketCapChange24H           *float64  `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H *float64  `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64   `json:"circulating_supply"`
	TotalSupply                  *float64  `json:"total_supply"`
	MaxSupply                    *float64  `json:"max_supply,omitempty"`
	Ath                          float64   `json:"ath"`
	AthChangePercentage          float64   `json:"ath_change_percentage"`
	AthDate                      time.Time `json:"ath_date"`
	Atl                          float64   `json:"atl"`
	AtlChangePercentage          float64   `json:"atl_change_percentage"`
	AtlDate                      time.Time `json:"atl_date"`
	Roi                          *struct {
		Times      float64 `json:"times"`
		Currency   string  `json:"currency"`
		Percentage float64 `json:"percentage"`
	} `json:"roi"`
	LastUpdated                        string                  `json:"last_updated"`
	SparklineIn7D                      *CoingeckoSparkLineData `json:"sparkline_in_7d,omitempty"`
	PriceChangePercentage24HInCurrency *float64                `json:"price_change_percentage_24h_in_currency"`

}
*/



/*


type CoingeckoAnswer struct {
	ID                      string `json:"_id"`
	Ticker                  string `json:"ticker"`
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	Logo                    string `json:"logo"`
	IsActive                bool   `json:"isActive"`
	IsMaintenance           bool   `json:"isMaintenance"`
	MaintenanceNotes        string `json:"maintenanceNotes"`
	IsDelisting             bool   `json:"isDelisting"`
	DelistingNotes          string `json:"delistingNotes"`
	DelistingStartsAt       int    `json:"delistingStartsAt"`
	DelistingEndsAt         int    `json:"delistingEndsAt"`
	HasChildren             bool   `json:"hasChildren"`
	IsChild                 bool   `json:"isChild"`
	IsToken                 bool   `json:"isToken"`
	TokenDetails            string `json:"tokenDetails"`
	UseParentAddress        bool   `json:"useParentAddress"`
	UsdValue                string `json:"usdValue"`
	DepositActive           bool   `json:"depositActive"`
	DepositNotes            string `json:"depositNotes"`
	DepositPayid            bool   `json:"depositPayid"`
	WithdrawalActive        bool   `json:"withdrawalActive"`
	WithdrawalNotes         string `json:"withdrawalNotes"`
	WithdrawalPayid         bool   `json:"withdrawalPayid"`
	WithdrawalPayidRequired bool   `json:"withdrawalPayidRequired"`
	ConfirmsRequired        int    `json:"confirmsRequired"`
	WithdrawDecimals        int    `json:"withdrawDecimals"`
	WithdrawFee             string `json:"withdrawFee"`
	Explorer                string `json:"explorer"`
	ExplorerTxid            string `json:"explorerTxid"`
	ExplorerAddress         string `json:"explorerAddress"`
	Website                 string `json:"website"`
	CoinMarketCap           string `json:"coinMarketCap"`
	CoinGecko               string `json:"coinGecko"`
	Nomics                  string `json:"nomics"`
	CoinPaprika             string `json:"coinPaprika"`
	CoinCodex               string `json:"coinCodex"`
	LiveCoinWatch           string `json:"liveCoinWatch"`
	AddressRegEx            string `json:"addressRegEx"`
	PayidRegEx              string `json:"payidRegEx"`
	SocialCommunity         struct {
		Twitter     string `json:"Twitter"`
		Github      string `json:"Github"`
		Discord     string `json:"Discord"`
		Telegram    string `json:"Telegram"`
		Reddit      string `json:"Reddit"`
		Facebook    string `json:"Facebook"`
		YouTube     string `json:"YouTube"`
		BitcoinTalk string `json:"BitcoinTalk"`
	} `json:"socialCommunity"`
	CoinGeckoAPIID         string `json:"coinGeckoApiId"`
	CoinMarketCapAPIID     string `json:"coinMarketCapApiId"`
	CoinPaprikaAPIID       string `json:"coinPaprikaApiId"`
	NomicsAPIID            string `json:"nomicsApiId"`
	CoinCodexAPIID         string `json:"coinCodexApiId"`
	LiveCoinWatchAPIID     string `json:"liveCoinWatchApiId"`
	IsEVM                  bool   `json:"isEVM"`
	OverrideCirculationURL string `json:"overrideCirculationUrl"`
	About                  string `json:"about"`
	CanLiquidityPool       bool   `json:"canLiquidityPool"`
	CanStake               bool   `json:"canStake"`
	CanVote                bool   `json:"canVote"`
	HasNfts                bool   `json:"hasNfts"`
	ImageUUID              string `json:"imageUUID"`
	CreatedAt              int64  `json:"createdAt"`
	UpdatedAt              int64  `json:"updatedAt"`
	Circulation            string `json:"circulation"`
	IsProofOfWork          bool   `json:"isProofOfWork"`
	IsHighRisk             bool   `json:"isHighRisk"`
	LastPriceUpdate        int64  `json:"lastPriceUpdate"`
	MarketcapNumber        int    `json:"marketcapNumber"`
	HighlightTill          int    `json:"highlightTill"`
	RecentBlockTimeAverage int    `json:"recentBlockTimeAverage"`
	Cryptorank             string `json:"cryptorank"`
	WithdrawFeeAsset       string `json:"withdrawFeeAsset"`
	WithdrawFeeAlt         string `json:"withdrawFeeAlt"`
	WithdrawFeeAltAsset    string `json:"withdrawFeeAltAsset"`
	UnifiedCryptoAssetID   int    `json:"unifiedCryptoAssetId"`
	WithdrawFeeCostAsset   string `json:"withdrawFeeCostAsset"`
	DevEmail               string `json:"devEmail"`
	DevTelegram            string `json:"devTelegram"`
	Developer              string `json:"developer"`
	ID0                    string `json:"id"`
}
*/



type CoingeckoAnswer struct {
	Ticker                  string `json:"ticker"`
	UsdValue                string `json:"usdValue"`
}







//const gCoingeckoEndpoint = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids="
//const gCoingeckoEndpoint = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids="
const gCoingeckoEndpoint = "https://api.xeggex.com/api/v2/asset/getbyticker/"
var CoingeckoPriceRegistry sync.Map

func NewCoingeckoRequest(page int) string {
	url := gCoingeckoEndpoint
	all_coins := getGeckoCoinsList()
	url += strings.Join(all_coins, ",")
//	url += "&x_cg_api_key=CG-8XDVCpAhU2YLcD3EGAU4bzCJ"
//	url += "&order=id_asc&price_change_percentage=24h&sparkline=true&per_page=250"
//	url += "&page=" + fmt.Sprintf("%d", page)
	return url
}

func processCoingecko() *[]CoingeckoAnswer {
var answer = &[]CoingeckoAnswer{}
	page := 1
	for {
		url := NewCoingeckoRequest(page)
		_ = glg.Infof("Processing coingecko request: %s", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Err != nil: %v\n", err)
		}
		if resp.StatusCode == http.StatusOK {
                             var resulta CoingeckoAnswer

			defer resp.Body.Close()
                       bodyBytes, _ := ioutil.ReadAll(resp.Body)



			//decodeErr := json.NewDecoder(resp.Body).Decode(page_answer)
			//if decodeErr != nil {
			//	fmt.Printf("decodeErr: %v\n", decodeErr)
			//}
			//fmt.Printf("Got %v coins form Gecko\n", len(*page_answer))
			// *answer = append(*answer, *page_answer...)

			//if len(*page_answer) == 0 || len(*page_answer) < 250 {
			//	return answer
			//}

                      decodeErr2 := json.Unmarshal(bodyBytes, &resulta)

                        if decodeErr2 != nil {
                               fmt.Printf("decodeErr2: %v\n", decodeErr2)
                        }
                        fmt.Printf("Got %+v coins from coingecko\n", resulta)



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












		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			glg.Errorf("Http status not OK: %s", bodyBytes)
			if len(*answer) == 0 {
				return nil
			}
			return answer
		}
		page += 1
		time.Sleep(600 * time.Second) 
	}



}

func StartCoingeckoService() {
	for {
		if resp := processCoingecko(); resp != nil {
			glg.Info("Coingecko request successfully processed")
			for _, cur := range *resp {
			CoingeckoPriceRegistry.Store(cur.Ticker, cur)
			}
		} else {
			glg.Error("Something went wrong when processing coingecko request")
		}
		time.Sleep(constants.GPricesLoopTime)
	}
}


func CoingeckoRetrieveUSDValIfSupported(coin string) (string, string, string) {
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	valStr := "0"
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
		val, ok := CoingeckoPriceRegistry.Load(cfg.CoingeckoID)
		if ok {
			resp := val.(CoingeckoAnswer)
			valStr = fmt.Sprintf("%f", resp.UsdValue)
//			dateStr = resp.LastUpdated
		}
		return valStr, dateStr, "coingecko"
	}
	return valStr, dateStr, "unknown"
}





func getGeckoCoinsList() []string {
	coins := []string{}
	for _, cur := range config.GCFGRegistry {
		if cur.CoingeckoID != "test-coin" && cur.CoingeckoID != "" {
			coins = append(coins, cur.CoingeckoID)
		}
	}
	coins = helpers.UniqueStrings(coins)
	sort.Strings(coins)
	return coins
}









































/*
func CoingeckoRetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	basePrice, baseDate, _ := CoingeckoRetrieveUSDValIfSupported(base)
	relPrice, relDate, _ := CoingeckoRetrieveUSDValIfSupported(rel)
	price := helpers.BigFloatDivide(basePrice, relPrice, 8)
	date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if helpers.RFC3339ToTimestamp(baseDate) <= helpers.RFC3339ToTimestamp(relDate) {
		date = baseDate
	} else {
		date = relDate
	}
	return price, true, date, "coingecko"
}

func CoingeckoGetTotalVolume(coin string) (string, string, string) {
	totalVolumeStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
		val, ok := CoingeckoPriceRegistry.Load(cfg.CoingeckoID)
		if ok {
			resp := val.(CoingeckoAnswer)
			totalVolumeStr = fmt.Sprintf("%f", resp.TotalVolume)
			dateStr = resp.LastUpdated
		}
		return totalVolumeStr, dateStr, "coingecko"
	}
	return totalVolumeStr, dateStr, "unknown"
}

func CoingeckoGetSparkline7D(coin string) (*[]float64, string, string) {
	var out *[]float64
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
		val, ok := CoingeckoPriceRegistry.Load(cfg.CoingeckoID)
		if ok {
			resp := val.(CoingeckoAnswer)
			if resp.SparklineIn7D != nil && resp.SparklineIn7D.Price != nil {
				out = resp.SparklineIn7D.Price
			}
			dateStr = resp.LastUpdated
		}
		return out, dateStr, "coingecko"
	}
	return out, dateStr, "unknown"
}

func CoingeckoGetChange24h(coin string) (string, string, string) {
	changePercent24h := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
		val, ok := CoingeckoPriceRegistry.Load(cfg.CoingeckoID)
		if ok {
			resp := val.(CoingeckoAnswer)
			if resp.PriceChangePercentage24H != nil {
				changePercent24h = fmt.Sprintf("%f", *resp.PriceChangePercentage24HInCurrency)
			}
			dateStr = resp.LastUpdated
		}
		return changePercent24h, dateStr, "coingecko"
	}
	return changePercent24h, dateStr, "unknown"
}
*/



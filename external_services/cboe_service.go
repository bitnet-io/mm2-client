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



//	"io"

//	"strconv"

)






// GlobalQuoteResponse - encapsulates global quote repsonse
type GlobalQuoteResponse struct {
	CBOEAnswer CBOEAnswer `json:"Global Quote"`
}

// GlobalQuote - encapsulates global quote
type CBOEAnswer struct {
	Id           string  `json:"01. symbol"`
	Open             float64 `json:"02. open,string"`
	High             float64 `json:"03. high,string"`
	Low              float64 `json:"04. low,string"`
	CurrentPrice            float64 `json:"05. price,string"`
	TotalVolume           int     `json:"06. volume,string"`
	LastUpdated string  `json:"07. latest trading day"`
	PreviousClose    float64 `json:"08. previous close,string"`
	Change           float64 `json:"09. change,string"`
	ChangePercentStr string  `json:"10. change percent"`
	ChangePercent    float64
}








const gCBOEEndpoint = "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol="

var CBOEPriceRegistry sync.Map


func NewCBOERequest(page int) string {

	url := gCBOEEndpoint
	all_coins := getCBOECoinsList()
	url += strings.Join(all_coins, "")
				//	url += "&order=id_asc&price_change_percentage=24h&sparkline=true&per_page=250"
				//	url += "&page=" + fmt.Sprintf("%d", page)
				//API KEY FOR ALPHA VANTAGE Y3NX64T240QWMYXF

	url += "&apikey=Y3NX64T240QWMYXF"

	return url
}








//func processCoingecko() *[]CoingeckoAnswer {
func processCBOE() *[]CBOEAnswer {
	var answers = &[]CBOEAnswer{}




	page := 1
	for {
		url := NewCBOERequest(page)
		_ = glg.Infof("Processing CBOE request: %s", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Err != nil: %v\n", err)
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
		var page_answer = &[]CBOEAnswer{}
			decodeErr := json.NewDecoder(resp.Body).Decode(page_answer)
			if decodeErr != nil {
				fmt.Printf("decodeErr: %v\n", decodeErr)
			}
			//fmt.Printf("Got %v coins form CBOE\n", len(*page_answer))
			*answers = append(*answers, *page_answer...)
//			if len(*page_answer) == 0 || len(*page_answer) < 250 {
			return answers


//			}
		} else {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			glg.Errorf("Http status not OK: %s", bodyBytes)
			if len(*answers) == 0 {
				return nil
			}
			return answers


		}
		page += 1
		time.Sleep(10 * time.Second) 
	}




}




func StartCBOEService() {
	for {
		if resp := processCBOE(); resp != nil {
			glg.Info("CBOE request successfully processed")
			for _, cur := range *resp {
				CBOEPriceRegistry.Store(cur.Id, cur)
			}
		} else {
			glg.Error("Something went wrong when processing cboe request")
		}
		time.Sleep(constants.GPricesLoopTime)
	}
}








func CBOEUSDValIfSupported(coin string) (string, string, string) {
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	valStr := "0"
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
		val, ok := CoingeckoPriceRegistry.Load(cfg.CBOEID)
		if ok {
			resp := val.(CBOEAnswer)
			valStr = fmt.Sprintf("%f", resp.CurrentPrice)
			dateStr = resp.LastUpdated
		}
		return valStr, dateStr, "cboe"
	}
	return valStr, dateStr, "unknown"
}




func CBOERetrieveCEXRatesFromPair(base string, rel string) (string, bool, string, string) {
	basePrice, baseDate, _ := CBOEUSDValIfSupported(base)
	relPrice, relDate, _ := CBOEUSDValIfSupported(rel)
	price := helpers.BigFloatDivide(basePrice, relPrice, 8)
	date := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if helpers.RFC3339ToTimestamp(baseDate) <= helpers.RFC3339ToTimestamp(relDate) {
		date = baseDate
	} else {
		date = relDate
	}
	return price, true, date, "cboe"
}




func CBOEVolume(coin string) (string, string, string) {
	totalVolumeStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if cfg, cfgOk := config.GCFGRegistry[coin]; cfgOk {
		val, ok := CoingeckoPriceRegistry.Load(cfg.CBOEID)
		if ok {
			resp := val.(CBOEAnswer)
			totalVolumeStr = fmt.Sprintf("%f", resp.TotalVolume)
			dateStr = resp.LastUpdated
		}
		return totalVolumeStr, dateStr, "cboe"
	}
	return totalVolumeStr, dateStr, "unknown"
}

func getCBOECoinsList() []string {
	coins := []string{}
	for _, cur := range config.GCFGRegistry {
		if cur.CBOEID != "test-coin" && cur.CBOEID != "" {
			coins = append(coins, cur.CBOEID)
		}
	}
	coins = helpers.UniqueStrings(coins)
	sort.Strings(coins)
	return coins
}

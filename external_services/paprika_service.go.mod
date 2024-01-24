package external_services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
//	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/helpers"
	"net/http"
	"sync"
	"time"

	"github.com/kpango/glg"
)

type CoinpaprikaAnswer struct {
	Id                string    `json:"symbol"`
	Name              string    `json:"symbol"`
//	Symbol            string    `json:"symbol"`
	LastUpdated       string    `json:"last_trade_time"`
	Data            struct {
			Price               float64   `json:"current_price"`
	} `json:"data"`
}

//const gPaprikaEndpoint = "https://api.coinpaprika.com/v1/tickers"
const gPaprikaEndpoint = "https://cdn.cboe.com/api/global/delayed_quotes/options/F.json"

var CoinpaprikaRegistry sync.Map

func processCoinpaprika() *[]CoinpaprikaAnswer {
	url := gPaprikaEndpoint
	glg.Infof("Processing coinpaprika request: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		var answer = &[]CoinpaprikaAnswer{}
		decodeErr := json.NewDecoder(resp.Body).Decode(answer)
		if decodeErr != nil {
			fmt.Printf("Err: %v\n", decodeErr)
			return nil
		}
		return answer
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		glg.Errorf("Http status not OK: %s", bodyBytes)
		return nil
	}
}

func StartCoinpaprikaService() {
//	functorVerification := func(ticker string, answer CoinpaprikaAnswer) {
//		if val, ok := config.GCFGRegistry[ticker]; ok && val.CoinpaprikaID == answer.Id {
//			ticker = helpers.RetrieveMainTicker(ticker)
//			CoinpaprikaRegistry.Store(ticker, answer)
//		}
//	}
/*
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
			glg.Error("Something went wrong when processing coinpaprika request")
		}
*/
		time.Sleep(constants.GPricesLoopTime)
//	}

}

func CoinpaprikaRetrieveUSDValIfSupported(coin string) (string, string, string) {
	coin = helpers.RetrieveMainTicker(coin)
	val, ok := CoinpaprikaRegistry.Load(coin)
	valStr := "0"
	dateStr := helpers.GetDateFromTimestampStandard(time.Now().UnixNano())
	if ok {
		resp := val.(CoinpaprikaAnswer)
		valStr = fmt.Sprintf("%f", resp.Data.Price)
//		dateStr = resp.LastUpdated
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


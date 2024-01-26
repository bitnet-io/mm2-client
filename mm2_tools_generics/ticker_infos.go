package mm2_tools_generics

import (
	"encoding/json"
	"fmt"
	"mm2_client/config"
	"mm2_client/external_services"
	"mm2_client/helpers"
)

type TickerInfosRequest struct {
	Ticker string `json:"ticker"`
}


type TickerInfosAnswer struct {
	Ticker               string     `json:"ticker"`
	LastPrice            string     `json:"last_price"`
	LastUpdated          string     `json:"last_updated"`
	LastUpdatedTimestamp int64      `json:"last_updated_timestamp"`
	PriceProvider        string     `json:"price_provider"`
	Sparkline7D          *[]float64 `json:"sparkline_7d"`
	SparklineProvider    string     `json:"sparkline_provider"`
}





type CBOEAnswer struct {

        QuoteResponse struct {


//                                              Result []struct {
  //                                                    RegularMarketPrice                float64  `json:"regularMarketPrice"`
    //                                                  Symbol                string `json:"symbol"`
      //                                        } `json:"result"`

        Result []Result `json:"result"`

        } `json:"quoteResponse"`
}

type Result struct {
                  RegularMarketPrice                string  `json:"regularMarketPrice"`
//                        Symbol                string `json:"symbol"`

	Ticker               string     `json:"ticker"`
	LastPrice            string     `json:"last_price"`
	LastUpdated          string     `json:"last_updated"`
	LastUpdatedTimestamp int64      `json:"last_updated_timestamp"`
	PriceProvider        string     `json:"price_provider"`
	Sparkline7D          *[]float64 `json:"sparkline_7d"`
	SparklineProvider    string     `json:"sparkline_provider"`

}





















//type CBOEAnswer struct {
//        Id                           string    `json:"id"`
  //      Symbol                       string    `json:"symbol"`
    //    CurrentPrice                 float64   `json:"current_price"`
      //  LastUpdated                        string                  `json:"last_updated"`
 //       SparklineIn7D                      *CoingeckoSparkLineData `json:"sparkline_in_7d,omitempty"`

//	Ticker               string     `json:"ticker"`
//	LastPrice            string     `json:"last_price"`
//	LastUpdated          string     `json:"last_updated"`
//	LastUpdatedTimestamp int64      `json:"last_updated_timestamp"`
//	PriceProvider        string     `json:"price_provider"`
//	Sparkline7D          *[]float64 `json:"sparkline_7d"`
//	SparklineProvider    string     `json:"sparkline_provider"`

  //      QuoteResponse struct {
//	              Result []struct {
  //                                                    RegularMarketPrice                float64  `json:"regularMarketPrice"`
    //                                                  Symbol                string `json:"symbol"`
      //                                        } `json:"result"`

	//	        Result []Result `json:"result"`
      //  } `json:"quoteResponse"`
//}

//type Result struct {
  //                  RegularMarketPrice                float64  `json:"regularMarketPrice"`
    //                    Symbol                string `json:"symbol"`
//}














func (req *TickerInfosAnswer) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}
func (req *CBOEAnswer) ToJson() string {
	b, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(b)
}

func (req *TickerInfosAnswer) ToWeb() map[string]interface{} {
	out := make(map[string]interface{})
	if req != nil {
		_ = json.Unmarshal([]byte(req.ToJson()), &out)
		return out
	}
	return nil
}

func (req *CBOEAnswer) ToWeb() map[string]interface{} {
	out := make(map[string]interface{})
	if req != nil {
		_ = json.Unmarshal([]byte(req.ToJson()), &out)
		return out
	}
	return nil
}

func GetTickerInfos(ticker string, expirePriceValidity int) *TickerInfosAnswer {
	outTicker := ticker
	if cfg, cfgOk := config.GCFGRegistry[ticker]; cfgOk && cfg.AliasTicker != nil {
		outTicker = *cfg.AliasTicker
	}
	val, date, provider := external_services.RetrieveUSDValIfSupported(outTicker, expirePriceValidity)
	return &TickerInfosAnswer{Ticker: ticker, LastPrice: val,
		 LastUpdated: date,
		LastUpdatedTimestamp: helpers.RFC3339ToTimestampSecond(date),
		PriceProvider:        provider}
}


func CBOETickerInfos(ticker string, expirePriceValidity int) *Result {
	outTicker := ticker
	if cfg, cfgOk := config.GCFGRegistry[ticker]; cfgOk && cfg.AliasTicker != nil {
		outTicker = *cfg.AliasTicker
	}
	val, date, provider := external_services.RetrieveUSDValIfSupported(outTicker, expirePriceValidity)
//	return &CBOEAnswer{}

	return &Result{RegularMarketPrice: val, 
	Ticker: ticker, LastPrice: val,
		 LastUpdated: date,
		LastUpdatedTimestamp: helpers.RFC3339ToTimestampSecond(date),
		PriceProvider:        provider}

}

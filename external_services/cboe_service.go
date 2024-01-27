
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























type CBOEAnswer struct {
//  Id                           string    `json:"id"`
//        Symbol                       string    `json:"symbol"`
  //      CurrentPrice                 float64   `json:"current_price"`
    //    LastUpdated                        string                  `json:"last_updated"`

      //  SparklineIn7D                      *CoingeckoSparkLineData `json:"sparkline_in_7d,omitempty"`

	QuoteResponse struct {


				//		Result []struct {
				//			RegularMarketPrice                float64  `json:"regularMarketPrice"`
				//			Symbol                string `json:"symbol"`
				//		} `json:"result"`

	Result []Result `json:"result"`

	} `json:"quoteResponse"`
}

type Result struct {
		    RegularMarketPrice                float64  `json:"regularMarketPrice"`
                        Symbol                string `json:"symbol"`
}

const gCBOEEndpoint = "https://apidojo-yahoo-finance-v1.p.rapidapi.com/market/v2/get-quotes?region=US&symbols="


var CBOEPriceRegistry sync.Map


func NewCBOERequest(page int) string {
        url := gCBOEEndpoint
        all_coins := getCBOECoinsList()
        url += strings.Join(all_coins, ",")
	url += "%2C"
	url += "&rapidapi-key=91ac13b7e0msh58a82c551837a5cp18ee79jsn4664b59ac23e"
        return url

}



























var CBOEAnswers struct {
  //      Symbol                       string    `json:"symbol"`
        QuoteResponse struct {
        Result []Result `json:"result"`
        } `json:"quoteResponse"`
}

var Results struct {
                    RegularMarketPrice                float64  `json:"regularMarketPrice"`
                        Symbol                string `json:"symbol"`
}




type Resultsx struct {
    Resultsx []Resultx `json:"resultx"`
}

type Resultx struct {
                    RegularMarketPrice                float64  `json:"regularMarketPrice"`
                        Symbol                string `json:"symbol"`
}





func processCBOE() *[]CBOEAnswer {
	var answer = &[]CBOEAnswer{}
	page := 1
	for {
		url := NewCBOERequest(page)
		_ = glg.Infof("Processing cboe request: %s", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Err != nil: %v\n", err)
		}
		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
//			var page_answer = &[]CBOEAnswer{}
	                     var resulta CBOEAnswer
	               bodyBytes, _ := ioutil.ReadAll(resp.Body)

		//	decodeErr := json.NewDecoder(resp.Body).Decode(resulta)
                      decodeErr := json.Unmarshal(bodyBytes, &resulta)

			if decodeErr != nil {
				fmt.Printf("decodeErr: %v\n", decodeErr)
			}
//			fmt.Printf("Got %v stocks form CBOE\n", len(*resulta))
		        fmt.Printf("Got %+v stocks from CBOE\n", resulta)
//			*answer = append(*answer, resulta...)
//			if len(*resulta) == 0 || len(*resulta) < 250 {
//				return answer
//			}


//    var resultsx Resultx

//    rankings := Resultx{}
    // read our opened xmlFile as a byte array.
//    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
//    json.Unmarshal(&resultsx)

    rankingsJson, _ := json.Marshal(resulta)
    err = ioutil.WriteFile("./stocks/stocks.json", rankingsJson, 0644)
    fmt.Printf("%+v", resulta)


	//Open our jsonFile
    jsonFile, err := os.Open("./stocks/stocks.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened stocks.json")
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
		time.Sleep(60 * time.Second) 
	}






}









/*
func processSTART() *[]Result {



 var answer2 = &[]Result{}


        page := 1
        for {
                url := NewCBOERequest(page)
                _ = glg.Infof("Processing cboe http request: %s", url)
                resp, err := http.Get(url)
                if err != nil {
                        fmt.Printf("Err != nil: %v\n", err)
                }
                if resp.StatusCode == http.StatusOK {
                        defer resp.Body.Close()
                  var page_answer2 = &[]Result{}

					//		      body := ioutil.ReadAll(resp.Body) // response body is []byte
	               bodyBytes, _ := ioutil.ReadAll(resp.Body)



                     var result2 Result
 ///	                   var result = &[]Result{}

  	              decodeErr := json.Unmarshal(bodyBytes, &result2)
                        if decodeErr != nil {
                               fmt.Printf("decodeErr: %v\n", decodeErr)
                        }

//		        fmt.Printf("Got %+v stocks from cboe\n", result2)
//                        fmt.Printf("Got %v stocks form cboe\n", len(*page_answer2))

                       // fmt.Printf("Got %v stocks form cboe\n", len(*result))
 
                        *answer2 = append(*answer2, *page_answer2...)
			if len(*page_answer2) == 0 || len(*page_answer2) < 250 {
                                return answer2
                        }

                } else {
                        bodyBytes, _ := ioutil.ReadAll(resp.Body)
                        glg.Errorf("Http status not OK: %s", bodyBytes)
                      if len(*answer2) == 0 {
                                return nil
                        }
                        return answer2
                }
                page += 1
                time.Sleep(10 * time.Second) 
        }


}
*/












func StartCBOEService() {
	 for {
               if resp := processCBOE(); resp != nil {
                        glg.Info("CBOE request successfully processed")
//                        for _, cur := range *resp {
//                               CBOEPriceRegistry.Store(cur.Symbol, cur)
  //                     }
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
                val, ok := CBOEPriceRegistry.Load(cfg.CBOEID)
                if ok {

                        resp := val.(Result)
	                valStr = fmt.Sprintf("%f", resp.RegularMarketPrice)
		          dateStr = resp.Symbol

                }
                return valStr, dateStr, "cboe"
        }
        return valStr, dateStr, "unknown"
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

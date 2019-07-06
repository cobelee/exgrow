package alistock

import (
	// "encoding/json"
	"fmt"
	// "io"
	"io/ioutil"
	"net/http"
)

func FetchTradeData() {

	// Query string
	END_TIME := "20180101"
	START_TIME := "20190417"
	STOCK_CODE := "600744"
	queryStr := "?END_TIME=" + END_TIME + "&START_TIME=" + START_TIME + "&STOCK_CODE=" + STOCK_CODE
	url := "http://istock.market.alicloudapi.com/ai_fintech_knowledge/ai_stock_trade_market"
	url += queryStr
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "APPCODE c9b51bc37d69424ab32087785248a49a")
	resp, _ := client.Do(req)
	fmt.Print("Request.\n")
	defer resp.Body.Close() // The body maybe is nil, can't be read.
	// var mapStock map[string]StockCode
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Print(string(body))
		// decoder := json.NewDecoder(resp.Body)

		// for {
		// 	// err := decoder.Decode(&mapStock)
		// 	// if err == io.EOF {
		// 	// 	break
		// 	// }

		// 	// if err != nil {
		// 	// 	fmt.Println("Error decoding JSON:", err)
		// 	// 	return nil
		// 	// }
		// }
	} else {
		fmt.Println(resp.Status)
	}

	// return mapStock

}

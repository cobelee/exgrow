package alistock

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

type StockCode struct {
	Country_ch   string `json:"STOCK_COUNTRY_CH"`
	Country_en   string `json:"STOCK_COUNTRY_EN"`
	Country_abbr string `json:"STOCK_COUNTRY_ABBR"`
	Market_ch    string `json:"STOCK_MARKET_CH"`
	Market_en    string `json:"STOCK_MARKET_EN"`
	Market_abbr  string `json:"STOCK_MARKET_ABBR"`
	Board_ch     string `json:"STOCK_BOARD_CH"`
	Board_en     string `json:"STOCK_BOARD_EN"`
	Name_ch_abbr string `json:"STOCK_NAME_CH_ABBR"`
	Name_en_abbr string `json:"STOCK_NAME_EN_ABBR"`
	Code         string `json:"STOCK_CODE"`
}

func GetStockMap() map[string]StockCode{

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://istock.market.alicloudapi.com/ai_fintech_knowledge/ai_stock_code", nil)
	req.Header.Set("Authorization", "APPCODE c9b51bc37d69424ab32087785248a49a")

	resp, _ := client.Do(req)

	defer resp.Body.Close()		// The body maybe is nil, can't be read.
	var mapStock map[string]StockCode
	if resp.StatusCode == 200 {
//		body, _ := ioutil.ReadAll(resp.Body)
		
		decoder := json.NewDecoder(resp.Body)

		for {
			err := decoder.Decode(&mapStock)
			if err == io.EOF{
				break
			}

			if err != nil{
				fmt.Println("Error decoding JSON:", err)
				return nil
			}
		}
	}


	return mapStock

}

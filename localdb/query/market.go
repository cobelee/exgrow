package query

import (
	c "exgrow/localdb/config"
	h "exgrow/localdb/dbhelp"
	"fmt"
	"strings"
)

var opt ShowStockCodeOpt

func init() {
	opt = ShowStockCodeOpt{
		Market: "",
		Col:    14,
	}
}

// Get the count of stocks in the market.
func GetCountofStocks() int {
	var codes []string
	codes = getStockCode("")
	cCount := len(codes)
	return cCount
}

// Get the count of stocks in shanghai market.
func GetCountofStocksInShanghai() int {
	var codes []string
	codes = getStockCode("sh")
	c := len(codes)
	return c
}

// Get the count of stocks in shenzhen market.
func GetCountofStocksInShenzhen() int {
	var codes []string
	codes = getStockCode("sz")
	c := len(codes)
	return c
}

/* Get list of stock code in market.

   param m   market name.       May be "sh", "sz" or ""
*/
func getStockCode(m string) []string {
	dbName := c.DBConfig.DBName.StockD1
	names, _ := h.GetCollectionNames(dbName)

	var namesFiltered []string

	switch m {
	case "sh":
		for _, name := range names {
			if strings.HasPrefix(name, "sh") {
				namesFiltered = append(namesFiltered, name)
			}

		}
		return namesFiltered
	case "sz":
		for _, name := range names {
			if strings.HasPrefix(name, "sz") {
				namesFiltered = append(namesFiltered, name)
			}

		}
		return namesFiltered
	default:
		return names

	}
}

type ShowStockCodeOpt struct {
	Market string
	Col    int
}

func NewDefaultShowStockCodeOpt() ShowStockCodeOpt {
	return ShowStockCodeOpt{
		Market: "", // Stock market, may be "sh", "sz" or ""
		Col:    14, // Column count, default value 14
	}
}

// Show stock code in market.
func ShowStockCode(opt ShowStockCodeOpt) {
	var codes []string
	codes = getStockCode(opt.Market)

	var col int
	if opt.Col > 5 {
		col = opt.Col
	} else {
		col = 5
	}

	for i, name := range codes {
		if i%col == 0 {
			fmt.Println()
		}
		fmt.Printf("%12s", name)
	}
	fmt.Println()
}

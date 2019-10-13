package clicmd

import (
	q "exgrow/localdb/query"
	"fmt"

	"github.com/urfave/cli"
)

var CmdMarket cli.Command
var short bool
var market string

func init() {

	CmdMarket = cli.Command{
		Name:     "market",
		Usage:    "证券市场基础信息查询。",
		Category: "查询",
		Subcommands: []cli.Command{
			{
				Name:    "stockcode",
				Aliases: []string{"sc"},
				Usage:   "List stock code.",
				Action: func(c *cli.Context) error {
					opt := q.NewDefaultShowStockCodeOpt()
					switch c.String("m") {
					case "sh", "shanghai":
						opt.Market = "sh"
						opt.Col = 14
						q.ShowStockCode(opt)

					case "sz", "shenzhen":
						opt.Market = "sz"
						opt.Col = 14
						q.ShowStockCode(opt)
					default:
						opt.Market = ""
						opt.Col = 14
						q.ShowStockCode(opt)
					}

					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "market,m",
						Value: "",
						Usage: `List stock code. The available value could be:
								'sh', shanghai		shanghai market
								'sz', shenzhen		shenzhen market
								'',					All market`,
						Destination: &market,
					},
				},
			},
		},
		Action: func(c *cli.Context) error {
			var csh, csz, cTotal int
			csh = q.GetCountofStocksInShanghai()
			csz = q.GetCountofStocksInShenzhen()
			cTotal = q.GetCountofStocks()

			if short {
				fmt.Println("Stock Count")
				fmt.Printf("      Shanghai		%v\n", csh)
				fmt.Printf("      Shenzhen		%v\n", csz)
				fmt.Printf("      Total		%v\n", cTotal)
			} else {
				fmt.Printf("The count of stocks in market: %v\n", cTotal)
			}

			return nil
		},
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "info,i",
				Usage:       "Show market information.",
				Destination: &short,
			},
		},
	}

}

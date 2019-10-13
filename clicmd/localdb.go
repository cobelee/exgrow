package clicmd

import (
	"exgrow/localdb"
	k "exgrow/localdb/analysis/keltner"
	m "exgrow/localdb/maintain"
	o "exgrow/localdb/object"
	"exgrow/testcode"
	"fmt"
	"time"

	"github.com/urfave/cli"
)

var CmdLocaldb cli.Command
var fromdatestr, period, mode string

func init() {
	CmdLocaldb = cli.Command{
		Name:     "db",
		Usage:    "证券市场基础数据维护。",
		Category: "数据库维护",
		Subcommands: []cli.Command{
			{
				Name:    "syncsc",
				Aliases: []string{"ssc"},
				Usage:   "将证券品种的基础代码信息，从阿里云同步到本地数据库。",
				Action: func(c *cli.Context) error {
					fmt.Println("正在同步证券品种代码，请稍候...")
					localdb.SyncStockCode()
					return nil
				},
			},
			{
				Name:    "readfiles",
				Aliases: []string{"ls"},
				Usage:   "读取 /data/ 目录下的所有文件。",
				Action: func(c *cli.Context) error {
					testcode.ReadFiles()
					return nil
				},
			},
			{
				Name:    "importHistoryData",
				Aliases: []string{"ihd"},
				Usage:   "将来自于“预测者”网站（www.yucezhe.com）的证券历史日线数据导入mongo数据库。",
				Action: func(c *cli.Context) error {
					localdb.ImportSHD_FromDir()
					var str string
					fmt.Scan(&str)
					return nil
				},
			},
			{
				Name:    "importDailyData",
				Aliases: []string{"idd"},
				Usage:   "将来自于“预测者”网站（www.yucezhe.com）的单日日线数据导入mongo数据库。包括证券以及指数数据文件。",
				Action: func(c *cli.Context) error {
					localdb.ImportDD_FromDir()
					return nil
				},
			},
			{
				Name:    "TypifyHeaderLine",
				Aliases: []string{"thl"},
				Usage:   "预处理“预测者”网站（www.yucezhe.com）的证券历史日线数据。\n对首行Fields进行格式化。",
				Action: func(c *cli.Context) error {
					localdb.TypifyHeaderLine_FromDir()
					return nil
				},
			},
			{
				Name:    "MongoImportStock",
				Aliases: []string{"mis"},
				Usage:   "利用MongoImport工具，将“预测者”网站（www.yucezhe.com）的证券历史日线数据导入数据库。",
				Action: func(c *cli.Context) error {
					localdb.MongoImportSHD_FromDir()
					return nil
				},
			},
			{
				Name:    "MongoImportIndex",
				Aliases: []string{"mii"},
				Usage:   "利用MongoImport工具，将“预测者”网站（www.yucezhe.com）的指数历史日线数据导入数据库。",
				Action: func(c *cli.Context) error {
					localdb.MongoImportIHD_FromDir()
					return nil
				},
			},

			{
				Name:    "GenerateStockD1",
				Aliases: []string{"gsd"},
				Usage:   "生成 StockD1 数据库，这是所有 StockX1 数据库的基本库。",
				Action: func(c *cli.Context) error {
					m.GenerateStockD1()
					return nil
				},
			},
			{
				Name:    "SyncStockRTD",
				Aliases: []string{"ssrtd"},
				Usage:   "依据 StockMarketRawD1 数据库，增量数据同步至  StockD1 数据库。",
				Action: func(c *cli.Context) error {
					m.BeginSyncRTD()
					return nil
				},
			},
			{
				Name:    "sync",
				Aliases: []string{"sync"},
				Usage: `Synchronize base-info from StockD1 to other period-database. 
						Flag: --period, -p`,
				Action: func(c *cli.Context) error {
					switch c.String("p") {
					case "w", "W", "week":
						m.BeginSync(o.DTW)

					case "m", "M", "month":
						m.BeginSync(o.DTM)

					case "q", "Q", "quarter":
						m.BeginSync(o.MTQ)
					case "y", "Y", "year":
						m.BeginSync(o.MTY)

					default:
						fmt.Print("    Wrong value of flag p.\n    the available value is w, m, q or y.\n")
					}

					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "period,p",
						Value: "",
						Usage: `根据日线数据库，同步至周线、月线、季线数据库。
								W, week		生成周线数据库
								M, month		生成月线数据库
								Q, quarter	生成季线数据库
								Y, year      生成年线数据库`,
					},
				},
			},

			{
				Name:    "FillIndic",
				Aliases: []string{"fillindic"},
				Usage: `Fill the indications for StockD1 W1 M1 Q1 Y1.
						Flag: --period|-p  ,  --mode|-m`,
				Action: func(c *cli.Context) error {
					m.FillIndications(period, mode)
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "period,p",
						Value:       "d",
						Usage:       "指定要更新数据库的周期，其值可能为 d w m q y，all(全部周期)",
						Destination: &period,
					},
					cli.StringFlag{
						Name:        "mode,m",
						Value:       "increment",
						Usage:       "指标数据填充模式，可tf值可能为 increment(增量模式), rerun（重新计算）",
						Destination: &mode,
					},
				},
			},
			{
				Name:    "kent",
				Aliases: []string{"kent"},
				Usage:   "Give a kent band count report",
				Action: func(c *cli.Context) error {
					k.BandRanking()
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "Remove stock record from database Stock D1 W1 M1 Q1 Y1.",
				Action: func(c *cli.Context) error {
					t, _ := time.Parse("2006-01-02", fromdatestr)
					m.RemoveStockRecordSinceDate(t)

					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:        "fromtime, t",
						Value:       "",
						Usage:       "指定此时间之后的记录将被删除。",
						Destination: &fromdatestr,
					},
				},
			},
		},
	}
}

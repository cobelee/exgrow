package main

import (
	"bufio"
	"exgrow/localdb"
	"exgrow/testcode"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "exgrow"
	app.Usage = "股票市场行情自动跟踪系统"
	app.Version = "1.0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "系统显示语言选项",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "add",
			Aliases:  []string{"a"},
			Usage:    "测试命令，添加一个任务到列表.",
			Category: "Calulation",
			Action: func(c *cli.Context) error {
				fmt.Println("added task:", c.Args().First())
				return nil
			},
		},
		{
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
					Usage:   "将来自于“预测者”网站（www.yucezhe.com）的单日日线数据导入mongo数据库。\n 包括证券以及指数数据文件。",
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
					Name:    "TestSort",
					Aliases: []string{"test"},
					Usage:   "测试函数。",
					Action: func(c *cli.Context) error {
						localdb.TestSDCard()
						return nil
					},
				},
			},
		},
		{
			Name:     "showday",
			Aliases:  []string{"sd"},
			Usage:    "测试命令，显示当前星期数.",
			Category: "Test",
			Action: func(c *cli.Context) error {
				testcode.ShowDay()
				return nil
			},
		},
		{
			Name:     "testUpsert",
			Aliases:  []string{"upsert"},
			Usage:    "测试命令，将用户对象更新到mongo数据库.",
			Category: "Test",
			Action: func(c *cli.Context) error {
				testcode.TestUpsert()
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) {

		fmt.Printf(strings.Join(c.Args(), " ") + "\n")
		if c.NArg() == 0 {
			return
		} else {
			fmt.Printf("Can not find command: %s\nRun command %s help to get help.\n", c.Args().Get(0), app.Name)
			return
		}

		if c.String("lang") == "english" {
			fmt.Printf("Set language to english.\n")
		}
		if c.String("lang") == "chinese" {
			fmt.Printf("Set language to chinese. \n")
		}

		return
	}
	// app.Run(os.Args)

	var prompt string
	prompt = fmt.Sprintf("-------------------------\n %c[33;33;1m@>>%c[0m ", 0x1B, 0x1B)
L:

	for {
		var input string
		fmt.Print(prompt)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use 'for scanner.Scan()' to keep reading
		input = scanner.Text()

		switch input {
		case "close":
			fmt.Println("close.\n")
			break L
		case "exit":
			fmt.Printf("%s is exited.\n", app.Name)
			break L
		case "quit":
			fmt.Printf("%s is quited.\n", app.Name)
			break L
		default:
		}

		cmdArgs := strings.Split(input, " ")
		if len(cmdArgs) == 0 {
			continue
		}

		s := []string{app.Name}
		s = append(s, cmdArgs...)

		app.Run(s)

	}

}

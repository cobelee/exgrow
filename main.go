package main

import (
	"bufio"
	c "exgrow/clicmd"
	_ "exgrow/localdb/query"
	_ "exgrow/localdb/test"
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
			Name:     "query",
			Usage:    "查询证券市场基础数据。",
			Category: "数据库查询",
			Subcommands: []cli.Command{
				{
					Name:    "stock",
					Aliases: []string{"stock"},
					Usage:   "测试",
					Action: func(c *cli.Context) error {

						return nil
					},
				},
			},
		},
		{
			Name:     "testtime",
			Aliases:  []string{"tt"},
			Usage:    "测试命令，显示当前星期数.",
			Category: "Test",
			Action: func(c *cli.Context) error {
				testcode.TestTime()
				return nil
			},
		},
		{
			Name:     "test",
			Aliases:  []string{"test"},
			Usage:    "测试命令，将用户对象更新到mongo数据库.",
			Category: "Test",
			Action: func(c *cli.Context) error {
				testcode.TestSDCard()
				return nil
			},
		},
	}

	// 添加其它控制台命令
	app.Commands = append(app.Commands, c.CmdLocaldb)
	app.Commands = append(app.Commands, c.CmdMarket)

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

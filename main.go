package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"sort"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "katze"
	app.Usage = "katze a simple rest api framework for golang"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Abdesselem Aymen",
			Email: "abdesselemaymen@gmail.com",
		},
	}
	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{
	// 		Name:  "new",
	// 		Usage: "create a new golang api project",
	// 	},
	// }
	app.Commands = []cli.Command{
		{
			Name: "new",
			// Aliases: []string{"c"},
			Usage: "create new golang rest api project",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start go server",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:      "generate",
			Aliases:   []string{"g"},
			Usage:     "generator for create controller, model ...",
			ArgsUsage: "[controller, model]",
			Category:  "Generator",
			Subcommands: []cli.Command{
				{
					Name: "controller",
					// Aliases:  []string{"c"},
					Usage:    "add a new controller (exp: g controller userController)",
					Category: "Generator",
					Action: func(c *cli.Context) error {
						fmt.Println("new controller: ", c.Args().First())
						return nil
					},
				},
				{
					Name: "model",
					// Aliases:  []string{"m"},
					Usage:    "add a new model (exp: g model userModel)",
					Category: "Generator",
					Action: func(c *cli.Context) error {
						fmt.Println("new model: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("katze a simple rest api framework for golang")
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
}

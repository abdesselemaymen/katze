package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "katze"
	app.Usage = "this is may first framework go"
	app.Action = func(c *cli.Context) error {
		fmt.Println("good exec")
		return nil
	}

	app.Run(os.Args)
}

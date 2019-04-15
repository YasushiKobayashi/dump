package main

import (
	"os"

	"github.com/YasushiKobayashi/dump/handler/cli_handler"
	"github.com/urfave/cli"
)

var Version string = "0.0.2"

func main() {
	app := cli.NewApp()
	app.Name = "dump"
	app.Usage = "Upload sql dump tool"
	app.Version = Version
	app.Author = "Yasushi Kobayashi"
	app.Email = "ptpadan@gmail.com"
	app.Commands = cli_handler.Commands

	app.Run(os.Args)
}

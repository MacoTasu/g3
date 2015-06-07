package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "g3"
	app.Version = Version
	app.Usage = ""
	app.Author = "MacoTasu"
	app.Email = "maco.tasu@gmail.com"
	app.Commands = Commands
	app.Run(os.Args)
}

package main

import (
	"log"
	"os"

	"github.com/hqpko/gomod/mod"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomod"
	app.Version = "0.1.1"
	app.Usage = "go mod for great firewall"

	app.Commands = []cli.Command{
		{
			Name:   "tidy",
			Usage:  "add missing and remove unused modules",
			Action: mod.ActionGoModTidy,
		},
		{
			Name:   "tidy1",
			Usage:  "add missing and remove unused modules",
			Action: mod.ActionGomodTidy1,
		},
		{
			Name:   "graph",
			Usage:  "print module requirement graph",
			Action: mod.ActionGoModGraph,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

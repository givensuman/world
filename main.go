package main

import (
	"log"
	"os"

	"github.com/Jguer/go-alpm/v2"
	cli "github.com/urfave/cli/v2"

	// "github.com/JoshuaDoes/alpm-go/v2"
	paconf "github.com/Morganamilo/go-pacmanconf"
)

func main() {
	pkg, err := alpm.Initialize("/", "/var/lib/pacman")
	if err != nil {
		print(err)
		os.Exit(1)
	}

	PacmanConfig, _, err := paconf.ParseFile("/etc/pacman.conf")
	if err != nil {
		print(err)
		os.Exit(1)
	}

	app := &cli.App{
		Name:  "world",
		Usage: "Declarative package manager for Arch, btw",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a", "ad"},
				Usage:   "Adds package(s) to world file and installs it",
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r", "rem"},
				Usage:   "Removes package(s) from world file(s) and uninstalls it",
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "install",
				Aliases: []string{"i", "in", "inst"},
				Usage:   "Installs package(s) without adding it to world file",
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "uninstall",
				Aliases: []string{"u", "ui", "un", "uninst"},
				Usage:   "Uninstalls package(s) without removing it from world file(s)",
			},
			{
				Name:    "sync",
				Aliases: []string{"s", "sy", "syn"},
				Usage:   "Synchronizes system with world files",
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "backup",
				Aliases: []string{"b", "bk", "bu", "back"},
				Usage:   "Creates a backup of current world state",
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
		},
		Flags: []cli.Flag{},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

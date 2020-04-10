package main

import (
	"log"
	"os"
	"path"

	"github.com/urfave/cli/v2"
)

var configFilePath string
var version string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configFilePath = path.Join(homeDir, ".kanrc.yml")
	if _, err := os.Stat(configFilePath); err == nil {

	} else if os.IsNotExist(err) {
		file, err := os.Create(configFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	} else {
		panic(err)
	}
}

func main() {
	app := &cli.App{
		Name:     "kan-config",
		Usage:    "make configuration for kan cli",
		HelpName: "kan-config",
		Version:  version,
	}
	app.UseShortOptionHandling = true
	/*
		kan-config init --access-key 123 --secret-key 456
	*/
	app.Commands = []*cli.Command{
		_init(configFilePath),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

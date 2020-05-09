package main

import (
	"fmt"
	"os"

	"git.lozi.space/beast/be-mandat/api"
	"git.lozi.space/beast/be-mandat/app"
	"git.lozi.space/beast/be-mandat/util"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func runServer(c *cli.Context) error {
	configFile := c.String("config")
	app, err := app.New(configFile)
	if err != nil {
		return err
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	api := api.Init(app, router)
	err = api.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "st"
	app.Flags = []cli.Flag{
		util.StringFlag("ST_CONFIG", "config", "configuration file path", "config.yaml"),
	}
	app.Action = runServer

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

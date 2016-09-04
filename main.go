package main

import (
	"github.com/slapers/ping-fastcgi/action"
	"github.com/urfave/cli"
	"os"
)

const version = "0.0.1"

func main() {

	app := cli.NewApp()
	app.Name = "ping-fastcgi"
	app.Version = version
	app.Usage = "Commandline tool to sent a http ping request to fastcgi"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Value: "127.0.0.1", Usage: "Ip or hostname of the fastcgi process"},
		cli.IntFlag{Name: "port", Value: 9000, Usage: "Port of the fastcgi process"},
		cli.StringFlag{Name: "uri", Value: "/ping", Usage: "Uri of the ping endpoint, default is /ping"},
		cli.StringFlag{Name: "response", Value: "pong", Usage: "Expected response from the endpoint"},
	}
	app.Action = action.PingAction
	app.Run(os.Args)
}

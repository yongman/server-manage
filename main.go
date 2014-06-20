package main

import (
	"./command/redis"
	"./command/test"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	cmds := []cli.Command{
		test.Command,
		redis.Command,
	}

	app := cli.NewApp()
	app.Name = "server-manage"
	app.Usage = "server manage toolkit"
	app.Commands = cmds

	app.Run(os.Args)

}

package main

import (
	"./command/gather"
	"./command/redis"
	"./command/test"
	//"./modules/db"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	cmds := []cli.Command{
		test.Command,
		redis.Command,
		gather.Command,
	}

	app := cli.NewApp()
	app.Name = "server-manage"
	app.Usage = "server manage toolkit"
	app.Commands = cmds

	app.Run(os.Args)
	//defer db.Mon.Close()
}

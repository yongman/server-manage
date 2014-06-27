package main

import (
	"./command/print"
	"./command/redis"
	"./command/redisproxy"
	"./command/test"
	"./command/update"
	//"./modules/db"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	cmds := []cli.Command{
		test.Command,
		redis.Command,
		redisproxy.Command,
		update.Command,
		print.Command,
	}

	app := cli.NewApp()
	app.Name = "server-manage"
	app.Usage = "server manage toolkit"
	app.Commands = cmds
	app.Run(os.Args)
	//defer db.Mon.Close()
}

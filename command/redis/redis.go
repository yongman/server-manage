package redis

import (
	//mdao "../../modules/db/machine"
	"../../utils"
	"fmt"
	"github.com/codegangsta/cli"
)

var (
	Command = cli.Command{
		Name:      "redis",
		ShortName: "r",
		Usage:     "redis",
		Action:    redisAction,
		Flags: []cli.Flag{
			cli.StringFlag{"d", "", "pid, if provided, will avoid the servers that have deployed instances of this pid"},
			cli.StringFlag{"m", "", "memory size want to alloc, eg. 512M 1G"},
			cli.IntFlag{"bj", 0, "amount alloc from machineroom in beijing"},
			cli.IntFlag{"hz", 0, "amount alloc from machineroom in hangzhou"},
			cli.IntFlag{"nj", 0, "amount alloc from machineroom in nanjing"},
		},
		Description: Usage,
	}
	Usage = `
Name:
redis alloc tools
        `
)

func redisAction(c *cli.Context) {
	fmt.Println("redisAction")
	m := c.String("m")
	if m != "" {
		size := utils.CapToKB(m)
		fmt.Println(size)
	}
}

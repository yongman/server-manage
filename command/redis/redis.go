package redis

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var (
	Command = cli.Command{
		Name:        "redis",
		ShortName:   "redis",
		Usage:       "redis",
		Action:      redisAction,
		Flags:       []cli.Flag{},
		Description: Usage,
	}
	Usage = `
Name:
redis alloc tools
        `
)

func redisAction(c *cli.Context) {
	fmt.Println("redisAction")
}

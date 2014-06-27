package update

import (
	"../../modules/commit"
	"../../modules/log"
	"../../modules/update/machine"
	"../../modules/update/service"
	//"../../utils"
	"github.com/codegangsta/cli"
)

var (
	Command = cli.Command{
		Name:      "update",
		ShortName: "u",
		Usage:     "update [-s service.info] [-m machine.info]",
		Action:    updateAction,
		Flags: []cli.Flag{
			cli.StringFlag{"s", "", "service info raw file"},
			cli.StringFlag{"m", "", "machine info raw file"},
		},
	}
)

func updateAction(c *cli.Context) {
	s := c.String("s")
	if s != "" {
		count := service.UpdateService(s)
		log.Info("update", count, " services into db")
		log.Info("refreshing commits....")
		//扫描commit表，扫描对应的commit是否存在服务表中，不存在自动撤销相应commit
		commit.RefreshCommit()
	}

	m := c.String("m")
	if m != "" {
		count := machine.UpdateMachine(m)
		log.Info("update", count, " machines into db")
	}
}

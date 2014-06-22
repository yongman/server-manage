package gather

import (
	"../../modules/gather/machine"
	"../../modules/gather/service"
	"../../modules/log"
	//"../../utils"
	"github.com/codegangsta/cli"
)

var (
	Command = cli.Command{
		Name:      "gather",
		ShortName: "g",
		Usage:     "gather [-s service.info] [-m machine.info]",
		Action:    gatherAction,
		Flags: []cli.Flag{
			cli.StringFlag{"s", "", "service info raw file"},
			cli.StringFlag{"m", "", "machine info raw file"},
		},
	}
)

func gatherAction(c *cli.Context) {
	s := c.String("s")
	if s != "" {
		count := service.UpdateService(s)
		log.Info("update", count, " services into db")
	}

	m := c.String("m")
	if m != "" {
		count := machine.UpdateMachine(m)
		log.Info("update", count, " machines into db")
	}
}

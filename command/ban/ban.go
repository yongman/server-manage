package ban

import (
	mdao "../../modules/db/machine"
	"../../modules/fmtoutput"
	"../../modules/log"
	"../../modules/update/machine"
	"../../utils"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

var (
	Command = cli.Command{
		Name:      "ban",
		ShortName: "b",
		Usage:     "ban -host <hostname> -s <redis|redisproxy|memcache>",
		Action:    banAction,
		Flags: []cli.Flag{
			cli.StringFlag{"host", "", "the hostname want to operate"},
			cli.StringFlag{"s", "", "the service want to ban to this host "},
		},
		Description: Usage,
	}

	Usage = `
NAME:
  ban -host <hostname> -s <redis|redisproxy|memcache>
`
)

func banAction(c *cli.Context) {
	h := c.String("host")
	s := c.String("s")

	if h == "" || s == "" {
		log.Fatal("arg error: -host <hostname> -s <redis|redisproxy|memcache>")
		os.Exit(1)
	}
	if strings.HasSuffix(h, ".baidu.com") == false {
		h = fmt.Sprintf("%s%s", h, ".baidu.com")
	}
	fmtoutput.PrintMachineHeader()

	m := mdao.GetMachineByHost(h)
	fmt.Println("Machine Status:")

	fmtoutput.PrintMachine(m)
	var pos uint8
	if s == "redis" {
		pos = utils.REDIS_POS
	} else if s == "redisproxy" {
		pos = utils.REDISPROXY_POS
	} else if s == "memcache" {
		pos = utils.MEMCACHE_POS
	} else {
		os.Exit(0)
	}
	res := machine.BanMachine(h, pos)
	if res == false {
		log.Info("Operation failed")
	} else {
		log.Info("Operation finished")
	}
	fmt.Println("Machine Status:")
	m = mdao.GetMachineByHost(h)
	fmtoutput.PrintMachine(m)
}

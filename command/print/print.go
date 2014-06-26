package print

import (
	"../../modules/gather/machine"
	"../../modules/gather/service"
	//"../../utils"
	"../../modules/fmtoutput"
	"github.com/codegangsta/cli"
)

var (
	Command = cli.Command{
		Name:      "print",
		ShortName: "p",
		Usage:     "print",
		Action:    printAction,
		Flags: []cli.Flag{
			cli.StringFlag{"t", "", "must be one of machine,service"},
		},
		Description: Usage,
	}

	Usage = `
NAME:
  print the machines and services
`
)

func printAction(c *cli.Context) {
	t := c.String("t")
	if t == "machine" {
		machines := machine.LoadMachine()
		fmtoutput.PrintMachineHeader()
		for _, m := range *machines {
			fmtoutput.PrintMachine(&m)
		}
	} else if t == "service" {
		services := service.LoadService()
		fmtoutput.PrintServiceHeader()
		for _, s := range *services {
			fmtoutput.PrintService(&s)
		}
	}
}

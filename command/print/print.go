package print

import (
	"../../modules/update/machine"
	"../../modules/update/service"
	//"../../utils"
	"../../modules/fmtoutput"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

var (
	Command = cli.Command{
		Name:      "print",
		ShortName: "p",
		Usage:     "print",
		Action:    printAction,
		Flags: []cli.Flag{
			cli.StringFlag{"t", "", "must be one of machine,service"},
			cli.IntFlag{"p", 0, "print used over the percentage services"},
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
	p := c.Int("p")
	if p < 0 || p > 100 {
		fmt.Println("-p must be 0--100")
		os.Exit(0)
	}
	if t == "machine" {
		machines := machine.LoadMachine()
		fmtoutput.PrintMachineHeader()
		if p != 0 {
			for _, m := range *machines {
				if float32(m.Mem.Free)/float32(m.Mem.Total) > float32(p)/float32(100) {
					fmtoutput.PrintMachine(&m)
				}
			}
		} else {
			for _, m := range *machines {
				fmtoutput.PrintMachine(&m)
			}
		}
	} else if t == "service" {
		services := service.LoadService()
		fmtoutput.PrintServiceHeader()
		if p != 0 {
			for _, s := range *services {
				if float32(s.UsedMem)/float32(s.BoxMem) > float32(p)/float32(100) {
					fmtoutput.PrintService(&s)
				}
			}
		} else {
			for _, s := range *services {
				fmtoutput.PrintService(&s)
			}
		}
	}
}

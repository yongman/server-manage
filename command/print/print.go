package print

import (
	//"../../modules/alloc"
	//mdao "../../modules/db/machine"
	"../../modules/gather/machine"
	//"../../modules/gather/service"
	//"../../utils"
	"fmt"
	"github.com/codegangsta/cli"
)

var (
	Command = cli.Command{
		Name:        "print",
		ShortName:   "p",
		Usage:       "print",
		Action:      printAction,
		Flags:       []cli.Flag{},
		Description: Usage,
	}

	Usage = `
NAME:
  print the machines and services commits
`
)

func printAction(c *cli.Context) {

	//machine.UpdateMachine("/home/users/yanming02/workspace/server-manage/host_mem.info")

	machines := machine.LoadMachine()
	for _, m := range *machines {
		fmt.Println(m)
	}
	fmt.Println("=====")

	//service.UpdateService("/home/users/yanming02/workspace/server-manage/host_redis.info")
	/*services := service.LoadService()
	for _, s := range *services {
		//mem := mdao.GetMemByHost(s.Host)
		fmt.Println("===>", s)
		fmt.Println("===>", utils.GetPidByDir(s.DirName))
	}
	*/
}

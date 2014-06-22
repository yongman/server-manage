package test

import (
	mdao "../../modules/db/machine"
	//"../../modules/gather/machine"
	//"../../modules/gather/service"
	//"../../utils"
	"../../utils"
	"fmt"
	"github.com/codegangsta/cli"
	//"labix.org/v2/mgo/bson"
)

var (
	Command = cli.Command{
		Name:        "test",
		ShortName:   "t",
		Usage:       "test",
		Action:      testAction,
		Flags:       []cli.Flag{},
		Description: Usage,
	}

	Usage = `
NAME:
  fmt - format proxy.conf
`
)

func testAction(c *cli.Context) {

	//machine.UpdateMachine("/home/users/yanming02/workspace/server-manage/host_mem.info")

	//machines := machine.LoadMachine()
	//for _, m := range *machines {
	//		fmt.Println(m)
	//	}
	fmt.Println("=====")
	mem := mdao.GetMemByHost("yf-arch-redis-wise12.yf01.baidu.com")
	fmt.Println(mem)
	fmt.Println(utils.GetIDCByHost("yf-arch-redis-wise12.yf01.baidu.com"))
	fmt.Println(utils.GetLogicByHost("yf-arch-redis-wise12.yf01.baidu.com"))
	fmt.Println(utils.GetRegionByHost("yf-arch-redis-wise12.yf01.baidu.com"))
	fmt.Println(utils.RandPercent())
	fmt.Println(utils.RandPercent())
	fmt.Println(len(*mdao.GetHostsByIDC("bj")))
	/*
		//service.UpdateService("/home/users/yanming02/workspace/server-manage/host_redis.info")
		services := service.LoadService()
		for index, s := range *services {
			//mem := mdao.GetMemByHost(s.Host)
			fmt.Println(index, "===>", s.Host)
		}
	*/

}

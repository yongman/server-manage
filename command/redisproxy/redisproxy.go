package redisproxy

import (
	//mdao "../../modules/db/machine"
	"../../modules/alloc"
	//"../../modules/db"
	"../../modules/log"
	"../../utils"
	"github.com/codegangsta/cli"
	"os"
	"sort"
)

var (
	Command = cli.Command{
		Name:      "redisproxy",
		ShortName: "rp",
		Usage:     "redisproxy tool",
		Action:    redisproxyAction,
		Flags: []cli.Flag{
			cli.StringFlag{"d", "", "pid, if provided, will avoid the servers that have deployed instances of this pid"},
			cli.IntFlag{"bj", 0, "amount alloc from machineroom in beijing"},
			cli.IntFlag{"hz", 0, "amount alloc from machineroom in hangzhou"},
			cli.IntFlag{"nj", 0, "amount alloc from machineroom in nanjing"},
			cli.StringFlag{"cid", "", "the commit id to drop"},
		},
		Description: Usage,
	}
	Usage = `
Name:
redisproxy alloc tools
        `
)

//用于排序
type freeN struct {
	index int
	free  float64
}

type freeNs []freeN

func (f freeNs) Len() int {
	return len(f)
}
func (f freeNs) Less(i, j int) bool {
	return f[i].free > f[j].free
}
func (f freeNs) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

//区域结构，用于循环处理
type region struct {
	cnt  int
	name string
}

func redisproxyAction(c *cli.Context) {
	pid := c.String("d")

	bj := c.Int("bj")
	hz := c.Int("hz")
	nj := c.Int("nj")
	regions := []region{
		{bj, "bj"},
		{hz, "hz"},
		{nj, "nj"},
	}
	//用于存放每个区域的分配结果，用于确认提交
	resultall := make([][]alloc.Instance, len(regions))
	for r_idx, r := range regions {
		//存放分配结果
		results := []alloc.Instance{}
		log.Info(r)
		if r.cnt != 0 {
			mach := alloc.AllocRedisproxyMachine(r.name, pid)
			fns := make(freeNs, len(*mach))
			for idx, ma := range *mach {
				fns[idx] = freeN{idx, float64(ma.Mem.Free) / float64(ma.Mem.Total)}
				log.Info(ma)
				//log.Info(ma.Host, "Getting a port....")
				//log.Info(alloc.AllocPort(ma.Host, REDIS_NAME))
			}

			log.Info("Total:", len(*mach), len(fns))
			//对列表进行按redis策略进行排序
			sort.Sort(fns)
			if len(fns) < r.cnt {
				log.Info("Not Enough Machine")
				os.Exit(1)
			}
			for i, fn := range fns {
				loop := r.cnt
				if i >= loop {
					break
				}
				ma := (*mach)[fn.index]
				port := alloc.AllocPort(ma.Host, utils.REDIS_NAME)
				if port == -1 {
					loop++ //端口分配失败，则继续在候选机器中选择
					continue
				}
				res := alloc.Instance{ma, port}
				results = append(results, res)
				log.Info("Alloc:", res.IMachine.Host, res.IPort)
			}
			if len(results) < r.cnt {
				log.Info("Not Enough Port")
				os.Exit(1)
			}
			log.Info(len(results))
			resultall[r_idx] = results
		}
	}
	os.Exit(0)
}

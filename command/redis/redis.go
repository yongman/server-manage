package redis

import (
	//mdao "../../modules/db/machine"
	"../../modules/alloc"
	"../../modules/db"
	"../../modules/log"
	"../../utils"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"sort"
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
	REDIS_NAME = string("redis")
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

type region struct {
	cnt  int
	name string
}

func redisAction(c *cli.Context) {
	fmt.Println("redisAction")
	m := c.String("m")
	pid := c.String("d")

	bj := c.Int("bj")
	hz := c.Int("hz")
	nj := c.Int("nj")
	regions := []region{
		{bj, "bj"}, {hz, "hz"}, {nj, "nj"},
	}
	if m != "" {
		size := utils.CapToKB(m)
		for _, r := range regions {
			//存放分配结果
			results := []db.Machine{}
			if r.cnt != 0 {
				mach := alloc.AllocMachine(size, r.name, pid)
				fns := make(freeNs, len(*mach))
				for idx, ma := range *mach {
					fns[idx] = freeN{idx, float64(ma.Mem.Free) / float64(ma.Mem.Total)}

					log.Info(ma.Host, ma.Status, ma.Mem)
					//log.Info(ma.Host, "Getting a port....")
					log.Info(alloc.AllocPort(ma.Host, REDIS_NAME))
				}

				log.Info("Total:", len(*mach), len(fns))
				//对列表进行按redis策略进行排序
				sort.Sort(fns)
				if len(fns) < r.cnt {
					log.Info("Not Enough Machine")
					os.Exit(1)
				}
				for i, fn := range fns {
					if i >= r.cnt {
						break
					}
					res := (*mach)[fn.index]
					results = append(results, res)
					log.Info(res)
				}
				log.Info(len(results))
			}
		}
	}
}

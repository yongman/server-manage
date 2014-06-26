package redis

import (
	//mdao "../../modules/db/machine"
	"../../modules/alloc"
	//"../../modules/db"
	"../../modules/commit"
	"../../modules/fmtoutput"
	"../../modules/log"
	"../../utils"
	"../../utils/filter"
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
			cli.StringFlag{"action", "", "must be one of [list,alloc,drop]"},
			cli.StringFlag{"d", "", "pid, if provided, will avoid the servers that have deployed instances of this pid"},
			cli.StringFlag{"m", "", "memory size want to alloc, eg. 512M 1G"},
			cli.IntFlag{"bj", 0, "amount alloc from machineroom in beijing"},
			cli.IntFlag{"hz", 0, "amount alloc from machineroom in hangzhou"},
			cli.IntFlag{"nj", 0, "amount alloc from machineroom in nanjing"},
			cli.StringFlag{"cid", "", "the commit id to drop"},
			cli.IntFlag{"n", 0, "used by action list. list the n newest commits"},
		},
		Description: Usage,
	}
	Usage = `
Name:
redis alloc tools
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

func redisAction(c *cli.Context) {
	act := c.String("action")
	if act == "" {
		log.Fatal("--action must be assign[list,commit,drop]")
		os.Exit(0)
	}
	m := c.String("m")
	pid := c.String("d")

	bj := c.Int("bj")
	hz := c.Int("hz")
	nj := c.Int("nj")
	regions := []region{
		{bj, "bj"},
		{hz, "hz"},
		{nj, "nj"},
	}
	if act == "alloc" {
		if m == "" {
			log.Fatal("arg is invalid:", "-m 1G [-bj 2] [-nj 3] [-hz 5]")
			os.Exit(0)
		}
		size := utils.CapToKB(m)
		//用于存放每个区域的分配结果，用于确认提交
		resultall := make([][]alloc.Instance, len(regions))
		for r_idx, r := range regions {
			//存放分配结果
			results := []alloc.Instance{}
			if r.cnt != 0 {
				mach := alloc.AllocRedisMachine(size, r.name, pid)
				fns := make(freeNs, len(*mach))

				//fmtoutput.PrintMachineHeader()
				for idx, ma := range *mach {
					fns[idx] = freeN{idx, float64(ma.Mem.Free) / float64(ma.Mem.Total)}
					//log.Info(ma.Host, ma.Status, ma.Mem)
				}

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
					//fmtoutput.PrintAlloc(i+1, &ma, port)
				}
				if len(results) < r.cnt {
					log.Info("Not Enough Port")
					os.Exit(1)
				}
				resultall[r_idx] = results
			}
		}
		//打印分配结果
		fmtoutput.PrintAllocHeader()
		var idx int = 1
		for _, cra := range resultall {
			for _, cr := range cra {
				fmtoutput.PrintAlloc(idx, &(cr.IMachine), cr.IPort)
				idx++
			}
		}

		var comm string = "no"
		fmt.Println("Commit:Yes|no [default:no]")
		fmt.Scanf("%s", &comm)
		if comm == "Yes" {
			for _, cra := range resultall {
				for _, cr := range cra {
					commit.DoCommit(cr.IMachine.Host, cr.IPort, filter.GetBoxType(size))
					//log.Info(cr)
				}
			}
		}
		//打印功能
	} else if act == "list" {
		n := c.Int("n")
		commit.ListCommit(n)
	} else if act == "drop" {
		cid := c.String("cid")
		if cid == "" {
			log.Fatal("arg error: -cid should be specified")
			os.Exit(0)
		}
		err := commit.DropCommit(cid)
		if err != nil {
			log.Fatal("Drop Commit Failed")
			log.Fatal(err)
		}
	}
	os.Exit(0)
}

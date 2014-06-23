package machine

import (
	"../../../utils"
	"../../db"
	"../../log"
	"bufio"
	"io"
	"labix.org/v2/mgo/bson"
	"os"
	"strconv"
	"strings"
)

func machineType(mem int64) (mtype string) {
	if mem > 128*1000*1000 {
		mtype = "T128G"
	} else if mem > 96*1000*1000 {
		mtype = "T96G"
	} else if mem > 64*1000*1000 {
		mtype = "T64G"
	} else if mem > 48*1000*1000 {
		mtype = "T48G"
	} else {
		mtype = "unknown"
	}
	return mtype
}

func UpdateMachine(rawfile string) int {
	f, err := os.Open(rawfile)
	if err != nil {
		log.Fatal("open file failed")
		os.Exit(1)
	}
	defer f.Close()
	br := bufio.NewReader(f)

	mongo := db.ClientDefault()
	instance, err := mongo.GetDB()
	if err != nil {
		log.Fatal("GetDB failed")
	}
	m_collec := instance.C("machine")
	m_collec.DropCollection()

	//格式
	//10.36.115.16 64423 43323 yf-arch-redis41.yf01.baidu.com
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			//一行数据
			line = strings.TrimFunc(line, utils.HasNewLine)
			xs := strings.Fields(line)

			mem := db.M_Mem{}
			var res int
			var err error

			res, err = strconv.Atoi(xs[1])
			if err != nil {
				log.Fatal("can not convert string", xs[1], "to int")
				os.Exit(1)
			}
			mem.Total = (int64)(res * 1024)

			res, err = strconv.Atoi(xs[2])
			if err != nil {
				log.Fatal("can not convert string", xs[2], "to int")
				os.Exit(1)
			}
			mem.Free = (int64)(res * 1024)
			mem.Cached = -1 //reserved

			mem.Percent = 0.90 //可以占用总内存的90%
			mem.Box10G, mem.Box5G, mem.Box1G = utils.DivideBox(int64(float32(mem.Free) * mem.Percent))
			machine := db.Machine{}
			machine.Mtype = machineType(mem.Total)
			machine.Host = xs[3]
			machine.Status = true
			machine.Mem = mem

			machine.Idc = utils.GetIDCByHost(machine.Host)
			machine.Logic = utils.GetLogicByHost(machine.Host)
			machine.Region = utils.GetRegionByHost(machine.Host)

			m_collec.Insert(&machine)
		}
	}

	count, err := m_collec.Count()
	if err != nil {
		log.Fatal("Count failed")
		os.Exit(1)
	}
	return count
}

func LoadMachine() *[]db.Machine {
	mongo := db.ClientDefault()
	instance, err := mongo.GetDB()
	if err != nil {
		log.Fatal("GetDB failed")
	}
	defer mongo.Close()
	m_collec := instance.C("machine")

	machines := &[]db.Machine{}
	m_collec.Find(bson.M{}).All(machines)
	return machines
}

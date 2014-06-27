package machine

import (
	"../../../utils"
	"../../db"
	mdao "../../db/machine"
	"../../log"
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func machineType(mem int64) (mtype string) {
	if mem > 128*1000*1000 {
		mtype = utils.MACHINE_128
	} else if mem > 96*1000*1000 {
		mtype = utils.MACHINE_96
	} else if mem > 64*1000*1000 {
		mtype = utils.MACHINE_64
	} else if mem > 48*1000*1000 {
		mtype = utils.MACHINE_48
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
			machine.Status = 0
			machine.Mem = mem

			//for test
			//CPU
			cpu := db.M_CPU{utils.GenerateCPUCore(machine.Mtype)}
			//Net
			net := utils.GenerateNet(machine.Mtype)
			//Disk
			disk := utils.GenerateDisk(machine.Mtype)
			machine.Cpu, machine.Net, machine.Disk = cpu, net, disk
			machine.Idc = utils.GetIDCByHost(machine.Host)
			machine.Mroom = utils.GetRoomByHost(machine.Host)
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
	return mdao.GetAllMachine()
}

//手动调整主机的内存box
func UpdateMachineBox(host string, b1g int8, b5g int8, b10g int8) error {
	m := mdao.GetMachineByHost(host)
	if m != nil {
		if b1g >= 0 {
			m.Mem.Box1G = b1g
		}
		if b5g >= 0 {
			m.Mem.Box5G = b5g
		}
		if b10g >= 0 {
			m.Mem.Box10G = b10g
		}
		mdao.UpdateMachineMem(m)
	}
	return nil
}

//将机器对具体的服务进行封禁
func BanMachine(host string, pos uint8) bool {
	if pos < 0 || pos > 31 {
		return false
	}
	m := mdao.GetMachineByHost(host)
	if m != nil {
		m.Status = m.Status | 1<<pos
		mdao.UpdateMachineStatus(m)
		return true
	} else {
		return false
	}
}

//解封
func UnBanMachine(host string, pos uint8) bool {
	if pos < 0 || pos > 31 {
		return false
	}
	m := mdao.GetMachineByHost(host)
	if m != nil {
		m.Status = m.Status & ^(1 << pos)
		mdao.UpdateMachineStatus(m)
		return true
	} else {
		return false
	}
}

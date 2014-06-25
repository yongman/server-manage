package alloc

import (
	"../../modules/db"
	mdao "../../modules/db/machine"
	sdao "../../modules/db/service"
	"../../utils"
	"../../utils/filter"
	"../log"
	"fmt"
	"net"
)

//分配结果结构
type Instance struct {
	IMachine db.Machine
	IPort    int32
}

//根据要求从指定的机房选择符合要求的count台redis机器
//size:分配机器大小
//region:区域
//pid:防止相同的pid出现在同一机器
//return 主机名列表
func AllocRedisMachine(size int64, region string, pid string) *[]db.Machine {
	machines := mdao.GetMachineByRegion(region)
	result := []db.Machine{}
	for _, m := range *machines {
		mem := m.Mem
		//主机被封禁
		if m.Status == false {
			continue
		}
		//redis数目检查
		if filter.FilterRedisNum(m.Host, 30) {
			continue
		}

		//内存限制filter
		if filter.FilterMem(utils.REDIS_NAME, &mem, size) {
			continue
		}
		//CPU限制filter
		cpu := m.Cpu
		if filter.FilterCPU(utils.REDIS_NAME, m.Host, &cpu) {
			continue
		}
		//网络限制filter
		net := m.Net
		if filter.FilterNet(utils.REDIS_NAME, &net) {
			continue
		}
		//磁盘限制filter
		disk := m.Disk
		if filter.FilterDisk(utils.REDIS_NAME, &disk) {
			continue
		}

		//query this machine has a particular pid or not
		//this used to avoid operation aggregation
		if pid != "" && sdao.PidInHost(pid, m.Host) {
			continue
		}
		//append a machine to array
		result = append(result, m)
	}
	return &result
}

//根据要求从指定的机房选择符合要求的count台redisproxy机器
//return 主机名列表
func AllocRedisproxyMachine(region string, pid string) *[]db.Machine {
	machines := mdao.GetMachineByRegion(region)
	result := []db.Machine{}
	for _, m := range *machines {
		log.Info(m.Host)
		mem := m.Mem
		//主机被封禁
		if m.Status == false {
			continue
		}
		//CPU限制filter
		cpu := m.Cpu
		if filter.FilterCPU(utils.REDISPROXY_NAME, m.Host, &cpu) {
			log.Info(m.Host, "cpu filter")
			continue
		}
		//网络限制filter
		net := m.Net
		if filter.FilterNet(utils.REDISPROXY_NAME, &net) {
			log.Info(m.Host, "net filter")
			continue
		}
		//磁盘限制filter
		disk := m.Disk
		if filter.FilterDisk(utils.REDISPROXY_NAME, &disk) {
			log.Info(m.Host, "disk filter")
			continue
		}
		//内存限制filter
		if filter.FilterMem(utils.REDISPROXY_NAME, &mem, 0) {
			log.Info(m.Host, "mem filter")
			continue
		}

		//query this machine has a particular pid or not
		//this used to avoid operation aggregation
		if pid != "" && sdao.PidInHost(pid, m.Host) {
			continue
		}
		//append a machine to array
		result = append(result, m)
	}
	return &result
}

//用于判断host机器port是否占用
func canDial(host string, port int32) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

//服务端口范围
var (
	REDIS_PORT_RANGE      = []int32{9100, 9999}
	REDISPROXY_PORT_RANGE = []int32{7000, 7999}
	//TODO add other services port range here
)

//根据host返回可用的空闲端口
//host:主机名
//stype:服务类型
//return 端口号
func AllocPort(host string, stype string) int32 {
	if stype == utils.REDIS_NAME {
		for port := REDIS_PORT_RANGE[0]; port <= REDIS_PORT_RANGE[1]; port++ {
			if canDial(host, port) == false {
				return port
			}
		}
	} else if stype == utils.REDISPROXY_NAME {
		for port := REDISPROXY_PORT_RANGE[0]; port <= REDISPROXY_PORT_RANGE[1]; port++ {
			if canDial(host, port) == false {
				return port
			}
		}
	}
	return -1
}

//对指定host对应的盒子计数减一
func DecBoxOne(host string, boxtype string) (err error) {
	m := mdao.GetMachineByHost(host)
	if boxtype == utils.Box10G {
		m.Mem.Box10G--
	} else if boxtype == utils.Box5G {
		m.Mem.Box5G--
	} else if boxtype == utils.Box1G {
		m.Mem.Box1G--
	}
	return mdao.UpdateMachineMem(m)
}

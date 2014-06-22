package alloc

import (
	"../../modules/db"
	mdao "../../modules/db/machine"
	"../../utils"
	"../log"
)

//根据要求从指定的机房选择符合要求的count台机器
//size:分配机器大小
//region:区域
//count:机器数目
//pid:防止相同的pid出现在同一机器
//return 主机名列表
func AllocMachine(size int64, region string, count int8, pid string) string {
	machines := mdao.GetMachineByRegion(region)
	//machine的指针数组
	result := []*db.Machine{}
	for _, m := range *machines {
		mem := m.Mem
		//内存限制filter
		if utils.FilterMem(&mem, size) {
			continue
		}

		//CPU限制filter
		cpu := m.Cpu
		if utils.FilterCPU(&cpu) {
			continue
		}
		//网络限制filter
		net := m.Net
		if utils.FilterNet(&net) {
			continue
		}
		//磁盘限制filter
		disk := m.Disk
		if utils.FilterDisk(&disk) {
			continue
		}

		//query this machine has a particular pid or not
		//this used to avoid operation aggregation
		if pid != "" && utils.PidInHost(pid, m.Host) {
			continue
		}
		//append a machine to array
		result = append(result, &m)
	}
	//result数组存放了符合条件的主机指针
	for _, x := range result {
		log.Info(x)
	}
	log.Info(len(result))
	return ""
}

//根据host返回可用的空闲端口
//host:主机名
//return 端口号
func AllocPort(host string) int32 {
	return 0
}

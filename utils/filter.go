package utils

import (
	"../modules/db"
)

func FilterMem(m *db.M_Mem, size int64) bool {
	if m.Total-m.Allocated > size && (float32(m.Allocated+size)/float32(m.Total)) < m.Percent {
		//符合条件，不应该对其过滤，返回false
		return false
	} else {
		return true
	}
}

func FilterCPU(c *db.M_CPU) bool {
	return false
}

func FilterNet(n *db.M_Net) bool {
	return false
}

func FilterDisk(d *db.M_Disk) bool {
	return false
}

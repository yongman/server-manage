package utils

import (
	"../modules/db"
	"../modules/log"
	"os"
)

var (
	M10G int64 = 10 * 1024 * 1024
	M5G  int64 = 5 * 1024 * 1024
	M1G  int64 = 1 * 1024 * 1024
)

func FilterMem(m *db.M_Mem, size int64) bool {
	/*
		if m.Total-m.Allocated > size && (float32(m.Allocated+size)/float32(m.Total)) < m.Percent {
			//符合条件，不应该对其过滤，返回false
			//log.Info("Total:", m.Total, "Allocated:", m.Allocated, "Free:", m.Total-m.Allocated, "size:", size)
			log.Info(float32(m.Allocated+size)/float32(m.Total), m.Percent)
			return false
		} else {
			return true
		}
	*/
	if size > M10G {
		log.Fatal("size too large")
		os.Exit(1)
	} else if size <= M10G && size > M5G {
		if m.Box10G > 0 {
			return false
		}
		return true
	} else if size <= M5G && size > M1G {
		if m.Box5G > 0 {
			return false
		}
		return true
	} else {
		if m.Box1G > 0 {
			return false
		}
		return true
	}
	return true
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

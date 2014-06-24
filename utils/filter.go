package utils

import (
	"../modules/db"
	"../modules/log"
	"os"
)

var (
	M10G   int64  = 10 * 1024 * 1024
	M5G    int64  = 5 * 1024 * 1024
	M1G    int64  = 1 * 1024 * 1024
	Box10G string = "M10G"
	Box5G  string = "M5G"
	Box1G  string = "M1G"
)

func GetBoxType(size int64) string {
	if size > M10G {
		log.Fatal("size too large")
		os.Exit(1)
	} else if size <= M10G && size > M5G {
		return Box10G
	} else if size <= M5G && size > M1G {
		return Box5G
	} else {
		return Box1G
	}
	return ""
}

//return filter
func FilterMem(m *db.M_Mem, size int64) bool {
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

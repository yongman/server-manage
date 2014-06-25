package utils

import (
	"../modules/db"
	"math/rand"
	"time"
)

func GenerateCPUCore(mtype string) int8 {
	if mtype == MACHINE_128 {
		return 16
	} else if mtype == MACHINE_96 {
		return 8
	} else if mtype == MACHINE_64 {
		return 4
	} else {
		return 2
	}
}
func RandPercent() (p float32) {
	rand.Seed(time.Now().UnixNano())
	p = float32(rand.Intn(65)) / 100.0
	return
}

func RandInt(m int, n int) int32 {
	rand.Seed(time.Now().UnixNano())
	return int32(rand.Intn(n) + m)
}

func GenerateNet(mtype string) db.M_Net {
	n := db.M_Net{}
	if mtype == MACHINE_128 || mtype == MACHINE_96 {
		n.Ntype = NETCARD_10G
		n.Up = 100 * 1000 * RandInt(2, 6)
		n.Down = 100 * 1000 * RandInt(2, 8)
	} else {
		n.Ntype = NETCARD_1G
		n.Up = 10 * 1000 * RandInt(2, 6)
		n.Down = 100 * 1000 * RandInt(2, 8)
	}
	return n
}

func GenerateDisk(mtype string) db.M_Disk {
	d := db.M_Disk{}
	if mtype == MACHINE_128 || mtype == MACHINE_96 {
		d.Total = 100 * 1024 * 1024
		d.Free = int64(float32(d.Total) * RandPercent())
	} else {
		d.Total = 10 * 1024 * 1024
		d.Free = int64(float32(d.Total) * RandPercent())
	}
	return d
}

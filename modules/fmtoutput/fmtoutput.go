package fmtoutput

import (
	"../db"
	"fmt"
	"strings"
)

func PrintCommitHeader() {
	fmt.Printf("%32s\t%40s\t%8s\t%10s\n", "CommitID", "Host", "Port", "AllocBox")
}
func PrintCommit(c *db.Commit) {
	fmt.Printf("%32s\t%40s\t%8d\t%10s\n", c.CommitID, c.Host, c.Port, c.AllocBox)
}

func PrintMachineHeader() {
	fmt.Printf("%5s%30s%10s%7s%5s%5s%5s%15s%15s%5s%5s%5s\n",
		"Type", "Host", "Status", "RM", "L", "RE", "IDC", "Total_Mem(KB)", "Free_Mem(KB)", "B1G", "B5G", "B10G")
}
func PrintMachine(m *db.Machine) {
	fmt.Printf("%5s%30s%10t%7s%5s%5s%5s%15d%15d%5d%5d%5d\n",
		m.Mtype, strings.TrimRight(m.Host, ".baidu.com"), m.Status, m.Mroom, m.Logic, m.Region, m.Idc, m.Mem.Total, m.Mem.Free, m.Mem.Box1G, m.Mem.Box5G, m.Mem.Box10G)

}

func PrintAllocHeader() {
	fmt.Println("Alloc Result:")
}
func PrintAlloc(index int, m *db.Machine, port int32) {
	fmt.Printf("%3d: %s:%d\n", index, strings.TrimRight(m.Host, ".baidu.com"), port)
}

func PrintServiceHeader() {
	fmt.Printf("%10s%30s%8s%15s%15s%50s%8s%10s\n",
		"Type", "Host", "Port", "Box", "Used", "DirName", "Role", "Version")
}
func PrintService(s *db.Service) {
	fmt.Printf("%10s%30s%8d%15d%15d%50s%8s%10s\n",
		s.Stype, strings.TrimRight(s.Host, ".baidu.com"), s.Port, s.BoxMem, s.UsedMem, s.DirName, s.Role, s.Version)
}

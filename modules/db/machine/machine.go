package machine

import (
	"../../db"
	"../../log"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
	//"strings"
)

func getCollection() *mgo.Collection {
	mongo := db.ClientDefault()
	instance, err := mongo.GetDB()
	if err != nil {
		log.Fatal("GetDB failed")
		os.Exit(1)
	}

	m_collec := instance.C("machine")
	return m_collec
}

func GetAllMachine() *[]db.Machine {
	m_collec := getCollection()

	machines := &[]db.Machine{}
	m_collec.Find(nil).All(machines)
	return machines
}

//get the memory infomation by hostname
func GetMemByHost(hostname string) *db.M_Mem {
	m_collec := getCollection()
	machine := db.Machine{}
	m_collec.Find(bson.M{"host": hostname}).One(&machine)
	mem := &(machine).Mem
	return mem
}

func GetMachineByHost(hostname string) *db.Machine {
	m_collec := getCollection()
	machine := db.Machine{}
	err := m_collec.Find(bson.M{"host": hostname}).One(&machine)
	if err != nil {
		return nil
	}
	return &machine
}

//
func GetHostsByIDC(idc string) *[]string {
	m_collec := getCollection()
	machines := []db.Machine{}
	m_collec.Find(bson.M{"idc": idc}).All(&machines)
	hosts := []string{}
	for _, m := range machines {
		host := m.Host
		hosts = append(hosts, host)
	}
	return &hosts
}

func GetMachineByRegion(region string) *[]db.Machine {
	m_collec := getCollection()
	machines := []db.Machine{}
	m_collec.Find(bson.M{"region": region}).All(&machines)
	return &machines
}

func UpdateMachineMem(m *db.Machine) (err error) {
	m_collec := getCollection()
	return m_collec.Update(bson.M{"_id": m.ID}, bson.M{"$set": bson.M{"mem": m.Mem}})
}

func UpdateMachineStatus(m *db.Machine) (err error) {
	m_collec := getCollection()
	return m_collec.Update(bson.M{"_id": m.ID}, bson.M{"$set": bson.M{"status": m.Status}})
}

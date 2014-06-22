package machine

import (
	"../../db"
	"../../log"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

func getCollection() *mgo.Collection {
	mongo := db.NewMongoDefault()
	instance, err := mongo.GetDB()
	if err != nil {
		log.Fatal("GetDB failed")
		mongo.Close()
		os.Exit(1)
	}

	m_collec := instance.C("machine")
	return m_collec
}

//get the memory infomation by hostname
func GetMemByHost(hostname string) *db.M_Mem {
	m_collec := getCollection()
	machine := db.Machine{}
	m_collec.Find(bson.M{"host": hostname}).One(&machine)
	mem := &(machine).Mem
	return mem
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

package service

import (
	"../../../utils"
	"../../db"
	"../../log"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"os"
)

func getCollection() *mgo.Collection {
	mongo := db.ClientDefault()
	instance, err := mongo.GetDB()
	if err != nil {
		log.Fatal("GetDB failed", err)
		os.Exit(1)
	}
	return instance.C("service")
}

//返回pid是否在host上
func PidInHost(pid string, host string) bool {
	s_collec := getCollection()
	services := []db.Service{}
	s_collec.Find(bson.M{"host": host}).All(&services)
	for _, s := range services {
		if pid == utils.GetPidByDir(s.DirName) {
			return true
		}
	}
	return false
}

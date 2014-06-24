package commit

import (
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
		log.Fatal("GetDB failed")
		os.Exit(1)
	}
	c_collec := instance.C("commit")
	return c_collec
}

//添加提交记录
func AddCommit(c *db.Commit) error {
	c_collec := getCollection()
	return c_collec.Insert(c)
}

func GetHostBoxByID(uuid string) (host string, boxtype string, err error) {
	c_collec := getCollection()
	c := db.Commit{}
	err = c_collec.Find(bson.M{"commitid": uuid}).One(&c)
	if err != nil {
		return "", "", err
	}
	return c.Host, c.AllocBox, nil

}
func DropCommit(uuid string) error {
	c_collec := getCollection()
	return c_collec.Remove(bson.M{"commitid": uuid})
}

func QueryLatestCommit(n int32) *[]db.Commit {
	latest := []db.Commit{}
	c_collec := getCollection()
	c_collec.Find(nil).All(&latest)
	return &latest
}

func DropAll() error {
	c_collec := getCollection()
	return c_collec.DropCollection()
}

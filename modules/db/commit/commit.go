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

//返回最新的n次提交
func QueryLatestCommit(n int) *[]db.Commit {
	latest := []db.Commit{}
	c_collec := getCollection()
	if n < 0 {
		c_collec.Find(nil).Sort("-_id").All(&latest)
	} else {
		c_collec.Find(nil).Sort("-_id").Limit(n).All(&latest)
	}
	return &latest
}

/*
func DropAll() error {
	c_collec := getCollection()
	return c_collec.DropCollection()
}
*/
//return all commitid,used to delete all commits
func GetAllCommitID() (*[]string, error) {
	c_collec := getCollection()
	commits := []db.Commit{}
	err := c_collec.Find(nil).All(&commits)
	if err != nil {
		return nil, err
	}
	cids := make([]string, len(commits))
	for idx, commit := range commits {
		cids[idx] = commit.CommitID
	}
	return &cids, nil
}

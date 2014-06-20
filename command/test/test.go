package test

import (
	"../../modules/db"
	"../../utils"
	"fmt"
	"github.com/codegangsta/cli"
	"labix.org/v2/mgo/bson"
)

var (
	Command = cli.Command{
		Name:      "test",
		ShortName: "t",
		Usage:     "test",
		Action:    testAction,
		Flags: []cli.Flag{
			cli.BoolFlag{"yaml", "convert to yaml"},
			cli.BoolFlag{"json", "convert to json"},
		},
		Description: Usage,
	}

	Usage = `
NAME:
  fmt - format proxy.conf
`
)

type test1 struct {
	CommitID string
	Ip       string
	Port     int32
}

func testAction(c *cli.Context) {
	mongo := db.NewMongo("st01-arch-agent00.st01:27017", "test", "machine")
	fmt.Println(utils.GenerateUUID())

	instance, err := mongo.GetDB()
	if err != nil {
		panic(err)
	}

	instance.DropDatabase()
	cl := instance.C(mongo.Collection)

	cl.Insert(&test1{utils.GenerateUUID(), "192.168.1.2", 1234})

	fmt.Println(cl.Count())
	result := []test1{}
	cl.Find(bson.M{}).All(&result)

	fmt.Println(result)

	defer mongo.Close()
}

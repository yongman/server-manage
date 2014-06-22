package service

import (
	"../../../utils"
	"../../db"
	"../../log"
	"bufio"
	"io"
	"labix.org/v2/mgo/bson"
	"os"
	"strconv"
	"strings"
)

func UpdateService(rawfile string) int {
	f, err := os.Open(rawfile)
	if err != nil {
		log.Fatal("open file", rawfile, "failed")
		os.Exit(1)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	//文件格式
	//10.38.167.56 redis-mco-webrssfeed-shard9 9911 1.8G 680351624 648.83M 696057856 680349280 648.83M slave 2.4.17 redis-mco-webrssfeed-shard9
	//services := []db.Service{}
	mongo := db.ClientDefault()
	instance, err := mongo.GetDB()
	if err != nil {
		log.Fatal("GetDB failed")
		os.Exit(1)
	}

	s_collec := instance.C("service")

	s_collec.DropCollection()

	var port int
	var mem int
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			//一行数据
			line = strings.TrimFunc(line, utils.HasNewLine)
			xs := strings.Fields(line)
			//log.Info(len(xs), "splited", xs)
			service := db.Service{}
			prefix := strings.Split(xs[1], "-")[0]
			switch prefix {
			case "redisproxy":
				//redisproxy
				service.Stype = prefix
				service.Host = xs[3]
				port, _ = strconv.Atoi(xs[2])
				service.Port = int32(port)
				service.MaxMem = -1
				service.DirName = xs[1]
			case "redis":
				service.Stype = prefix
				service.Host = xs[11]
				port, _ = strconv.Atoi(xs[2])
				service.Port = int32(port)
				service.MaxMem = -1
				service.DirName = xs[1]
				mem, _ = strconv.Atoi(xs[4])
				service.UsedMem = int64(mem)
				service.Role = xs[9]
				service.Version = xs[10]
			default:
				continue
			}
			//services = append(service,service)
			s_collec.Insert(&service)
		}
	}
	count, err := s_collec.Count()
	return count
}

func LoadService() *[]db.Service {
	mongo := db.ClientDefault()
	instance, err := mongo.GetDB()
	if err != nil {
		log.Fatal("GetDB failed")
		os.Exit(1)
	}
	defer mongo.Close()

	services := &[]db.Service{}
	s_collec := instance.C("service")
	s_collec.Find(bson.M{}).All(services)
	return services
}

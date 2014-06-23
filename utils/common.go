package utils

import (
	"../modules/log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	idc_map = map[string]string{
		"tc":    "bj",
		"yf01":  "bj",
		"dbl01": "bj",
		"m1":    "bj",
		"cq01":  "bj",
		"st01":  "bj",
		"cq02":  "bj",
		"jx":    "bj",
		"ai01":  "bj",
		"db01":  "bj",
		"cp01":  "bj",
		"hz01":  "hd",
		"nj02":  "hd"}
	region_map = map[string]string{
		"tc":    "bj",
		"yf01":  "bj",
		"dbl01": "bj",
		"m1":    "bj",
		"cq01":  "bj",
		"st01":  "bj",
		"cq02":  "bj",
		"jx":    "bj",
		"ai01":  "bj",
		"db01":  "bj",
		"cp01":  "bj",
		"hz01":  "hz",
		"nj02":  "nj"}
	logic_map = map[string]string{
		"m1":    "tc",
		"tc":    "tc",
		"st01":  "tc",
		"cq02":  "tc",
		"db01":  "tc",
		"yf01":  "jx",
		"dbl01": "jx",
		"cq01":  "jx",
		"jx":    "jx",
		"ai01":  "jx",
		"cp01":  "jx",
		"hz01":  "hz",
		"nj02":  "nj"}
)

func HasNewLine(r rune) bool {
	return r == '\n' || r == '\r'
}

func GetIDCByHost(host string) string {
	sp := regexp.MustCompile("[-]|[.]").Split(host, -1)
	if len(sp) >= 3 {
		return idc_map[sp[len(sp)-3]]
	}
	return ""
}

func GetLogicByHost(host string) string {
	sp := regexp.MustCompile("[-]|[.]").Split(host, -1)
	if len(sp) >= 3 {
		return logic_map[sp[len(sp)-3]]
	}
	return ""
}

func GetRegionByHost(host string) string {
	sp := regexp.MustCompile("[-]|[.]").Split(host, -1)
	if len(sp) >= 3 {
		return region_map[sp[len(sp)-3]]
	}
	return ""
}

func RandPercent() (p float32) {
	rand.Seed(time.Now().UnixNano())
	p = float32(rand.Intn(65)) / 100.0
	return
}

func isAlph(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

//convert cap to KB
//1G 1GB 512M 512MB 1gb 512mb
func CapToKB(c string) int64 {
	if strings.HasSuffix(c, "G") ||
		strings.HasSuffix(c, "GB") ||
		strings.HasSuffix(c, "gb") ||
		strings.HasSuffix(c, "g") {
		d := strings.TrimRightFunc(c, isAlph)
		num, err := strconv.Atoi(d)
		if err != nil {
			log.Fatal("CapToKB error")
			os.Exit(1)
		}
		log.Info(num)
		return int64(num) * 1024 * 1024
	} else if strings.HasSuffix(c, "M") ||
		strings.HasSuffix(c, "m") ||
		strings.HasSuffix(c, "MB") ||
		strings.HasSuffix(c, "mb") {
		d := strings.TrimRightFunc(c, isAlph)
		num, err := strconv.Atoi(d)
		if err != nil {
			log.Fatal("CapToKB error")
			os.Exit(1)
		}
		log.Info(num)
		return int64(num) * 1024
	} else {
		log.Fatal("size arguments invalid")
		os.Exit(1)
	}
	return 0
}

var (
	prefixs = []string{
		"redis-",
		"redisproxy-",
		"memcached-",
	}
)

//根据dirname得到PID
func GetPidByDir(dir string) string {
	var pid string
	for _, p := range prefixs {
		if strings.HasPrefix(dir, p) {
			pid = strings.TrimPrefix(dir, p)
			//对redis做特殊处理，dirname后有分片信息
			if p == "redis-" {
				pos := strings.Index(pid, "-shard")
				if pos == -1 {
					continue
				}
				rs := []rune(pid)
				pid = string(rs[0:pos])
				//log.Info(pid, pos)
			}
			break
		}
	}
	return pid
}

//根据可用空间大小，进行盒子划分，并返回对应盒子的类型 for test
func DivideBox(mem int64) (box10G int8, box5G int8, box1G int8) {
	b10 := (mem / 2) / (10 * 1024 * 1024) //50%用来划分10G盒子
	mem = mem - 10*1024*b10

	b5 := (mem / 2) / (5 * 1024 * 1024) //25%用来划分5G盒子
	mem = mem - 5*1024*b5

	b1 := mem / (1024 * 1024) //剩余25%用来划分1G盒子
	return int8(b10), int8(b5), int8(b1)
}

package utils

import (
	"math/rand"
	"regexp"
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

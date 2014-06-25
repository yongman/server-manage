package filter

import (
	"../../modules/db"
	sdao "../../modules/db/service"
	"../../modules/log"
	"../../utils"
	"os"
)

func GetBoxType(size int64) string {
	if size > utils.M10G {
		log.Fatal("size too large")
		os.Exit(1)
	} else if size <= utils.M10G && size > utils.M5G {
		return utils.Box10G
	} else if size <= utils.M5G && size > utils.M1G {
		return utils.Box5G
	} else {
		return utils.Box1G
	}
	return ""
}

//return filter
func FilterMem(service string, m *db.M_Mem, size int64) bool {
	switch service {
	case utils.REDIS_NAME:
		if size > utils.M10G {
			log.Fatal("size too large")
			os.Exit(1)
		} else if size <= utils.M10G && size > utils.M5G {
			if m.Box10G > 0 {
				return false
			}
			return true
		} else if size <= utils.M5G && size > utils.M1G {
			if m.Box5G > 0 {
				return false
			}
			return true
		} else {
			if m.Box1G > 0 {
				return false
			}
			return true
		}
	case utils.REDISPROXY_NAME:
		return false
	default:
		return false
	}
	return false
}

func FilterCPU(service string, hostname string, c *db.M_CPU) bool {
	switch service {
	case utils.REDIS_NAME:
		return false
	case utils.REDISPROXY_NAME:
		//proxy对CPU条件过滤规则:proxy数量=cpu核心数目-2
		proxy_cnt, err := sdao.GetProxyNum(hostname)
		if err != nil {
			log.Fatal(err)
			return true
		}
		if proxy_cnt >= (int)(c.CoreNum)-2 {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func FilterNet(service string, n *db.M_Net) bool {
	switch service {
	case utils.REDIS_NAME:
		return false
	case utils.REDISPROXY_NAME:
		if n.Ntype == utils.NETCARD_10G {
			//按总带宽的80%过滤
			if (n.Up + n.Down) > (1000 * 1000) {
				return true
			} else {
				return false
			}
		} else if n.Ntype == utils.NETCARD_1G {
			if n.Up+n.Down > 100*1000 {
				return true
			} else {
				return false
			}
		}
		return false
	default:
		return false
	}
}

func FilterDisk(service string, d *db.M_Disk) bool {
	switch service {
	case utils.REDIS_NAME:
		return false
	case utils.REDISPROXY_NAME:
		if d.Free < 1024*10 {
			//desc in MB
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func FilterRedisNum(host string, max int) bool {
	num, err := sdao.GetRedisNum(host)
	if err != nil {
		log.Fatal(err)
		return true
	}
	log.Info(host, "Redis num:", num)
	if num >= max {
		return true
	} else {
		return false
	}
}

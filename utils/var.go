package utils

//Commom vars put in this file
var (
	REDIS_NAME      string = "redis"
	REDISPROXY_NAME string = "redisproxy"

	//内存描述相关
	M10G   int64  = 10 * 1024 * 1024
	M5G    int64  = 5 * 1024 * 1024
	M1G    int64  = 1 * 1024 * 1024
	Box10G string = "M10G"
	Box5G  string = "M5G"
	Box1G  string = "M1G"

	//网卡类型
	NETCARD_1G  string = "1G"
	NETCARD_10G string = "10G"

	//机器类型
	MACHINE_48  string = "T48G"
	MACHINE_64  string = "T64G"
	MACHINE_96  string = "T96G"
	MACHINE_128 string = "T128G"
)

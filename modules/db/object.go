package db

type M_Mem struct {
	//desc in KB
	Total  int64 //总内存大小
	Free   int64 //空闲内存
	Cached int64 //缓存内存

	Prealloc int64 //总预分配大小
	//memery box amount
	//Box1G  int8
	//Box5G  int8
	//Box10G int8
}

type M_CPU struct {
}

type M_Net struct {
}

type M_Disk struct {
}

//describe a machine's status
type Machine struct {
	Mtype  string  //机器类型
	Host   string  //主机名
	RNum   int8    //Redis部署数目
	RPNum  int8    //RedisProxys部署数目
	Mroom  string  //机房
	Logic  string  //逻辑机房
	Region string  //地域
	Idc    string  //IDC
	Mem    *M_Mem  //内存描述
	Cpu    *M_CPU  //CPU描述
	Net    *M_Net  //网络描述
	Disk   *M_Disk //硬盘描述
}

//describe a service
type Service struct {
	Mach   *Machine //服务所在机器
	Port   int32    //服务端口
	MaxMem int64    //服务申请的最大内存
	pid    string   //服务的PID
}

package cnode

var (
	Ip     = GetLocalIP()
	Port   = GetRandom(1, 30000)
	errStr = "registerToZk error:%+v"
	msg    = "centerHost is empty"
)

type HeartBeat struct {
	Ip         string `json:"ip"`
	Port       string `json:"port"`
	ExecuteId  int    `json:"execute_id"`  //执行器ID
	Cmd        int    `json:"cmd"`         //发送给服务器命令类型
	Flag       int    `json:"flag"`        //标识客户端还是服务器端 0客户端通信 1服务端通信
	Token      string `json:"token"`       //将来用的验证标识 目前预留
	LogId      int    `json:"log_id"`      //日志id
	OpFlag     int    `json:"op_flag"`     //标示如何操作数据库
	JobId      int    `json:"job_id"`      //任务id
	JobHandler string `json:"job_handler"` //任务handler
	ExeParam   string `json:"exe_param"`   //任务参数
	CronMsg    string `json:"cron_msg"`    //日志内容
}

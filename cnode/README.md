集成方式：
1. 导入包：
    go get github.com/drewinner/cnode
2. 代码： 
```
func main() {
	cnode.Register("TestHandler2",&jobs.CallBackStruct{})
        //@param:后台服务器地址，执行器id 每个实例对应一个执行器
        //@param:heatBeatTime 心跳时间 单位秒
	_ =cnode.StartDisRpcServer("127.0.0.1:8080,127.0.0.1:8082",1,"F:/logs/dis_push/business_log/",4)
}
```

```
// 继承CallBackBaseStruct结构体
type CallBackStruct struct {
	go-dis-cron-cli.CallBackBaseStruct
}
//处理业务逻辑的方法 -- 需要重写的方法
func(c *CallBackStruct)  CallBack(request *go-dis-cron-cli.JobRequest,hostUrl,logUrl string) (bool,error) {
	fmt.Println("selfCallBack CallBack....")
	return true
}
```


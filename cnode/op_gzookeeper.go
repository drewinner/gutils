package cnode

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	OnlineServer = make([]interface{}, 0) //存放本地在线服务器列表
	lock         sync.Mutex               //互斥锁
)

/**
*	清空onlineServer
 */
func CleanOnlineServer() {
	OnlineServer = make([]interface{}, 0)
}

/**
*	向slice中添加元素
*	@param:v 添加元素的值
 */
func OnlineServerAppend(v interface{}) {
	lock.Lock()
	defer lock.Unlock()
	OnlineServer = append(OnlineServer, v)
}

/**
*	注册服务到zk中,使服务器端能够获取客户端地址
*	@param c 管道，判断是否初始化完成
*	@param hosts zk地址 格式 192.168.0.1:2181,192.168.0.2:2181
*	@param path 注册的地址
*	@param executeId 执行器id
 */
func registerToZk(c chan bool, hosts, path string, executeId int) {
	errStr := "registerToZk error:%+v"
	if hosts == "" {
		panic(fmt.Sprintf("zk host empty"))
	}
	addr := strings.Split(hosts, ",")
	zkClient, err := NewZKClient(path, addr...)
	if err != nil {
		panic(fmt.Sprintf(errStr, err))
	}
	zkValue := new(HeartBeat)
	zkValue.Ip = Ip
	zkValue.Port = Port
	zkValue.ExecuteId = executeId

	zkByte, err := json.Marshal(zkValue)
	if err != nil {
		panic(fmt.Sprintf(errStr, err))
	}
	err = zkClient.Register(strconv.Itoa(executeId)+"_"+Ip+":"+Port, zkByte)
	if err != nil {
		panic(fmt.Sprintf(errStr, err))
	}
	c <- true
	select {}
}

/**
*	客户端监听zk目录元素
*	@param : hosts zk地址
*	@param : path 监听元素路径
 */
func zkOnline(hosts, path string, executeId int) {
	//捕获zk异常，1分钟后重试
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("zkonline error:%+v\n", err)
			time.Sleep(time.Second * 60)
			zkOnline(hosts, path, executeId)
		}
	}()
	if hosts == "" {
		panic(fmt.Sprintf("zk host empty"))
	}
	addr := strings.Split(hosts, ",")
	zkClient, err := NewZKClient(path, addr...)
	if err != nil {
		panic(fmt.Sprintf(errStr, err))
	}
	defer zkClient.Close()
	snapshots, zkerr := zkClient.Mirror()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case snapshot := <-snapshots:
			CleanOnlineServer() //清除OnlineServer列表
			for ip := range snapshot {
				OnlineServerAppend(ip)
			}
		case err := <-zkerr:
			panic(fmt.Sprintf(errStr, err))
		case _ = <-ch:
			_ = zkClient.Close()
			os.Exit(1)
		}

	}
}

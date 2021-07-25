package cnode

import (
	"encoding/json"
	"fmt"
	"net"
)

/**
*	启动tcp客户端
*	@param:centerHost 注册中心地址
*	@param:执行器id
*	@param:logUrl 客户端写入的日志文件地址
*	@param:heatBeatTime心跳时间 -- 目前没用
 */
func StartTcpCli(centerHost string, executeId int) {
	if centerHost == "" {
		panic(msg)
	}
	c := make(chan bool, 0)
	//将客户端注册到zk中
	go registerToZk(c, centerHost, "/dis_cron/clients", executeId)
	if <-c {
		fmt.Println("zk注册完成...")
	}
	//获取可用节点
	go zkOnline(centerHost, "/dis_cron/onlineserver", executeId)
	//启动客户端tcp服务
	go startTcpServer()
	//客户端向服务端发送更新日志操作
	go subLogChan()
	select {}
}

/**
*	处理消息日志订阅通道
 */
func subLogChan() {
	for {
		select {
		case b := <-LogChan:
			//通过tcp请求服务器端
			upLogReqByte, err := json.Marshal(b)
			if err != nil {
				continue
			}
			if c, b := getAvailableNode(); b {
				_, err = c.Write(Packet(upLogReqByte))
				if err != nil {
					fmt.Printf("sublog err:+%v\n", err)
				}
				if c != nil {
					_ = c.Close()
				}
			} else {
				fmt.Println("无可用节点")
			}
		}
	}
}

/**
*	获取可用节点
 */
func getAvailableNode() (conn *net.TCPConn, b bool) {
	var (
		isSelected = false
		c          *net.TCPConn
	)
	Shuffle(OnlineServer)
	for _, server := range OnlineServer {
		tcpAddr, err := net.ResolveTCPAddr("tcp", server.(string))
		if err != nil {
			continue
		}
		c, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			continue
		}
		//选到可用节点就退出
		isSelected = true
		break
	}
	return c, isSelected
}

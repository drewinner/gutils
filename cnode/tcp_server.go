package cnode

import (
	"encoding/json"
	"fmt"
	"net"
)

/**
*	启动客户端tcp服务
 */
func startTcpServer() {
	listen, err := net.Listen("tcp", Ip+":"+Port)
	if err != nil {
		panic(fmt.Sprintf("StartTcpServer err:%+v", err.Error()))
	}
	fmt.Printf("startTcpServer success ...%s:%s\n", Ip, Port)
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

/**
*	tcp server
 */

func handleConnection(conn net.Conn) {
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)

	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		tmpBuffer = unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

/**
*	channel中读取消息--处理具体业务逻辑
 */
func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			if len(data) == 0 {
				continue
			}
			h := new(HeartBeat)
			err := json.Unmarshal(data, h)
			if err != nil {
				continue
			}
			//请求接口
			go realCallBackTcp(h)
		}
	}
}

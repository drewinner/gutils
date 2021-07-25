package cnode

import (
	"math/rand"
	"net"
	"strconv"
	"time"
)

//获取本地ip
func GetLocalIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

//获取随机数
func GetRandom(flag, num int) string {
	rand.Seed(time.Now().UnixNano())
	switch flag {
	case 1:
		x := rand.Intn(num) + 30000
		return strconv.Itoa(x)
	case 2:
		x := rand.Intn(num)
		return strconv.Itoa(x)
	}
	return "0"
}

/**
*	slice乱序
 */
func Shuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

/**
*	获取当前格式化时间2017-04-11 13:24:04
*	@param:flag 0 返回时间格式为 2017-04-11 13:24:04
 */
func GetFormatTime(flag int) string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	switch flag {
	case 0:
		timeStr = time.Now().Format("2006-01-02 15:04:05")
	case 1:
		timeStr = time.Now().Format("20060102")
	}
	return timeStr
}

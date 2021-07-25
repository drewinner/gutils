package cnode

import (
	"errors"
	"fmt"
)

type CallBackInterface interface {
	//客户端回调函数
	/**
	*	return : msg 执行后的信息
	 */
	CallBack(heartBeat *HeartBeat) (msg string)
}
type CallBackBaseStruct struct {
}

var (
	LogChan = make(chan *HeartBeat, 50)
)

/**
*	回调前执行方法
*	效率原因，不进行重试
 */
func before(logId int) (bool, error) {
	if logId == 0 {
		return false, errors.New("更新日志参数错误")
	}
	LogChan <- &HeartBeat{
		LogId:  logId,
		OpFlag: 0,
		Cmd:    1,
	}
	return true, nil
}

/**
*	处理业务逻辑的方法
*	@param:request rcp请求体
*	@param:hostUrl 业务回调地址
*	@param:logUrl 日志存放地址
*	@return: msg 作为执行内容返回，记录日志
 */
func (c *CallBackBaseStruct) CallBack(heartBeat *HeartBeat) (msg string) {
	return "执行成功"
}

/**
*	回调执行的方法
*	更新cron_cron_log表执行开始时间
*	@param:logId 日志id
*	@param:msg 执行完之后的信息
 */
func after(logId int, msg string) (bool, error) {
	if logId == 0 {
		return false, errors.New("参数错误")
	}
	//将消息内容写入日志，返回相对路径
	//path, _ := WriteFileContent(GlobalLogUrl, logId, msg)
	LogChan <- &HeartBeat{
		Flag:    0,
		LogId:   logId,
		OpFlag:  1,
		Cmd:     1,
		CronMsg: msg,
	}
	return true, nil
}

var MethMap = make(map[string]CallBackInterface)

/**
*	注册需要执行的方法
 */
func Register(funcName string, backInterface CallBackInterface) {
	if _, ok := MethMap[funcName]; !ok {
		MethMap[funcName] = backInterface
	}
}

/**
*	tcp回调函数
 */
func realCallBackTcp(heartBeat *HeartBeat) {
	if heartBeat.JobHandler == "" {
		fmt.Println("JobHandler 为空,请联系管理员")
		return
	}
	//判断是否注册方法，没注册直接返回
	v, ok := MethMap[heartBeat.JobHandler]
	if !ok {
		fmt.Printf("%s 未注册\n", heartBeat.JobHandler)
		return
	}
	if b, _ := before(heartBeat.LogId); !b {
		return
	}
	msg := v.CallBack(heartBeat)
	_, _ = after(heartBeat.LogId, msg)
}

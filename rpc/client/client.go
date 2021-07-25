package client

import (
	"context"
	"fmt"
	pb "github.com/drewinner/gutils/rpc/proto/rpc"
	"google.golang.org/grpc"
	"time"
)

/**
*	调用服务端
*	@param:id int32 任务id
*	@param:logId 日志id
*	@param:taskHandler 任务标识
*	@param:params 参数
 */
func Invoke(id, logId int32, taskHandler, params string) (resp *pb.TaskResp, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "127.0.0.1:8090", grpc.WithInsecure())
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)
	r, err := client.Run(ctx, &pb.TaskReq{
		Id:         id,
		LogId:      logId,
		JobHandler: taskHandler,
		Params:     params,
	})
	if err != nil {
		fmt.Println(err.Error())

	}
	return r, err
}

package client

import (
	pb "github.com/drewinner/gutils/rpc/proto/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

// GetClient /**
/**
*	获取client
 */
func GetClient(address string) (pb.TaskServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewTaskServiceClient(conn)
	return c, nil
}

package server

import "context"
import pb "github.com/drewinner/gutils/rpc/proto/rpc"

type Server struct {
	pb.UnimplementedTaskServiceServer
}

func (s *Server) Call(ctx context.Context, req *pb.TaskReq) (*pb.TaskResp, error) {
	resp := &pb.TaskResp{
		Id:            1,
		LogId:         2,
		Status:        1,
		ExecStartTime: "",
		ExecEndTime:   "",
		LogMsg:        "yes....",
	}
	return resp, nil
}

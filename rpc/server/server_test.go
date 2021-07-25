package server

import (
	"fmt"
	pb "github.com/drewinner/gutils/rpc/proto/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func TestServer_Run(t *testing.T) {
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("server:", err.Error())
	}
	select {}
}

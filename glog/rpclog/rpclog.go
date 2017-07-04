package main

import (
	"fmt"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/glog"
	pb "github.com/evanchen/bamboo/rpcpto"
	_ "golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type rpcServe struct {
}

var g_serve *rpcServe
var g_grpcServer *grpc.Server

func newServer() *rpcServe {
	g_serve = &rpcServe{}
	return g_serve
}

func (s *rpcServe) SendLog(stream pb.RpcLog_SendLogServer) error {
	for {
		pLogInfo, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.RetSend{Ret: 1})
		}
		if err != nil {
			return err
		}
		glog.WriteFile(pLogInfo.Path, pLogInfo.Content)
	}
}

func main() {
	etc.LoadConfig()
	var opts []grpc.ServerOption
	g_grpcServer = grpc.NewServer(opts...)
	pb.RegisterRpcLogServer(g_grpcServer, newServer())
	ret, port := etc.GetConfigInt("log_server_port")
	if !ret {
		log.Fatalf("[log.main] log_server_port error\n")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("[log.main] failed to listen: %v", err)
	}
	fmt.Println("[log.main] start rpc serve listening...")
	g_grpcServer.Serve(lis)

	closech := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.Tick(5 * time.Second):
				fmt.Println("tick tack")
			}
		}
		close(closech)
	}()
	<-closech
}

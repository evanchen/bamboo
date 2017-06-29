package rpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	pb "github.com/evanchen/bamboo/rpcpto"
	glog "github.com/evanchen/bamboo/glog"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/base"
	"log"
	"net"
	"fmt"
)

var all_rpc_clients map[string] pb.RpcRouteClient

func StartRPC() {
	StartRPCServer()
	if base.GetGsId() > 0 {
		cName := fmt.Sprintf("game_%d_port",0)
		ret, port := etc.GetConfigInt(cName)
		if !ret {
			log.Fatalf("[StartRPC] %s config error: %d",cName,port)
		}
		addrPort := fmt.Sprintf("127.0.0.1:%d",port)
		go func(){
			for {
				ret := StartRPCClient(addrPort)
				if ret != nil {
					base.ServerReady()
					break 
				}
			}
		}()
	} else {
		go func () {
			base.ServerReady()	
		}()
	}
}

func StartRPCClient(addrPort string) pb.RpcRouteClient {
	v,ok := all_rpc_clients[addrPort]
	if ok {
		log.Printf("[StartRPCClient] try to connect: %s twice",addrPort)
		return nil 
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addrPort, opts...)
	if err != nil {
		log.Printf("[StartRPCClient] fail to dial: %v", err)
		return nil 
	}
	client := pb.NewRpcRouteClient(conn)
	all_rpc_clients[addrPort] = client
	return client
}

func StartRPCServer() {
	cName := fmt.Sprintf("game_%d_port",base.GetGsId())
	ret, port := etc.GetConfigInt(cName)
	if !ret {
		log.Fatalf("[StartRPCServer] %s config error: %d",cName,port)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("[StartRPCServer] failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRpcRouteServer(grpcServer, newServer())
	err := grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("[StartRPCServer] failed: %v",err)
	}
}

func Test() {
	cName := fmt.Sprintf("game_%d_port",0)
	ret, port := etc.GetConfigInt(cName)
	if !ret {
		log.Fatalf("[StartRPC] %s config error: %d",cName,port)
	}
	addrPort := fmt.Sprintf("127.0.0.1:%d",port)
	client,ok := all_rpc_clients[addrPort]
	if !ok {
		log.Fatalf("[Test] error")
	}
	client.GetInfo()
}
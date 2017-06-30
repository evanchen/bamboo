package rpc

import (
	"fmt"
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/proc"
	pb "github.com/evanchen/bamboo/rpcpto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// rpc 服务,每个进程只有一个
var g_grpcServer *grpc.Server

// rpc 请求,0进程有多个
var g_grpcConn = make(map[int]*grpc.ClientConn)

// 0进程启动时,需等待game进程的rpc连接完成,然后主动连通各game进程的rpc服务
// 每个进程既作为rpc的client,也作为server,但game进程只与0进程相互连接
func StartRPC() {
	InitRPCServer()
	if base.GetGsId() == 0 {
		return
	} else {
		// 其他进程必须先连接0进程
		cName := fmt.Sprintf("game_%d_port", base.MASTER_GSID)
		ret, port := etc.GetConfigInt(cName)
		if !ret {
			log.Fatalf("[StartRPC] %s config error: %d", cName, port)
		}
		addrPort := fmt.Sprintf("127.0.0.1:%d", port)
		for {
			if _, ok := g_grpcConn[base.MASTER_GSID]; !ok {
				InitRPCClient(base.MASTER_GSID, addrPort)
			} else {
				break
			}
		}
	}
}

func InitRPCServer() {
	var opts []grpc.ServerOption
	g_grpcServer = grpc.NewServer(opts...)
	InitServeModule()
}

func InitRPCClient(gsId int, addrPort string) bool {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addrPort, opts...)
	if err != nil {
		log.Printf("[InitRPCClient] fail to dial: %v", err)
		return false
	}
	g_grpcConn[gsId] = conn
	InitClientModule(gsId)
	return true
}

func StartServe() {
	cName := fmt.Sprintf("game_%d_port", base.GetGsId())
	ret, port := etc.GetConfigInt(cName)
	if !ret {
		log.Fatalf("[InitRPCServer] %s config error: %d", cName, port)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("[InitRPCServer] failed to listen: %v", err)
	}
	g_grpcServer.Serve(lis)
}

func GetServer() *grpc.Server {
	return g_grpcServer
}

func GetClientConn(gsId int) *grpc.ClientConn {
	v, ok := g_grpcConn[gsId]
	if ok {
		return v
	} else {
		return nil
	}
}

func InitServeModule() {
	pb.RegisterRpcEngineServer(GetServer(), proc.NewRpcServe())
}

func InitClientModule(gsId int) {
	cc, ok := g_grpcConn[gsId]
	if !ok {
		log.Fatalf("[InitClientModule] no %d connection", gsId)
	}

	pb.NewRpcEngineClient(cc)
}

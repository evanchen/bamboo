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
	"time"
)

// rpc 服务方,每个进程只有一个
var g_grpcServer *grpc.Server

// rpc 请求方
// 多个game进程连接0进程; 同时0进程连接到多个game进程
// 每个连接都需要初始化各个rpc服务模块的client端
type connInfo struct {
	conn    *grpc.ClientConn
	clients map[string]interface{}
}

var g_grpcConn = make(map[int]*connInfo)

// 0进程启动时,需等待game进程的rpc连接完成,然后主动连通各game进程的rpc服务
// 每个进程既作为rpc的client,也作为server,但game进程只与0进程相互连接
func StartRPC() {
	//初始化进程服务
	InitRPCServer()
	//初始化所有rpc模块服务
	InitServeModule()
	//rpc服务启动
	go StartServe()
	//客户端连接服务
	StartClientConn()

	//检查所有rpc进程端口连接情况
	go func() {
		for {
			if ret := proc.CheckRegFinish(); !ret {
				<-time.After(1 * time.Second)
			} else {
				break
			}
		}
		base.ServerReady()
	}()
}

func InitRPCServer() {
	var opts []grpc.ServerOption
	g_grpcServer = grpc.NewServer(opts...)
}

func StartClientConn() {
	for gsId, pInOut := range proc.GetServe().RegGsIds {
		port := int(pInOut.Port())
		go func(id, p int) {
			addrPort := fmt.Sprintf("127.0.0.1:%d", p)
			for {
				fmt.Printf("[StartClientConn] connecting: gsId: %d, addrPort: %s\n", id, addrPort)
				if ret := InitRPCClient(id, addrPort); !ret {
					<-time.After(1 * time.Second)
				} else {
					break
				}
			}
		}(gsId, port)
	}
}

func InitRPCClient(gsId int, addrPort string) bool {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(addrPort, opts...)
	fmt.Printf("[InitRPCClient] %v, %v\n", conn, err)
	if err != nil {
		log.Printf("[InitRPCClient] fail to dial: %v", err)
		return false
	}
	g_grpcConn[gsId] = &connInfo{
		conn:    conn,
		clients: make(map[string]interface{}),
	}
	InitClientModule(gsId)
	return true
}

func StartServe() {
	port := etc.GetRpcPort(base.GetGsId())
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("[StartServe] failed to listen: %v", err)
	}
	fmt.Println("[StartServe] start rpc serve listening...")
	g_grpcServer.Serve(lis)
}

func GetServer() *grpc.Server {
	return g_grpcServer
}

func GetClientConn(gsId int) *grpc.ClientConn {
	v, ok := g_grpcConn[gsId]
	if ok {
		return v.conn
	} else {
		return nil
	}
}

// 初始化所有rpc服务模块,创建服务
func InitServeModule() {
	pb.RegisterRpcEngineServer(GetServer(), proc.NewRpcEngineServe())
}

// 每初始化一个rpc模块就创建一个该模块的client,以模块名为key
func InitClientModule(gsId int) {
	cInfo, ok := g_grpcConn[gsId]
	if !ok {
		log.Fatalf("[InitClientModule] no %d connection", gsId)
	}

	// proc
	client := proc.GetServe().ClientRegist(gsId, cInfo.conn)
	cInfo.clients["proc"] = client
}

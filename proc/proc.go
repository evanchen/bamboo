package proc

import (
	"fmt"
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/etc"
	pb "github.com/evanchen/bamboo/pto/rpcpto"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"time"
)

type InOut struct {
	in     bool
	out    bool
	port   int64
	client interface{} // 重连不会改变
}

func (p *InOut) In() bool {
	return p.in
}

func (p *InOut) Out() bool {
	return p.out
}

func (p *InOut) Port() int64 {
	return p.port
}

// 进程需要的rpc连接
// [gsId] = InOut
type RpcEngine struct {
	RegGsIds map[int]*InOut
}

func GetServe() *RpcEngine {
	return g_serve
}

var g_serve *RpcEngine

func CheckRegFinish() bool {
	if g_serve != nil {
		for _, pInOut := range g_serve.RegGsIds {
			if !pInOut.In() || !pInOut.Out() {
				return false
			}
		}
		return true
	}
	return false
}

func NewRpcEngineServe() *RpcEngine {
	s := &RpcEngine{
		RegGsIds: make(map[int]*InOut),
	}
	// 0进程需要等待所有game进程连接完成
	if base.GetGsId() == base.MASTER_GSID {
		_, num := etc.GetConfigInt("max_game_num")
		fmt.Printf("[NewRpcEngineServe] max_game_num: %d\n", num)
		for gsId := 1; gsId <= int(num); gsId++ {
			port := etc.GetRpcPort(gsId)
			s.RegGsIds[gsId] = &InOut{
				in:   false,
				out:  false,
				port: port,
			}
		}
	} else {
		// 其他game进程只与0进程连接
		port := etc.GetRpcPort(base.MASTER_GSID)
		s.RegGsIds[base.MASTER_GSID] = &InOut{
			in:   false,
			out:  false,
			port: port,
		}
	}
	g_serve = s
	return g_serve
}

func (s *RpcEngine) RegisterEngine(ctx context.Context, p *pb.ReqRegister) (*pb.RetRegister, error) {
	gsId := int(p.GsId)
	pInOut := s.RegGsIds[gsId]
	if pInOut.Out() && pInOut.In() {
		//本进程已经请求过对方的服务了,但是对方仍请求注册服务时,说明对方进程可能重启了
		//此时延长1秒(先让本次函数调用返回),然后重新向对方新启动的进程注册服务,向对方确认,这边的in/out都是ok的
		//注意,grpc.ClientConn 是支持自动重连的,所以网络连接不需要再次手动connect
		go func() {
			<-time.After(1 * time.Second)
			pInOut.client.(pb.RpcEngineClient).RegisterEngine(context.Background(), &pb.ReqRegister{GsId: int32(base.GetGsId())})
		}()
	}
	pInOut.in = true
	fmt.Printf("[RegisterEngine] finish: %d\n", gsId)
	return &pb.RetRegister{Ret: 1}, nil
}

func (s *RpcEngine) ClientRegist(gsId int, conn *grpc.ClientConn) interface{} {
	client := pb.NewRpcEngineClient(conn)
	_, err := client.RegisterEngine(context.Background(), &pb.ReqRegister{GsId: int32(base.GetGsId())})
	if err != nil {
		log.Fatalf("[proc.InitClient] RegisterEngine failed: %s\n", err.Error())
	}
	s.RegGsIds[gsId].out = true
	s.RegGsIds[gsId].client = client

	fmt.Printf("[ClientRegist] successed. %d -> %d\n", base.GetGsId(), gsId)
	return client
}

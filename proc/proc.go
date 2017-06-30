package proc

import (
	pb "github.com/evanchen/bamboo/rpcpto"
	"golang.org/x/net/context"
)

type rpcServe struct {
}

func NewRpcServe() *rpcServe {
	return &rpcServe{}
}

func (s *rpcServe) RegisterGame(context.Context, *pb.ReqRegister) (*pb.RetRegister, error) {
	return nil, nil
}

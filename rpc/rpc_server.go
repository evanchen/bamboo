package rpc

type routeServer struct {

}

func newServer() *routeServer {
	s := new(routeServer)
	return s
}

func (s *routeServer) ProcStart(context.Context, params *pb.ReqParams) (*pb.RetParams, error) {

}

func (s *routeServer) GetInfo(context.Context, params *pb.ReqParams) (*pb.RetParams, error) {

}


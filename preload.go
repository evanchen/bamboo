package main

import (
	"flag"
	"fmt"
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/db"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/glog"
	"github.com/evanchen/bamboo/gnet"
	_ "github.com/evanchen/bamboo/pto"
	_ "github.com/evanchen/bamboo/pto/ptohandler"
	"github.com/evanchen/bamboo/rpc"
	"runtime"
)

var (
	gsid = flag.Int("gsid", -1, "game server id")
)

func main() {
	flag.Parse()
	base.InitGsId(*gsid)
	etc.LoadConfig()
	etc.CheckSysConfig(base.GetGsId())
	glog.Init()
	runtime.GOMAXPROCS(4)
	go rpc.StartRPC()
	// 等待所有game进程连接完毕,开放玩家连接
	fmt.Println("waiting all rpc connections...")
	base.IsServerReady()
	fmt.Println("rpc connections are ok.")
	glog.ChangeSysLogType(glog.LOG_RPC)
	if base.GetGsId() == base.MASTER_GSID {
		gnet.Start()
	}
	db.Init()

	fmt.Println("***************start to serve***************")
	<-base.GetShutdownCh()
	fmt.Println("server shutdown.")
}

package main

import (
	"flag"
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/glog"
	"github.com/evanchen/bamboo/rpc"
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
	rpc.StartRPC()

	// rpc服务启动,放在所有需要注册rpc服务的模块后
	rpc.StartServe()

	// 等待所有game进程连接完毕,开放玩家连接
	base.IsServerReady()
}

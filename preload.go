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
	base.InitGsID(*gsid)
	etc.LoadConfig()
	glog.Init()
	rpc.StartRPC()
	base.IsServerReady()
}


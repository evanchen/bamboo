package main

import (
	"flag"
	"fmt"
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/glog"
	"github.com/evanchen/bamboo/rpc"
	"runtime"
	"time"
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
	fmt.Println("waiting rpc connection...")
	base.IsServerReady()
	fmt.Println("all rpc connections are ready")
	close := make(chan bool)
	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				fmt.Println("tick tack")
			}
		}
		close <- true
	}()

	<-close
}

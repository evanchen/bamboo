package main

import (
	"base"
	"etc"
	"flag"
	"glog"
	"log"
)

var (
	gsid = flag.Int("gsid", -1, "game server id")
)

func main() {
	flag.Parse()
	base.InitGsID(*gsid)
	etc.LoadConfig()
	ret, lv := etc.GetConfigInt("log_level")
	if !ret {
		log.Fatal("config log_level error")
	}
	glog.ChangeSysLogLevel(int(lv))
	glog.CreateLocalLog()
}

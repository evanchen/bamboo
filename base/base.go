package base

import (
	"log"
)

// 本进程id
var g_gsid = -1

var server_ready = make(chan struct{})
var shutdownch = make(chan bool)

// 0进程id
const MASTER_GSID = 0

// 初始化本进程id
func InitGsId(gsid int) {
	curId := GetGsId()
	if curId != -1 {
		log.Fatal("[InitGsId]GSID is initiated")
	}
	if gsid < 0 {
		log.Fatal("[InitGsId]param error")
	}
	g_gsid = gsid
}

func GetGsId() int {
	return g_gsid
}

func IsServerReady() {
	<-server_ready
}

func ServerReady() {
	close(server_ready)
}

func GetShutdownCh() chan bool {
	return shutdownch
}

func Shutdown() {
	shutdownch <- true
}

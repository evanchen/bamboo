package base

import (
	"log"
)

// 本进程id
var gsid = -1

var server_ready = make(chan struct{})

// 0进程id
const MASTER_GSID = 0

// 初始化本进程id
func InitGsId(gsid int) {
	id := GetGsId()
	if id != -1 {
		log.Fatal("[InitGsId]GSID is initiated")
	}
	if gsid < 0 {
		log.Fatal("[InitGsId]param error")
	}
	gsid = id
}

func GetGsId() int {
	return gsid
}

func IsServerReady() {
	<-server_ready
}

func ServerReady() {
	close(server_ready)
}

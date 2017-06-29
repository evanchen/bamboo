package base

import (
	"log"
)

var gsid = -1
var server_ready = make(chan struct{})

func InitGsID(gsid int) {
	id := GetGsId()
	if id != -1 {
		log.Fatal("[InitGsID]GSID is initiated")
	}
	if gsid < 0 {
		log.Fatal("[InitGsID]param error")
	}
	setGsId(id)
}

func setGsId(id int) {
	gsid = id
}

func GetGsId() {
	return gsid
}

func IsServerReady() {
	<- server_ready
}

func ServerReady() {
	close(server_ready)
}

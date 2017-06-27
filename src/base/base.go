package base

import (
	"log"
)

var GSID = -1

func InitGsID(gsid int) {
	if GSID != -1 {
		log.Fatal("[InitGsID]GSID is initiated")
	}
	if gsid < 0 {
		log.Fatal("[InitGsID]param error")
	}
	GSID = gsid
}

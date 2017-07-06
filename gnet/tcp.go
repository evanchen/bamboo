package gnet

import (
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/glog"
	"github.com/evanchen/bamboo/pto"
	"github.com/evanchen/bamboo/pto/ptohandler"
	"log"
	"net"
	"syscall"
)

var logger = glog.New("log/login.log")

func Start() {
	ret, port := etc.GetConfigInt("master_port")
	if !ret {
		log.Fatalf("server listen port error: %d", port)
	}
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("false listening port: %s", err.Error())
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				logger.Error("accept error: %s", err)
				return
			}
			logger.Info("accept conn: %v", conn)
			go HandleConn(conn)
		}

		base.Shutdown()
	}()
}

func HandleConn(conn *net.Conn) {
	defer conn.Close()
	for {
		ptoId,data,err := ptohandler.Recv(conn)
		if err != nil {
			if err == io.EOF {
				logger.Info("conn closed: %v", conn)
			} else {
				logger.Error("conn: %v \nerror: %s", conn, err.Error())
			}
			break
		}
		err = ptohandler.HandleMsg(conn, ptoId, data)
		if err != nil {
			logger.Error("handler proto: %d, error: %s", ptoId, err.Error())
		}
	}
}

package gnet

import (
	"fmt"
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/glog"
	"github.com/evanchen/bamboo/pto"
	"github.com/evanchen/bamboo/pto/ptohandler"
	"io"
	"log"
	"net"
)

var logger = glog.New("log/tcp.log")

func Start() {
	ret, port := etc.GetConfigInt("master_port")
	if !ret {
		log.Fatalf("server listen port error: %d", port)
	}
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("false listening port: %s", err.Error())
	}

	fmt.Println("gnet.Start()...")
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				logger.Error("accept error: %s", err)
				return
			}
			go HandleConn(conn)
		}

		base.Shutdown()
	}()
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	logger.Info("[HandleConn] accept new connection: %v", conn)

	for {
		ptoId, data, err := pto.Recv(conn)
		if err != nil {
			if err == io.EOF {
				logger.Info("conn normally closed: %v", conn)
			} else {
				logger.Error("conn: %v error: \n%s", conn, err.Error())
			}
			break
		}
		err = ptohandler.HandleMsg(conn, ptoId, data)
		if err != nil {
			logger.Error("handler proto: %d, error: %s", ptoId, err.Error())
		}
	}
	logger.Info("[HandleConn] connection ending: %v", conn)
}

package gnet

import (
	"fmt"
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/glog"
	_ "github.com/evanchen/bamboo/pto"
	"github.com/evanchen/bamboo/pto/ptohandler"
	"io"
	"log"
	"net"
	"errors"
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

	fmt.Println("gnet.Start()...")
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

func HandleConn(conn net.Conn) {
	defer conn.Close()
	logger.Info("[HandleConn] accept new connection: %v", conn)

	for {
		ptoId, data, err := ptohandler.Recv(conn)
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
	logger.Info("[HandleConn] connection ending: %v", conn)
}

func Send(conn net.Conn, ptoId uint16, data []byte) error {
	ptoLen := ptohandler.TCP_HEADER_LEN + len(data)
	if !(ptoLen >= 0 && ptoLen < ptohandler.MAX_TCP_DATA_LEN) {
		return errors.New(fmt.Sprintf("len error: ptoLen: %d, ptoId: %d", ptoLen, ptoId))
	}
	tData := make([]byte, ptoLen)
	ptohandler.EncodeHeader(uint16(ptoLen), ptoId, tData)
	copy(tData[ptohandler.TCP_HEADER_LEN:], data)
	slen, err := conn.Write(tData)
	if err != nil {
		return err
	}
	return nil 
}
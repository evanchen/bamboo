package main

import (
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/pto"
	"github.com/evanchen/bamboo/pto/ptohandler"
	"log"
	"net"
	"fmt"
)

var shutdownch = make(chan bool)

func main() {
	etc.LoadConfig()
	ret, port := etc.GetConfigInt("master_port")
	if !ret {
		log.Fatalf("server port error: %d", port)
	}
	servAddr := fmt.Sprintf("127.0.0.1:%s", port)
	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		log.Fatalf("failed to connect to err: %s", err.Error())
	}
	go HandleConn(conn)
	<-shutdownch
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	for {
		ptoId, data, err := ptohandler.Recv(conn)
		if err != nil {
			if err == io.EOF {
				log.Fatalf("conn closed: %v", conn)
			} else {
				log.Fatalf("conn: %v \nerror: %s", conn, err.Error())
			}
			break
		}
		err = HandleMsg(conn, ptoId, data)
		if err != nil {
			log.Fatalf("handler proto: %d, error: %s", ptoId, err.Error())
		}
	}
	shutdownch <- true
}

func HandleMsg(conn net.Conn, ptoId uint16, data []byte) error {
	ptoName, ok := pto.GetPtoName(ptoId)
	if !ok {
		return errors.New(fmt.Sprintf("protocol is not exist: %d", ptoId))
	}
	ptoObj := pto.GetNewPto(ptoName)
	err := proto.Unmarshal(data, ptoObj)
	if err != nil {
		return err
	}
	if ptoName == "CLogin" {

	} else if ptoName == "CLoginRet" {

	}

	return nil
}

package main

import (
	"errors"
	"fmt"
	"github.com/evanchen/bamboo/etc"
	"github.com/evanchen/bamboo/pto"
	"github.com/evanchen/bamboo/pto/ptohandler"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"net"
)

var shutdownch = make(chan bool)

func main() {
	etc.LoadConfig()
	ret, port := etc.GetConfigInt("master_port")
	if !ret {
		log.Fatalf("server port error: %d", port)
	}
	servAddr := fmt.Sprintf("127.0.0.1:%d", port)
	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		log.Fatalf("failed to connect to err: %s", err.Error())
	}
	fmt.Printf("connected to server\n")
	go HandleConn(conn)
	<-shutdownch
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	firstPtoId, ok := pto.GetPtoId("SLogin")
	if !ok {
		log.Fatalf("[pto.GetPtoId] error: %s", "SLogin")
	}
	firstPto := pto.GetNewPto("SLogin").(*pto.SLogin)
	firstPto.Ver = "636d366d40d9d37abeb46431bcf5e382"
	firstPto.Account = "golang"
	firstPto.Passwd = "123456"
	firstData, err := proto.Marshal(firstPto)
	if err != nil {
		log.Fatalf("marshaling error: %s", err.Error())
	}
	SendPto(conn, firstPtoId, firstData)

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
	switch ptoObj.(type) {
	case *pto.CLogin:
		p := ptoObj.(*pto.CLogin)
		uid := p.Uid
		fmt.Printf("%v\n", p)

		sendPto := pto.GetNewPto("SLoginReq").(*pto.SLoginReq)
		sendPto.Uid = uid
		sData, err := proto.Marshal(sendPto)
		if err != nil {
			log.Fatalf("marshaling error: %s", err.Error())
		}
		SendPto(conn, ptoId, sData)
	case *pto.CLoginRet:
		p := ptoObj.(*pto.CLoginRet)
		fmt.Printf("%v\n", p)

	}

	return nil
}

func SendPto(conn net.Conn, ptoId uint16, data []byte) {
	ptoLen := ptohandler.TCP_HEADER_LEN + len(data)
	if !(ptoLen >= 0 && ptoLen < ptohandler.MAX_TCP_DATA_LEN) {
		log.Fatalf("len error: ptoLen: %d, ptoId: %d", ptoLen, ptoId)
	}
	header := make([]byte, ptohandler.TCP_HEADER_LEN)
	ptohandler.EncodeHeader(uint16(ptoLen), ptoId, header)
	header = append(header, data...)
	_, err := conn.Write(header)
	if err != nil {
		log.Fatalf("[SendPto] ptoId: %d, error: %s", ptoId, err.Error)
	}
	fmt.Printf("[SendPto] ptoId: %d ok. \n", ptoId)
}

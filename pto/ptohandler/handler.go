package ptohandler

import (
	"errors"
	"fmt"
	"github.com/evanchen/bamboo/pto"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
)

const (
	TCP_HEADER_LEN   = 4
	MAX_TCP_DATA_LEN = 65535
)

func Recv(conn net.Conn) (uint16, []byte, error) {
	header := make([]byte, TCP_HEADER_LEN)
	_, err := io.ReadFull(conn, header)
	if err != nil {
		return 0, nil, err
	}
	ptoLen, ptoId := ParseHeader(header)
	if !(ptoLen >= 0 && ptoLen < MAX_TCP_DATA_LEN) {
		return 0, nil, errors.New(fmt.Sprintf("len error: ptoLen: %d, ptoId: %d", ptoLen, ptoId))
	}
	data := make([]byte, ptoLen)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		return 0, nil, err
	}
	return ptoId, data, nil
}

func ParseHeader(header []byte) (uint16, uint16) {
	// 大端
	ptoLen := uint16(header[0]) << 8
	ptoLen |= uint16(header[1])

	ptoId := uint16(header[2]) << 8
	ptoId |= uint16(header[3])

	return ptoLen, ptoId
}

func HandleMsg(conn net.Conn, ptoId uint16, data []byte) error {
	ptoName, ok := pto.GetPtoName(ptoId)
	if !ok {
		return errors.New(fmt.Sprintf("protocol is not exist: %d", ptoId))
	}
	ptoObj := pto.GetNewPto(ptoName)
	fn, ok := handlerFunc[ptoName]
	if !ok {
		return errors.New(fmt.Sprintf("protocol handler is not exist: %d, %s", ptoId, ptoName))
	}
	err := proto.Unmarshal(data, ptoObj)
	if err != nil {
		return err
	}
	return fn(conn, ptoObj)
}

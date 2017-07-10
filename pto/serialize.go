// pto 模块作为协议模块(.pb.go集合)
// 本模块出了依赖系统包和proto包外,不依赖其他包
package pto

import (
	"errors"
	"fmt"
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
	ptoLen, ptoId := DecodeHeader(header)
	if !(ptoLen >= 0 && ptoLen < MAX_TCP_DATA_LEN) {
		return 0, nil, errors.New(fmt.Sprintf("len error: ptoLen: %d, ptoId: %d\n", ptoLen, ptoId))
	}
	//fmt.Printf("[ptohandler.Recv]: header: %v, ptoLen: %d, ptoId: %d", header, ptoLen, ptoId)
	data := make([]byte, ptoLen-TCP_HEADER_LEN)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		return 0, nil, err
	}
	return ptoId, data, nil
}

func Send(conn net.Conn, ptoId uint16, data []byte) error {
	ptoLen := TCP_HEADER_LEN + len(data)
	if !(ptoLen >= 0 && ptoLen < MAX_TCP_DATA_LEN) {
		return errors.New(fmt.Sprintf("len error: ptoLen: %d, ptoId: %d", ptoLen, ptoId))
	}
	tData := make([]byte, ptoLen)
	EncodeHeader(uint16(ptoLen), ptoId, tData)
	copy(tData[TCP_HEADER_LEN:], data)
	_, err := conn.Write(tData)
	if err != nil {
		return err
	}
	return nil
}

func DecodeHeader(header []byte) (uint16, uint16) {
	// 大端
	ptoLen := uint16((header[0] << 8))
	ptoLen |= uint16(header[1])
	ptoId := uint16((header[2]) << 8)
	ptoId |= uint16(header[3])
	return ptoLen, ptoId
}

func EncodeHeader(ptoLen, ptoId uint16, header []byte) {
	// 大端
	header[0] = byte((ptoLen >> 8) & 0xff)
	header[1] = byte(ptoLen & 0xff)
	header[2] = byte((ptoId >> 8) & 0xff)
	header[3] = byte(ptoId & 0xff)
}

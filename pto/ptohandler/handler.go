// 本模块负责处理客户端消息请求
// 为每个.proto独立创建一个handler文件,在完成消息解码后,调用各逻辑模块处理消息
package ptohandler

import (
	"errors"
	"fmt"
	"github.com/evanchen/bamboo/pto"
	"github.com/golang/protobuf/proto"
	"net"
)

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

package ptohandler

import (
	_ "errors"
	"fmt"
	"github.com/evanchen/bamboo/pto"
	"github.com/golang/protobuf/proto"
	"net"
)

func HandleSLogin(conn net.Conn, d interface{}) error {
	p, _ := d.(*pto.SLogin)
	fmt.Printf("%v\n", p)

	if p.Ver != pto.GetVersion() {
		ptoId, _ := pto.GetPtoId("CLoginVer")
		retPto := pto.GetNewPto("CLoginVer").(*pto.CLoginVer)
		retPto.Ret = false
		data, err := proto.Marshal(retPto)
		if err != nil {
			return err
		}
		return pto.Send(conn, ptoId, data)
	} else {
		ptoId, _ := pto.GetPtoId("CLogin")
		retPto := pto.GetNewPto("CLogin").(*pto.CLogin)
		retPto.Uid = 54321
		data, err := proto.Marshal(retPto)
		if err != nil {
			return err
		}
		return pto.Send(conn, ptoId, data)
	}

	return nil
}

func HandleSLoginReq(conn net.Conn, d interface{}) error {
	p, _ := d.(*pto.SLoginReq)
	fmt.Printf("%v\n", p)
	return nil
}

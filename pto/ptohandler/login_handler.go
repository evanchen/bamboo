package ptohandler

import (
	"fmt"
	"github.com/evanchen/bamboo/pto"
	"net"
	"github.com/evanchen/bamboo/gnet"
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
		return gnet.Send(conn,ptoId,data)
	}
	return nil
}

func HandleSLoginReq(conn net.Conn, d interface{}) error {
	p, _ := d.(*pto.SLoginReq)
	fmt.Printf("%v\n", p)
	return nil
}

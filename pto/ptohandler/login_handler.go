package ptohandler

import (
	"fmt"
	"github.com/evanchen/bamboo/pto"
	"net"
)

func HandleSLogin(conn net.Conn, d interface{}) error {
	p, _ := d.(*pto.SLogin)
	fmt.Printf("%v\n", p)
	return nil
}

func HandleSLoginReq(conn net.Conn, d interface{}) error {
	p, _ := d.(*pto.SLoginReq)
	fmt.Printf("%v\n", p)
	return nil
}
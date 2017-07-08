// This file is created by ptoVersion. DO NOT EDIT.
package ptohandler

import(
	"net"
)

var handlerFunc = make(map[string](func(net.Conn, interface{}) error))

func init() {
	handlerFunc["SLogin"] = HandleSLogin
	handlerFunc["SLoginReq"] = HandleSLoginReq
}


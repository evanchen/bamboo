// This file is created by ptoVersion. DO NOT EDIT.
package pto

import (
	"github.com/golang/protobuf/proto"
)

var md5sum string
var id2Name = make(map[uint16]string)
var name2Id = make(map[string]uint16)
var name2Func = make(map[string](func() proto.Message))

func init() {
	id2Name[1] = "SLogin"
	name2Id["SLogin"] = 1
	name2Func["SLogin"] = func() proto.Message { return &SLogin{} }

	id2Name[2] = "CLogin"
	name2Id["CLogin"] = 2
	name2Func["CLogin"] = func() proto.Message { return &CLogin{} }

	id2Name[3] = "CLoginVer"
	name2Id["CLoginVer"] = 3
	name2Func["CLoginVer"] = func() proto.Message { return &CLoginVer{} }

	id2Name[4] = "SLoginReq"
	name2Id["SLoginReq"] = 4
	name2Func["SLoginReq"] = func() proto.Message { return &SLoginReq{} }

	id2Name[5] = "CLoginRet"
	name2Id["CLoginRet"] = 5
	name2Func["CLoginRet"] = func() proto.Message { return &CLoginRet{} }

	md5sum = "09b0da6dcab93421f105a77f7388656d"
}

func GetPtoName(id uint16) (string, bool) {
	var v, ok = id2Name[id]
	return v, ok
}

func GetPtoId(name string) (uint16, bool) {
	var v, ok = name2Id[name]
	return v, ok
}

func GetNewPto(name string) proto.Message {
	var fn, ok = name2Func[name]
	if !ok {
		return nil
	}
	return fn()
}

func GetVersion() string {
	return md5sum
}

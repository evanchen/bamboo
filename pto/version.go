// This file is created by ptoVersion. DO NOT EDIT.
package pto

var md5sum string
var id2Name = make(map[uint16]string)
var name2Id = make(map[string]uint16)
var name2Func = make(map[string](func() interface{}))

func init() {
	id2Name[1] = "SLogin"
	name2Id["SLogin"] = 1
	name2Func["SLogin"] = func() interface{} { return &SLogin{} }

	id2Name[2] = "CLogin"
	name2Id["CLogin"] = 2
	name2Func["CLogin"] = func() interface{} { return &CLogin{} }

	id2Name[3] = "SLoginReq"
	name2Id["SLoginReq"] = 3
	name2Func["SLoginReq"] = func() interface{} { return &SLoginReq{} }

	id2Name[4] = "CLoginRet"
	name2Id["CLoginRet"] = 4
	name2Func["CLoginRet"] = func() interface{} { return &CLoginRet{} }

	md5sum = "636d366d40d9d37abeb46431bcf5e382"
}

func GetPtoName(id uint16) (string, bool) {
	var v, ok = id2Name[id]
	return v, ok
}

func GetPtoId(name string) (uint16, bool) {
	var v, ok = name2Id[name]
	return v, ok
}

func GetNewPto(name string) interface{} {
	var fn, ok = name2Func[name]
	if !ok {
		return nil
	}
	return fn()
}

func GetVersion() string {
	return md5sum
}

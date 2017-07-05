// This file is created by ptoVersion. DO NOT EDIT.
package pto

var md5sum string
var id2Name = make(map[uint16]string)
var name2Id = make(map[string]uint16)
var name2Func = make(map[string](func() interface{}))

func Init() {
	id2Name[1] = "Person"
	name2Id["Person"] = 1
	name2Func["Person"] = func() interface{} { return &Person{} }

	id2Name[2] = "Person_PhoneNumber"
	name2Id["Person_PhoneNumber"] = 2
	name2Func["Person_PhoneNumber"] = func() interface{} { return &Person_PhoneNumber{} }

	id2Name[3] = "AddressBook"
	name2Id["AddressBook"] = 3
	name2Func["AddressBook"] = func() interface{} { return &AddressBook{} }

	md5sum = "989628efd57e125c3cbf3878eb426351"
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

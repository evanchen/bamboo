package main

import (
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	addrs := &AddressBook{
		People: []*Person{&Person{
			Name:  "haha",
			Id:    123,
			Email: "haha@gmail.com",
			Phones: []*Person_PhoneNumber{
				&Person_PhoneNumber{
					Number: "22112233",
					Type:   Person_MOBILE,
				},
			},
		},
		}}
	data, err := proto.Marshal(addrs)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newAddrs := &AddressBook{}
	err = proto.Unmarshal(data, newAddrs)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	if newAddrs.People[0].Name == addrs.People[0].Name {
		log.Printf("ok!\n")
	}
}

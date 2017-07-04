package ptoVersion

import (
	"fmt"
	"gloger"
	"msghandler"
	"net"
	"os"
	"path/filepath"
	"protocol"
	"runtime"
	"strings"
)

var filenamech = make(chan string,5)
var structnamech = make(chan string,5)
//walk through child directories and files
func DoAnalysis() {
	path := "." //当前目录
	go func() {
		defer close(filenamech)
		filepath.Walk(path, walkFunc)
	}()
	
	go func () {
		for pbFile := range filenamech {
			go doFile(pbFile)
		}
	}()

	go func () {
		verFileName_tmp := path+"/version.tmp"
		verFileName := path+"/version.go"
		wf, err := os.OpenFile(verFileName_tmp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatalf("create version file: %s error: %v",verFileName_tmp,err)
		}
		defer wf.Close()
		fmt.Fprintln(wf, "// This file is automatically created. DO NOT EDIT.")
		fmt.Fprintln(wf, "package pto")
		fmt.Fprintln(wf,"\n\n")
		fmt.Fprintln(wf, "var Name2Id = make(map[string]uint16)")
		fmt.Fprintln(wf, "var Id2Name = make(map[uint16]string)")
		fmt.Fprintln(wf, "var Name2NewFunc = make(map[string]func(){})")
		fmt.Fprintln(wf,"\n\n")
		fmt.Fprintln(wf, "func Init() {")
		lineNum := 0
		for sname := range structnamech {
			lineNum ++
			fmt.Fprintf(wf, "\tName2Id[%s] = %d\n",lineNum,sname)
			fmt.Fprintf(wf, "\tId2Name[%d] = %s\n",sname,lineNum)
			fmt.Fprintf(wf, "\tName2NewFunc[%s] = func(){ return &%s{} }\n",sname,sname)
		}
		fmt.Fprintln(wf, "}")
		fmt.Fprintln(wf,"\n")
		fmt.Fprintln(wf, "func GetPtoName(id uint16) (string,bool) { return Name2Id[id] }")
		fmt.Fprintln(wf, "func GetPtoId(sname string) (uint16,bool) { return Id2Name[sname] }")
		fmt.Fprintln(wf, "func GetNewPto(sname string) (fuc(){},bool) { return Name2NewFunc[sname] }")
	}()
}

//walk function
func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		panic(err)
	}
	isdir := info.IsDir()
	name := info.Name()

	if isdir {
		if name == "." || name == ".." {
			return nil
		}
	} else {
		if !strings.HasSuffix(name,".proto") {
			return nil
		}
	}

	filenamech <- name

	return nil
}

func doFile(fname string) {
	rf,err := os.Open(fname)
	if err != nil {
		log.Fatalf("create version file: %s error: %v",fname,err)
	}
	defer rf.Close()


}

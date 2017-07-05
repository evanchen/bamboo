package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var filenamech = make(chan string)
var structnamech = make(chan string, 5)
var wg1 sync.WaitGroup
var wg2 sync.WaitGroup
var uniqueStructName = make(map[string]bool)

//walk through child directories and files
func Start() {
	path := ".." //上一层目录
	go func() {
		defer close(filenamech)
		filepath.Walk(path, walkFunc)
	}()

	go func() {
		for pbFile := range filenamech {
			wg1.Add(1)
			go doFile("../" + pbFile)
		}
		wg1.Wait()
		close(structnamech)
	}()

	go func() {
		defer wg2.Done()
		verFileName_tmp := path + "/version.tmp"
		verFileName := path + "/version.go"
		wf, err := os.OpenFile(verFileName_tmp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatalf("create version file: %s error: %v", verFileName_tmp, err)
		}
		fmt.Fprintln(wf, "// This file is created by ptoVersion. DO NOT EDIT.")
		fmt.Fprintln(wf, "package pto\n")
		fmt.Fprintln(wf, "var md5sum string")
		fmt.Fprintln(wf, "var id2Name = make(map[uint16]string)")
		fmt.Fprintln(wf, "var name2Id = make(map[string]uint16)")
		fmt.Fprintln(wf, "var name2Func = make(map[string](func() interface{}))\n")
		fmt.Fprintln(wf, "func Init() {")
		var content []byte
		lineNum := 0
		for sname := range structnamech {
			lineNum++
			fmt.Fprintf(wf, "\tid2Name[%d] = \"%s\"\n", lineNum, sname)
			fmt.Fprintf(wf, "\tname2Id[\"%s\"] = %d\n", sname, lineNum)
			fmt.Fprintf(wf, "\tname2Func[\"%s\"] = func() interface{} { return &%s{} }\n\n", sname, sname)

			txt := ([]byte)(fmt.Sprintf("%d=%s", lineNum, sname))
			content = append(content, txt...)
		}
		fmt.Fprintf(wf, "\tmd5sum = \"%x\"\n", md5.Sum(content))
		fmt.Fprintln(wf, "}\n")
		fmt.Fprintln(wf, "func GetPtoName(id uint16) (string, bool) {")
		fmt.Fprintln(wf, "\tvar v, ok = id2Name[id]")
		fmt.Fprintln(wf, "\treturn v, ok")
		fmt.Fprintln(wf, "}\n")

		fmt.Fprintln(wf, "func GetPtoId(name string) (uint16, bool) {")
		fmt.Fprintln(wf, "\tvar v, ok = name2Id[name]")
		fmt.Fprintln(wf, "\treturn v, ok")
		fmt.Fprintln(wf, "}\n")

		fmt.Fprintln(wf, "func GetNewPto(name string) interface{} {")
		fmt.Fprintln(wf, "\tvar fn, ok = name2Func[name]")
		fmt.Fprintln(wf, "\tif !ok {")
		fmt.Fprintln(wf, "\t\treturn nil")
		fmt.Fprintln(wf, "\t}")
		fmt.Fprintln(wf, "\treturn fn()")
		fmt.Fprintln(wf, "}\n")

		fmt.Fprintln(wf, "func GetVersion() string {")
		fmt.Fprintln(wf, "\treturn md5sum")
		fmt.Fprintln(wf, "}\n")

		wf.Sync()
		wf.Close()
		os.Rename(verFileName_tmp, verFileName)
		fmt.Println("finish.")
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
		return nil
	}
	if !strings.HasSuffix(name, ".pb.go") {
		return nil
	}

	fmt.Printf("path: %s, name: %s\n", path, name)
	filenamech <- name

	return nil
}

func doFile(fname string) {
	defer wg1.Done()
	rf, err := os.Open(fname)
	if err != nil {
		log.Fatalf("open file: %s error: %v", fname, err)
	}
	defer rf.Close()
	info, _ := rf.Stat()
	txt := make([]byte, info.Size())
	size, err := rf.Read(txt)
	if size != int(info.Size()) {
		log.Fatalf("doFile: size not match: %s, Size().%d == Read().%d\n", fname, info.Size(), size)
	}
	reg1 := regexp.MustCompile(`type ([\w]+) struct \{`)
	structs := reg1.FindAllStringSubmatch(string(txt), -1)
	for _, v := range structs {
		fmt.Printf("%v\n", v[1])
		_, ok := uniqueStructName[v[1]]
		if ok {
			log.Fatalf("uniqueStructName: file: %s, struct: %s", fname, v[1])
		}
		uniqueStructName[v[1]] = true
		structnamech <- v[1]
	}
}

func main() {
	wg2.Add(1)
	Start()
	wg2.Wait()
}
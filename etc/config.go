package etc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Data map[string]string
}

var mConfig = &Config{
	Data: make(map[string]string),
}

func LoadConfig() *Config {
	file, err := os.Open("./etc/config.sys")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	newCfg := &Config{
		Data: make(map[string]string),
	}

	reader := bufio.NewReader(file)
	lineNum := 0
	for {
		line, _, rerr := reader.ReadLine()
		if rerr == io.EOF {
			break
		} else if rerr != nil {
			log.Fatal(rerr)
		}
		if len(line) == 0 {
			continue
		}
		lineNum++
		newCfg.ParseLine(line)
	}

	mConfig = newCfg
	return mConfig
}

func (cfg *Config) ParseLine(line []byte) {
	for pos, char := range line {
		if char == '#' {
			if pos == 0 {
				return
			}
			line = line[:pos-1]
			break
		}
	}
	tarLine := strings.Split(string(line), "=")
	if len(tarLine) != 2 {
		log.Fatalf("\"%s\" should be 'key = value'", line)
	}
	cName, cValue := strings.TrimSpace(tarLine[0]), strings.TrimSpace(tarLine[1])
	if len(cName) == 0 || len(cValue) == 0 {
		log.Printf("\"%s\": key or value is empty", line)
		return
	}
	if _, ok := cfg.Data[cName]; ok {
		log.Printf("\"%s\": key repeated!", line)
	}
	cfg.Data[cName] = cValue
	fmt.Println(cName, cValue)
}

func GetConfigInt(key string) (bool, int64) {
	v, ok := mConfig.Data[key]
	if !ok {
		return false, -1
	}
	value, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		log.Printf("[GetConfigInt]: %s error: :%s", key, err.Error())
		return false, -1
	}
	return true, value
}

func GetConfigStr(key string) (bool, string) {
	v, ok := mConfig.Data[key]
	if !ok {
		return false, ""
	}
	return true, v
}

// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func GetConfigBool(key string) (bool, bool) {
	v, ok := mConfig.Data[key]
	if !ok {
		return false, false
	}
	value, err := strconv.ParseBool(v)
	if err != nil {
		log.Printf("[GetConfigBool]: %s error: :%s", key, err.Error())
		return false, false
	}
	return true, value
}

func GetConfigFloat(key string) (bool, float64) {
	v, ok := mConfig.Data[key]
	if !ok {
		return false, -1
	}
	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Printf("[GetConfigFloat]: %s error: :%s", key, err.Error())
		return false, -1
	}
	return true, value
}

func Test() {
	for k, v := range mConfig.Data {
		key1, value1 := GetConfigInt(k)
		key2, value2 := GetConfigFloat(k)
		key3, value3 := GetConfigBool(k)
		key4, value4 := GetConfigStr(k)

		fmt.Println("============", k, v)
		fmt.Println(key1, value1)
		fmt.Println(key2, value2)
		fmt.Println(key3, value3)
		fmt.Println(key4, value4)
	}
}

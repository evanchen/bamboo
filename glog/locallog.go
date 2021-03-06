package glog

import (
	"fmt"
	"github.com/evanchen/bamboo/base"
)

// 在rpc日志没启动前,进程里写入的所有日志都会写到本地日志
var g_logger *Logger

func CreateLocalLog() {
	path := fmt.Sprintf("log/engine/s%d_local.log", base.GetGsId())
	g_logger = CreateLog(path)
}

package glog

import (
	"base"
	"fmt"
)

// 系统日志等级
type type_log_level int

const (
	LOG_LEVEL_DEBUG type_log_level = iota // 0
	LOG_LEVEL_INFO                        // 1
	LOG_LEVEL_ERROR                       // 2
)

// 系统日志类型
type type_log int

const (
	LOG_LOCAL type_log = iota // 0,本地写
	LOG_RPC                   // 1,远程日志
)

type ModuleLogger struct {
	path string
}

var g_log_level = LOG_LEVEL_DEBUG
var g_log_type = LOG_LOCAL
var g_runtime_log_path = "log/engine/runtime.log"

func ChangeSysLogLevel(lv int) {
	g_log_level = type_log_level(lv)
}

func ChangeSysLogType(tp int) {
	g_log_type = type_log(tp)
}

func WriteLog(path, ctn string) {
	if g_log_type == LOG_RPC {
		//RPC.SendLog(path,ctn)
	} else {
		WriteFile(g_logger.path, ctn)
	}
}

func New(path string) *ModuleLogger {
	return &ModuleLogger{
		path: path,
	}
}

func (lg *Logger) WriteFunc(lv type_log_level, cls, format string, args ...interface{}) {
	if g_log_level < lv {
		return
	}

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		errStr := fmt.Sprintf("[%s][%02d] %s\n", cls, base.GSID, format)
	// 		fmt.Println(errStr)
	// 		WriteLog(g_runtime_log_path, errStr)
	// 	}
	// }()
	ctn := fmt.Sprintf(format, args...)
	ctn = fmt.Sprintf("[%s][%02d] %s\n", cls, base.GSID, ctn)
	WriteLog(lg.path, ctn)
}

func (lg *Logger) Debug(format string, args ...interface{}) {
	lg.WriteFunc(LOG_LEVEL_DEBUG, "Debug", format, args...)
}

func (lg *Logger) Info(format string, args ...interface{}) {
	lg.WriteFunc(LOG_LEVEL_INFO, "Info", format, args...)
}

func (lg *Logger) Error(format string, args ...interface{}) {
	lg.WriteFunc(LOG_LEVEL_ERROR, "Error", format, args...)
}

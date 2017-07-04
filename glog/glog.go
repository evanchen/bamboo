package glog

import (
	"fmt"
	"github.com/evanchen/bamboo/base"
	"github.com/evanchen/bamboo/etc"
	pb "github.com/evanchen/bamboo/rpcpto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
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

// 发送到日志服务的日志,通过channel发送
var chmsg = make(chan *pb.LogInfo, 50)

func ChangeSysLogLevel(lv type_log_level) {
	g_log_level = lv
}

func ChangeSysLogType(tp type_log) {
	g_log_type = tp
}

func WriteLog(path, ctn string) {
	if g_log_type == LOG_RPC {
		chmsg <- &pb.LogInfo{Path: path, Content: ctn}
	} else {
		WriteFile(g_logger.path, ctn)
	}
}

func New(path string) *ModuleLogger {
	return &ModuleLogger{
		path: path,
	}
}

func (lg *ModuleLogger) WriteFunc(lv type_log_level, cls, format string, args ...interface{}) {
	if g_log_level < lv {
		return
	}
	ctn := fmt.Sprintf(format, args...)
	ctn = fmt.Sprintf("[%s][%02d] %s", cls, base.GetGsId(), ctn)
	WriteLog(lg.path, ctn)
}

func (lg *ModuleLogger) Debug(format string, args ...interface{}) {
	lg.WriteFunc(LOG_LEVEL_DEBUG, "Debug", format, args...)
}

func (lg *ModuleLogger) Info(format string, args ...interface{}) {
	lg.WriteFunc(LOG_LEVEL_INFO, "Info", format, args...)
}

func (lg *ModuleLogger) Error(format string, args ...interface{}) {
	lg.WriteFunc(LOG_LEVEL_ERROR, "Error", format, args...)
}

func Init() {
	ret, lv := etc.GetConfigInt("log_level")
	if !ret {
		log.Fatal("config log_level error")
	}
	ChangeSysLogLevel(type_log_level(lv))
	CreateLocalLog()
	StartRpcLog()
}

func StartRpcLog() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	ret, port := etc.GetConfigInt("log_server_port")
	if !ret {
		log.Fatalf("[StartRpcLog] log_server_port error\n")
	}
	addrPort := fmt.Sprintf("127.0.0.1:%d", port)
	conn, err := grpc.Dial(addrPort, opts...)
	fmt.Printf("[StartRpcLog] %v, %v\n", conn, err)
	if err != nil {
		log.Fatalf("[StartRpcLog] fail to dial: %v", err)
	}
	client := pb.NewRpcLogClient(conn)

	go func() {
		stream, err := client.SendLog(context.Background())
		if err != nil {
			log.Fatalf("%v.SendLog(_) = _, %v", client, err)
		}
		for msg := range chmsg {
			if err := stream.Send(msg); err != nil {
				log.Fatalf("%v.Send(%v) = %v", stream, msg, err)
			}
		}

		if _, err := stream.CloseAndRecv(); err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
	}()
}

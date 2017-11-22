package xlog

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/ebitgo/ConsoleColor"
)

// XLogger 日志结构
type XLogger struct {
	FilePath     string
	flogger      *logs.BeeLogger
	OpenInfo     bool
	InfoColor    int
	OpenError    bool
	ErrorColor   int
	OpenTrace    bool
	TraceColor   int
	OpenWarning  bool
	WarningColor int
	OpenDebug    bool
	DebugColor   int
}

// Init 初始化
func (ths *XLogger) Init(file string) {
	dir := path.Dir(file)
	err := os.MkdirAll(dir, os.FileMode(0777))
	if err != nil {
		panic(fmt.Sprintf("Can not create log folder! error is \n%+v\n", err))
	}
	ths.FilePath = file
	ths.flogger = logs.NewLogger(10000)
	ths.flogger.EnableFuncCallDepth(true)
	ths.flogger.SetLogger("file", `{"filename":"`+ths.FilePath+".log"+`"}`)
	ths.OpenInfo = true
	ths.InfoColor = ConsoleColor.C_WHITE
	ths.OpenError = true
	ths.ErrorColor = ConsoleColor.C_RED
	ths.OpenTrace = true
	ths.TraceColor = ConsoleColor.C_CYAN
	ths.OpenWarning = true
	ths.WarningColor = ConsoleColor.C_BLUE
	ths.OpenDebug = true
	ths.DebugColor = ConsoleColor.C_YELLOW
}

// SetLogFunCallDepth 设置深度
func (ths *XLogger) SetLogFunCallDepth(d int) {
	ths.flogger.SetLogFuncCallDepth(d)
}

// Info 写入日志 等级是Infomation
func (ths *XLogger) Info(format string, v ...interface{}) {
	if ths.OpenInfo {
		if v != nil {
			ths.flogger.Info(format, v...)
		} else {
			ths.flogger.Info(format)
		}
	}
}

// InfoPrint 写入日志并打印到屏幕 等级是Infomation
func (ths *XLogger) InfoPrint(format string, v ...interface{}) {
	if ths.OpenInfo {
		ret := ""
		if v == nil {
			ret += fmt.Sprintf(format)
		} else {
			ret += fmt.Sprintf(format, v...)
		}
		ConsoleColor.Printf(ths.InfoColor, "%s [I] %s", time.Now().Format("2006-01-02 15:04:05.000"), ret)
		ths.Info(ret)
	}
}

// Error 写入日志 等级是Error
func (ths *XLogger) Error(format string, v ...interface{}) {
	if ths.OpenError {
		if v != nil {
			ths.flogger.Error(format, v...)
		} else {
			ths.flogger.Error(format)
		}
	}
}

// ErrorPrint 写入日志并打印到屏幕 等级是Error
func (ths *XLogger) ErrorPrint(format string, v ...interface{}) {
	if ths.OpenError {
		ret := ""
		if v == nil {
			ret += fmt.Sprintf(format)
		} else {
			ret += fmt.Sprintf(format, v...)
		}
		ConsoleColor.Printf(ths.ErrorColor, "%s [E] %s", time.Now().Format("2006-01-02 15:04:05.000"), ret)
		ths.Error(ret)
	}
}

// Trace 写入日志 等级是Trace
func (ths *XLogger) Trace(format string, v ...interface{}) {
	if ths.OpenTrace {
		if v != nil {
			ths.flogger.Trace(format, v...)
		} else {
			ths.flogger.Trace(format)
		}
	}
}

// TracePrint 写入日志并打印到屏幕 等级是Trace
func (ths *XLogger) TracePrint(format string, v ...interface{}) {
	if ths.OpenTrace {
		ret := ""
		if v == nil {
			ret += fmt.Sprintf(format)
		} else {
			ret += fmt.Sprintf(format, v...)
		}
		ConsoleColor.Printf(ths.TraceColor, "%s [T] %s", time.Now().Format("2006-01-02 15:04:05.000"), ret)
		ths.Trace(ret)
	}
}

// Warning 写入日志 等级是Warning
func (ths *XLogger) Warning(format string, v ...interface{}) {
	if ths.OpenWarning {
		if v != nil {
			ths.flogger.Warning(format, v...)
		} else {
			ths.flogger.Warning(format)
		}
	}
}

// WarningPrint 写入日志并打印到屏幕 等级是Warning
func (ths *XLogger) WarningPrint(format string, v ...interface{}) {
	if ths.OpenWarning {
		ret := ""
		if v == nil {
			ret += fmt.Sprintf(format)
		} else {
			ret += fmt.Sprintf(format, v...)
		}
		ConsoleColor.Printf(ths.WarningColor, "%s [W] %s", time.Now().Format("2006-01-02 15:04:05.000"), ret)
		ths.Warning(ret)
	}
}

// Debug 写入日志 等级是Debug
func (ths *XLogger) Debug(format string, v ...interface{}) {
	if ths.OpenDebug {
		if v != nil {
			ths.flogger.Debug(format, v...)
		} else {
			ths.flogger.Debug(format)
		}
	}
}

// DebugPrint 写入日志并打印到屏幕 等级是Debug
func (ths *XLogger) DebugPrint(format string, v ...interface{}) {
	if ths.OpenDebug {
		ret := ""
		if v == nil {
			ret += fmt.Sprintf(format)
		} else {
			ret += fmt.Sprintf(format, v...)
		}
		ConsoleColor.Printf(ths.DebugColor, "%s [D] %s", time.Now().Format("2006-01-02 15:04:05.000"), ret)
		ths.Debug(ret)
	}
}

package xlog

// LogInstance : log instance
var LogInstance *XLogger

// CreateDefaultLogInstance : create log instance
func CreateDefaultLogInstance(f string, d int) *XLogger {
	LogInstance = new(XLogger)
	LogInstance.Init(f)
	LogInstance.SetLogFunCallDepth(d)
	return LogInstance
}

// Info : print info message to stdout and file
func Info(format string, v ...interface{}) {
	if LogInstance != nil {
		LogInstance.InfoPrint("> "+format+"\n", v...)
	}
}

// Error : print error message to stdout and file
func Error(format string, v ...interface{}) {
	if LogInstance != nil {
		LogInstance.ErrorPrint("> "+format+"\n", v...)
	}
}

// Trace : print trace message to stdout and file
func Trace(format string, v ...interface{}) {
	if LogInstance != nil {
		LogInstance.TracePrint("> "+format+"\n", v...)
	}
}

// Warning : print warning message to stdout and file
func Warning(format string, v ...interface{}) {
	if LogInstance != nil {
		LogInstance.WarningPrint("> "+format+"\n", v...)
	}
}

// Debug : print debug message to stdout and file
func Debug(format string, v ...interface{}) {
	if LogInstance != nil {
		LogInstance.DebugPrint("> "+format+"\n", v...)
	}
}

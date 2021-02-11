package log

import (
	"crypto-balancer/src/core/environment"
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

const (
	debugTag      = "[Debug]"
	errorTag      = "[Error]"
	fatalErrorTag = "[Fatal]"
	warningTag    = "[Warning]"
)

func init() {
	logger = log.New(os.Stderr, "["+environment.AppName()+"]: ", log.LstdFlags|log.Lshortfile)
}

func logInfo(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func logFatal(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func LogInfo(format string, v ...interface{}) {
	logInfo(fmt.Sprintf("%s %s", debugTag, format), v)
}

func LogDebug(format string, v ...interface{}) {
	logInfo(fmt.Sprintf("%s %s", debugTag, format), v)
}

func LogWarning(format string, v ...interface{}) {
	logInfo(fmt.Sprintf("%s %s", warningTag, format), v)
}

func LogFatal(format string, v ...interface{}) {
	logFatal(fmt.Sprintf("%s %s", fatalErrorTag, format), v)
}

func LogError(format string, v ...interface{}) {
	logFatal(fmt.Sprintf("%s %s", errorTag, format), v)
}

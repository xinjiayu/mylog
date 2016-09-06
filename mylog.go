package mylog

import (
	"flag"

	"github.com/xinjiayu/log"
)

//定义日志处理参数
var (
	logLevel = flag.String("L", "info", "log level: info,debug,warn,error,fatal")
	logFile  = flag.String("log-file", "", "log file path")
)

func LogInit() {
	flag.Parse()
	if len(*logLevel) > 0 {
		log.SetLevelByString(*logLevel)
	}
	// set log options
	if len(*logFile) > 0 {
		err := log.SetOutputByName(*logFile)
		if err != nil {
			log.Fatal(err)
		}
		log.SetRotateByDay()
	}

}

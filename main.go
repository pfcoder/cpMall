package main

import (
    "github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

func initilize() {
	// Init log control
	log = logs.NewLogger(10000)
	log.SetLogger("console", "")
	log.EnableFuncCallDepth(true)
	//log.Async()
}

func main() {
	initilize()
	log.Info("main init done")
}
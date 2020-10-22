package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/liuhengloveyou/ipdb/api"
	"github.com/liuhengloveyou/ipdb/common"
	"github.com/liuhengloveyou/ipdb/db"
)

var addr = flag.String("addr", ":10000", "http listen addr.")
var ipdbFile = flag.String("ipfile", "./ipip.txtx", "ip databases file.")
var logDir = flag.String("logdir", "./logs/", "log dir.")
var logLevel = flag.String("loglevel", "debug", "log level.")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	if err := common.InitLogger(*logDir, *logLevel); err != nil {
		panic(err)
	}
	defer common.Logger.Sync()

	if err := db.InitGEO("ipip", *ipdbFile); err != nil {
		panic(err)
	}

	fmt.Println("go http ", *addr)
	if err := api.InitHttpApi(*addr); err != nil {
		panic("HTTPAPI: " + err.Error())
	}
}

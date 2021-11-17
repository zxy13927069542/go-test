package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

/**
定制四种不同级别的日志
*/
var (
	Trace   *log.Logger //记录所有日志
	Info    *log.Logger //记录普通日志
	Warning *log.Logger //记录注意日志
	Error   *log.Logger //记录错误日志
)

/**
初始化日志
*/
func init() {
	//创建错误日志的输出文件
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("创建错误日志文件失败", err)
	}

	//定制Trace等级日志
	Trace = log.New(ioutil.Discard, "[Trace] ", log.Ldate|log.Ltime|log.Llongfile)
	//定制Info等级日志
	Info = log.New(os.Stdout, "[Info] ", log.Ldate|log.Ltime|log.Llongfile)
	//定制Warning等级日志
	Warning = log.New(os.Stdout, "[Warning] ", log.Ldate|log.Ltime|log.Llongfile)
	//定制Error等级日志,将日志同时输出到file文件里和os.Stderr
	Error = log.New(io.MultiWriter(file, os.Stderr), "[Error] ", log.Ldate|log.Ltime|log.Llongfile)
}

func main() {
	Trace.Println("这里是Trace日志")
	Info.Println("这里是Info日志")
	Warning.Println("这里是Warning日志")
	Error.Println("这里是Error日志")
}

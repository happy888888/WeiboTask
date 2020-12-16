// @Title        WeiboTask
// @Description  包括程序入口，日志初始化
// @Author       星辰
// @Update
package main

import (
	"WeiboClient"
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"sync"
)

// @title         main
// @description   程序入口
// @auth          星辰
// @param
// @return
func main() {
	var configPath, logPath string
	flag.StringVar(&logPath, "l", "./log.log", "日志文件路径,默认为./log.log")
	logFile, err := os.OpenFile(logPath, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0600)
	if err != nil {
		println(err)
		initLog(nil)
	}else{
		initLog(logFile)
		defer logFile.Close()
	}
	flag.StringVar(&configPath, "c", "./config.json", "配置文件路径,默认为./config.json")
	config, err := LoadConfig(configPath)
	if err != nil {
		log.Println("配置文件加载失败:"+err.Error())
		os.Exit(6)
	}
	defer sendToServerChan(config.SCKEY)
	wb := WeiboClient.New()
	if wb.LoginByCookies(config.Cookies) {
		defer func() { config.Cookies = wb.GetCookies(); _ = SaveConfig(configPath, config) }()
		var wg sync.WaitGroup
		wg.Add(1)
		go runTasks(wb, &wg)
		wg.Wait()
	}else{
		log.Println("登录失败")
	}
}

// @title         initLog
// @description   初始化日志
// @auth          星辰
// @param         logFile       *os.File   日志文件
// @return
func initLog(logFile *os.File) {
	log.SetFlags(log.LstdFlags)
	mBuffer = bytes.NewBufferString("")
	logIo := io.MultiWriter(os.Stdout, mBuffer)
	if logFile != nil {
		logIo = io.MultiWriter(logIo, logFile)
	}
	log.SetOutput(logIo)
}
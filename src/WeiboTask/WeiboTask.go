// @Title        WeiboTask
// @Description  包括程序入口，日志初始化
// @Author       星辰
// @Update
package main

import (
	"WeiboClient"
	"WeiboTask/Deamon"
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"time"
)

// @title         main
// @description   程序入口
// @auth          星辰
// @param
// @return
func main() {
	var configPath, logPath string
	var isDeamon, isLoop bool
	flag.StringVar(&logPath, "l", "./WeiboTask.log", "日志文件路径,默认为./log.log")
	flag.StringVar(&ConfigPath, "c", "./config.json", "配置文件路径,默认为./config.json")
	flag.BoolVar(&isDeamon, "D", false, "是否以守护进程方式循环运行")
	flag.BoolVar(&isLoop, "L", false, "循环运行")
	flag.Parse()

	if isDeamon {
		Deamon.InitDeamon()
	}

	logFile, err := os.OpenFile(logPath, os.O_CREATE | os.O_APPEND | os.O_RDWR, 0600)
	if err != nil {
		println(err)
		initLog(nil)
	}else{
		initLog(logFile)
		defer logFile.Close()
	}
	err = LoadConfig()
	if err != nil {
		log.Println("配置文件加载失败:"+err.Error())
		os.Exit(6)
	}
	if isDeamon || isLoop {
		runLoop(configPath)
	}else{
		runOnce(false)
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

// @title         runOnce
// @description   单次运行任务
// @auth          星辰
// @param         configPath       string   配置文件路径
// @param         reloadConfig     bool     执行前是否重载配置文件
// @return
func runOnce(reloadConfig bool) {
	if reloadConfig {
		err := LoadConfig()
		if err != nil {
			log.Println("配置文件加载失败:"+err.Error())
			os.Exit(6)
		}
	}
	defer sendToServerChan()
	wb := WeiboClient.New(MyConfig.C, MyConfig.S)
	if wb.LoginByCookies(MyConfig.Cookies) {
		defer func() { MyConfig.Cookies = wb.GetCookies(); _ = SaveConfig() }()
		runTasks(wb)
	}else{
		log.Println("登录失败")
	}
}

// @title         runLoop
// @description   周期运行任务
// @auth          星辰
// @param         configPath       string   配置文件路径
// @return
func runLoop(configPath string) {
	Now := time.Now()
	todayTime := time.Date(Now.Year(), Now.Month(), Now.Day(), MyConfig.Stime[0], MyConfig.Stime[1], 0, 0, time.Local)
	tomorrowTime := time.Date(Now.Year(), Now.Month(), Now.Day() + 1, MyConfig.Stime[0], MyConfig.Stime[1], 0, 0, time.Local)
	if Now.Before(todayTime) {
		time.Sleep(todayTime.Sub(Now))
	}else{
		time.Sleep(tomorrowTime.Sub(Now))
	}
	for {
		go runOnce(true)
		time.Sleep(24 * time.Hour)
	}
}

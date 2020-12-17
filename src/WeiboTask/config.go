// @Title        config
// @Description  包括配置文件的加载与存放
// @Author       星辰
// @Update
package main

import (
	"WeiboClient"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config 配置文件结构
type Config struct{
	SCKEY string                  //server酱的推送key
	Cookies []WeiboClient.Cookie  //cookies用于保持登录状态
	Stime [2]int                  //定时启动时间，格式为{小时, 分钟}
}

var MyConfig Config  //全局配置

// @title         LoadConfig
// @description   加载配置文件
// @auth          星辰
// @param         path          string     配置文件路径(需要读取权限)
// @return
func LoadConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &MyConfig)
	if err != nil {
		return err
	}
	return nil
}

// @title         LoadConfig
// @description   保存配置文件
// @auth          星辰
// @param         path          string     配置文件路径(需要写入权限)
// @return
func SaveConfig(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE,0600)
	if err != nil {
		return err
	}
	defer f.Close()
	s, err := json.Marshal(MyConfig)
	if err != nil {
		return err
	}
	var str bytes.Buffer
	_ = json.Indent(&str, s, "", "    ")
	_, err = f.Write([]byte(str.String()))
	return err
}
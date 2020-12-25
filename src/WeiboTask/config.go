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
	"runtime"
)

// Config 配置文件结构
type Config struct{
	SCKEY string                  //server酱的推送key
	Cookies []WeiboClient.Cookie  //cookies用于保持登录状态
	Stime [2]int                  //定时启动时间，格式为{小时, 分钟}
	C string                      //平台,如"android"
	S string                      //app秘钥,6位字符串
}

var MyConfig Config  //配置文件对象
var ConfigPath string //配置文件路径

// @title         LoadConfig
// @description   加载配置文件
// @auth          星辰
// @param         path          string     配置文件路径(需要读取权限)
// @return
func LoadConfig() error {
	_, err := os.Lstat(ConfigPath)
	if os.IsNotExist(err) && runtime.GOOS == "linux" {
		ConfigPath = "/etc/WeiboTask/config.json"
	}
	f, err := os.Open(ConfigPath)
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
func SaveConfig() error {
	f, err := os.OpenFile(ConfigPath, os.O_WRONLY|os.O_CREATE,0600)
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
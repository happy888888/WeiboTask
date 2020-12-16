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
}

// @title         LoadConfig
// @description   加载配置文件
// @auth          星辰
// @param         path          string     配置文件路径(需要读取权限)
// @return                      Config     配置
func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var rconfig Config
	err = json.Unmarshal(data, &rconfig)
	if err != nil {
		return nil, err
	}
	return &rconfig, nil
}

// @title         LoadConfig
// @description   保存配置文件
// @auth          星辰
// @param         path          string     配置文件路径(需要写入权限)
// @param                       *Config    配置
// @return
func SaveConfig(path string, config *Config) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE,0600)
	if err != nil {
		return err
	}
	defer f.Close()
	s, err := json.Marshal(config)
	if err != nil {
		return err
	}
	var str bytes.Buffer
	_ = json.Indent(&str, s, "", "    ")
	_, err = f.Write([]byte(str.String()))
	return err
}
// @Title        sendMessage
// @Description  消息推送
// @Author       星辰
// @Update
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var mBuffer *bytes.Buffer //用于在内存中保存日志

// @title         sendToServerChan
// @description   推送消息至server酱
// @auth          星辰
// @param
// @return
func sendToServerChan() {
	defer mBuffer.Reset()
	if MyConfig.SCKEY == "" {
		log.Println("未定义消息推送")
		return
	}
	resp, err := http.PostForm(
		"https://sc.ftqq.com/"+MyConfig.SCKEY+".send",
		url.Values{
			"text": {"Weibo_sign_in消息推送"},
			"desp": {strings.Replace(mBuffer.String(), "\n", "\n\n", -1 )},
		},
	)
	if err != nil {
		log.Println("消息推送异常："+err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("消息推送异常："+err.Error())
		return
	}
	var data map[string]interface{}
	print()
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("消息推送读取响应异常："+err.Error())
		return
	}
	log.Println("server酱推送信息："+data["errmsg"].(string))
}

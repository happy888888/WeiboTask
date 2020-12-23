// @Title        AppSignIn
// @Description  微博客户端签到
// @Author       星辰
// @Update
package Tasks

import (
	"WeiboClient"
	"fmt"
	"log"
	"sync"
)

// @title         AppSignIn
// @description   微博客户端签到
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func AppSignIn(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	data, err := w.CheckinSignIn()
	if err != nil {
		log.Println("微博app签到异常:"+err.Error())
		return
	}
	if data["ok"].(float64) != 1 {
		log.Println("微博app签到错误:"+data["msg"].(string))
		return
	}
	sign_in, ok := data["data"].(map[string]interface{})["sign_in"].(map[string]interface{})
	if !ok {
		log.Println("获取微博app签到信息失败")
		return
	}
	continuous := sign_in["continuous"].(float64)
	sign_in = sign_in["sign_in"].(map[string]interface{})
	if sign_in["show"].(float64) == 1 {
		gift := sign_in["content"].(map[string]interface{})["gift"].(map[string]interface{})
		if money, ok := gift["money"]; ok {
			log.Println(fmt.Sprintf("微博app签到获得%s%s,已连续签到%0.0f天",
				money.(map[string]interface{})["value"].(string),
				money.(map[string]interface{})["unit"].(string),
				continuous))
		}
		if points, ok := gift["points"]; ok {
			log.Println(fmt.Sprintf("微博app签到获得%s%s,已连续签到%0.0f天",
				points.(map[string]interface{})["value"].(string),
				points.(map[string]interface{})["unit"].(string),
				continuous))
		}
	}else{
		log.Println(fmt.Sprintf("微博app已连续签到%0.0f天", continuous))
	}
}
// @Title        SuperCheckIn
// @Description  超话签到任务
// @Author       星辰
// @Update
package Tasks

import (
	"WeiboClient"
	"log"
	"regexp"
	"sync"
)

// @title         SuperCheckIn
// @description   超话列表签到
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func SuperCheckIn(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	ch := make(chan [2]string, 3)
	go GetSuperTopics(w, ch)
	for {
		if item, ok := <-ch; ok {
			data, err := w.SuperCheckin(item[1])
			//响应例子{'code': '100000', 'msg': '已签到', 'data': {'tipMessage': '今日签到，经验值+4', 'alert_title': '今日签到 第14482名', 'alert_subtitle': '经验值+4', 'alert_activity': ''}}
			if err != nil {
				log.Println("签到超话("+item[0]+")异常:"+err.Error())
				continue
			}
			if data["code"] == "100000" {
				data = data["data"].(map[string]interface {})
				log.Println("签到超话("+item[0]+")成功:"+data["tipMessage"].(string)+","+data["alert_title"].(string))
			}else{
				log.Println("签到超话("+item[0]+")失败:"+data["msg"].(string))
			}
		} else {
			break
		}
	}
}

// @title         GetSuperTopics
// @description   获取关注的超话列表
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         ch         chan [2]string            管道，返回[超话名称, 超话id]
// @return
func GetSuperTopics(w *WeiboClient.WeiboClient, ch chan<- [2]string) {
	reg := regexp.MustCompile("[0-9a-z]{38}")
	sinceId := ""
	for {
		data, err := w.ContainerGetIndex("100803_-_followsuper", sinceId)
		if err != nil {
			log.Println("获取超话列表异常:"+err.Error())
			break
		}
		//python里直接["data"]["cardlistInfo"]["since_id"]
		sinceId, ok := data["data"].(map[string]interface{})["cardlistInfo"].(map[string]interface{})["since_id"]
		if !ok {
			log.Println("获取超话列表sinceId错误")
			break
		}
		//python里直接["data"]["cards"][0]["card_group"]
		cards, ok := data["data"].(map[string]interface{})["cards"].([]interface{})
		if !ok {
			log.Println("获取超话列表cards错误")
			break
		}
		cardGroup, ok := cards[0].(map[string]interface{})["card_group"].([]interface{})
		if !ok {
			log.Println("获取超话列表card_group错误")
			break
		}
		for _, v := range cardGroup {
			item := v.(map[string]interface{})
			if item["card_type"] == "8" {
				name := item["title_sub"].(string)
				id := reg.FindString(item["scheme"].(string))
				ch <- [2]string{name, id}
			}
		}
		if sinceId == "" {
			break
		}
	}
	close(ch)
}

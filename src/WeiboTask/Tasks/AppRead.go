// @Title        AppRead
// @Description  微博app刷关注微博任务
// @Author       星辰
// @Update
package Tasks

import (
	"WeiboClient"
	"log"
	"strconv"
	"sync"
	"time"
)

// @title         AppRead
// @description   微博app刷关注微博
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func AppRead(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for i, c := 0, 0; i < 5 && c < 7; c++ {
		data, err := w.UnreadFriendsTimeline()
		time.Sleep(time.Second * 6)
		if err != nil {
			log.Println("微博app刷关注微博异常:"+err.Error())
			continue
		}
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app刷关注微博失败:"+errmsg.(string))
			continue
		}else{
			log.Println("微博app刷关注微博第"+strconv.Itoa(i)+"次")
		}
		i++
	}
	data, err := w.ScoreClaim("54")
	if err != nil {
		log.Println("微博app刷关注微博任务完成异常:"+err.Error())
		return
	}
	if data["ok"].(float64) == 1 {
		data = data["data"].(map[string]interface{})
		log.Println("微博app刷关注微博任务完成")
	}else{
		log.Println("微博app刷关注微博任务完成失败:"+data["msg"].(string))
	}
}
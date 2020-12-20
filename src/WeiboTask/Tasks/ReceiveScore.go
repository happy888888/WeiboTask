// @Title        ReceiveScore
// @Description  每日积分获取
// @Author       星辰
// @Update
package Tasks

import (
	"WeiboClient"
	"log"
	"sync"
)

// @title         ReceiveScore
// @description   每日积分获取
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func ReceiveScore(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	data, err := w.SuperReceiveScore()
	if err != nil {
		log.Println("每日积分获取异常:"+err.Error())
		return
	}
	if data["code"].(float64) == 100000 {
		data = data["data"].(map[string]interface{})
		log.Println("每日积分获取增加："+data["add_score"].(string)+"积分")
	}else{
		log.Println("每日积分获取失败："+data["msg"].(string))
	}
}

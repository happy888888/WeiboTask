// @Title        AppFollowUser
// @Description  微博app关注用户任务
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

// @title         AppFollowUser
// @description   微博app关注用户
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func AppFollowUser(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	uids := GetUids(w)
	if len(uids) < 5 {
		log.Println("未得到足够的用户，跳过微博app关注取关")
		return
	}
	if wg != nil {
		wg.Add(1)
	}
	go AppFollowWithUndo(w, uids, wg)
}

// @title         GetUids
// @description   微博app获取首页帖子列表
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @return        uids       []string                  帖子mid列表
func GetUids(w *WeiboClient.WeiboClient) (uids []string) {
	data, err := w.CardList("231093_-_selfrecomm")
	if err != nil {
		log.Println("微博app获取用户列表异常:"+err.Error())
		return
	}
	if errmsg, ok := data["errmsg"]; ok {
		log.Println("微博app获取用户列表失败:"+errmsg.(string))
		return
	}
	cards := data["cards"].([]interface{})
	for _, card := range cards {
		card_group := card.(map[string]interface{})["card_group"].([]interface{})
		for _, group := range card_group {
			uid := group.(map[string]interface{})["user"].(map[string]interface{})["id"].(float64)
			uids = append(uids, strconv.Itoa(int(uid)))
		}
	}
	return
}

// @title         AppFollowWithUndo
// @description   微博app关注用户并取关
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         uids       []string                  用户id列表
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func AppFollowWithUndo(w *WeiboClient.WeiboClient, uids []string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for ii, uid := range uids {
		if ii >= 5 {
			break
		}
		data, err := w.AppFriendshipsCreate(uid)
		time.Sleep(time.Second * 6)
		if err != nil {
			log.Println("微博app关注用户"+uid+"异常:"+err.Error())
			continue
		}
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app关注用户"+uid+"失败:"+errmsg.(string))
			continue
		}else{
			log.Println("微博app关注用户"+uid+"成功")
		}
		data, err = w.AppFriendshipsDestroy(uid)
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app关注后取关用户失败:"+errmsg.(string))
		}else{
			log.Println("微博app关注后取关用户成功")
		}
	}
	data, err := w.ScoreClaim("8")
	if err != nil {
		log.Println("微博app用户关注任务完成异常:"+err.Error())
		return
	}
	if data["ok"].(float64) == 1 {
		data = data["data"].(map[string]interface{})
		log.Println("微博app用户关注任务完成")
	}else{
		log.Println("微博app用户关注任务完成失败:"+data["msg"].(string))
	}
}
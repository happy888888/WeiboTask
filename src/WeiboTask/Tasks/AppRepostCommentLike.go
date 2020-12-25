// @Title        AppRepostCommentLike
// @Description  微博app转发评论点赞任务
// @Author       星辰
// @Update
package Tasks

import (
	"WeiboClient"
	"log"
	"sync"
	"time"
)

// @title         AppRepostCommentLike
// @description   微博app转发评论点赞
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func AppRepostCommentLike(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	mids := GetMids(w)
	if len(mids) < 2 {
		log.Println("未得到足够的帖子，跳过微博app转发评论点赞")
		return
	}
	if wg != nil {
		wg.Add(3)
	}
	go AppRepostWithDel(w, mids, wg)
	go AppCommentsWithDel(w, mids, wg)
	go ApplikeWithUndo(w, mids, wg)
}

// @title         GetMids
// @description   微博app获取首页帖子列表
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @return        mids       []string                  帖子mid列表
func GetMids(w *WeiboClient.WeiboClient) (mids []string) {
	data, err := w.CardList("102803")
	if err != nil {
		log.Println("微博app获取帖子列表异常:"+err.Error())
		return
	}
	if errmsg, ok := data["errmsg"]; ok {
		log.Println("微博app获取帖子列表失败:"+errmsg.(string))
		return
	}
	cards := data["cards"].([]interface{})
	for _, card := range cards {
		mid := card.(map[string]interface{})["mblog"].(map[string]interface{})["mid"].(string)
		mids = append(mids, mid)
	}
	return
}

// @title         AppRepostWithDel
// @description   微博app转发帖子并删除
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         mids       []string                  帖子id列表
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func AppRepostWithDel(w *WeiboClient.WeiboClient, mids []string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for ii, mid := range mids {
		if ii >= 3 {
			break
		}
		data, err := w.AppRepost(mid)
		time.Sleep(time.Second * 6)
		if err != nil {
			log.Println("微博app转发帖子"+mid+"异常:"+err.Error())
			continue
		}
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app转发帖子"+mid+"失败:"+errmsg.(string))
			continue
		}else{
			log.Println("微博app转发帖子"+mid+"成功")
		}
		id := data["statuses"].(map[string]interface{})["fast_reposted_by_copy"].(map[string]interface{})["mid"].(string)
		data, err = w.AppDestroy(id)
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app转发后删除帖子"+id+"失败:"+errmsg.(string))
		}else{
			log.Println("微博app转发后删除帖子"+id+"成功")
		}
	}
	data, err := w.ScoreClaim("13")
	if err != nil {
		log.Println("微博app转发任务完成异常:"+err.Error())
		return
	}
	if data["ok"].(float64) == 1 {
		data = data["data"].(map[string]interface{})
		log.Println("微博app转发任务完成")
	}else{
		log.Println("微博app转发任务完成失败:"+data["msg"].(string))
	}
}

// @title         AppCommentsWithDel
// @description   微博app评论帖子并删除
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         mids       []string                  帖子id列表
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func AppCommentsWithDel(w *WeiboClient.WeiboClient, mids []string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for ii, mid := range mids {
		if ii >= 2 {
			break
		}
		data, err := w.AppCommentsCreate(mid, "o(￣▽￣)ｄ")
		time.Sleep(time.Second * 6)
		if err != nil {
			log.Println("微博app评论帖子"+mid+"异常:"+err.Error())
			continue
		}
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app评论帖子"+mid+"失败:"+errmsg.(string))
			continue
		}else{
			log.Println("微博app评论帖子"+mid+"成功")
		}
		id := data["mid"].(string)
		data, err = w.AppCommentsDestroy(id)
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app评论后删除"+id+"失败:"+errmsg.(string))
		}else{
			log.Println("微博app评论后删除"+id+"成功")
		}
	}
	data, err := w.ScoreClaim("12")
	if err != nil {
		log.Println("微博app评论任务完成异常:"+err.Error())
		return
	}
	if data["ok"].(float64) == 1 {
		data = data["data"].(map[string]interface{})
		log.Println("微博app评论任务完成")
	}else{
		log.Println("微博app评论任务完成失败:"+data["msg"].(string))
	}
}

// @title         ApplikeWithUndo
// @description   微博app点赞帖子并取消
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         mids       []string                  帖子id列表
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func ApplikeWithUndo(w *WeiboClient.WeiboClient, mids []string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for ii, mid := range mids {
		if ii >= 2 {
			break
		}
		data, err := w.AppSetLike(mid)
		time.Sleep(time.Second * 6)
		if err != nil {
			log.Println("微博app点赞帖子"+mid+"异常:"+err.Error())
			continue
		}
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app点赞帖子"+mid+"失败:"+errmsg.(string))
			continue
		}else{
			log.Println("微博app点赞帖子"+mid+"成功")
		}
		data, err = w.AppCancelLike(mid)
		if errmsg, ok := data["errmsg"]; ok {
			log.Println("微博app取消点赞"+mid+"失败:"+errmsg.(string))
		}else{
			log.Println("微博app取消点赞"+mid+"成功")
		}
	}
	data, err := w.ScoreClaim("11")
	if err != nil {
		log.Println("微博app点赞任务完成异常:"+err.Error())
		return
	}
	if data["ok"].(float64) == 1 {
		data = data["data"].(map[string]interface{})
		log.Println("微博app点赞任务完成")
	}else{
		log.Println("微博app点赞任务完成失败:"+data["msg"].(string))
	}
}
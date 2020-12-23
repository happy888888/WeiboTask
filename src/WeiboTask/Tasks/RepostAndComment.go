// @Title        RepostAndComment
// @Description  超话转发评论
// @Author       星辰
// @Update
package Tasks

import (
	"WeiboClient"
	"errors"
	"log"
	"regexp"
	"sync"
	"time"
)

// @title         RepostAndComment
// @description   完成转发和评论超话帖子任务
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func RepostAndComment(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	id, err := GetFirstSuperTopic(w)
	if err != nil {
		log.Println("获取超话异常："+err.Error())
		return
	}
	list, err := GetComposeList(w, id)
	if err != nil {
		log.Println("获取超话帖子异常："+err.Error())
		return
	}
	var mywg sync.WaitGroup
	mywg.Add(2)
	go RepostWithDel(w, list, &mywg)
	go CommentWithDel(w, list, &mywg)
	mywg.Wait()
}

// @title         RepostWithDel
// @description   转发2个帖子并删除转发的帖子
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         list       []string                  超话帖子列表
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func RepostWithDel(w *WeiboClient.WeiboClient, list []string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for ii, mid := range list {
		if ii >= 2 {
			break
		}
		data, err := w.ComposeRepost(mid, "🔔")
		time.Sleep(time.Second * 10)
		if err != nil {
			log.Println("转发超话帖子异常："+err.Error())
			continue
		}
		if data["ok"].(float64) == 1 {
			log.Println("转发帖子"+mid+"成功")
		}else{
			log.Println("转发帖子"+mid+"失败:"+data["msg"].(string))
			continue
		}
		mid = data["data"].(map[string]interface{})["mid"].(string)
		data, err = w.DelMyblog(mid)
		if err != nil {
			log.Println("删除超话帖子异常："+err.Error())
			continue
		}
		if data["ok"].(float64) == 1 {
			log.Println("删除帖子"+mid+"成功")
		}else{
			log.Println("删除帖子"+mid+"失败:"+data["msg"].(string))
		}
	}
}

// @title         CommentWithDel
// @description   评论6个帖子并删除评论
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         list       []string                  超话帖子列表
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func CommentWithDel(w *WeiboClient.WeiboClient, list []string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	for ii, mid := range list {
		if ii >= 6 {
			break
		}
		data, err := w.CommentsCreate(mid, "🔔")
		time.Sleep(time.Second * 6)
		if err != nil {
			log.Println("评论超话帖子"+mid+"异常："+err.Error())
			continue
		}
		if data["ok"].(float64) == 1 {
			log.Println("评论帖子"+mid+"成功")
		}else{
			log.Println("评论帖子"+mid+"失败:"+data["msg"].(string))
			continue
		}
		mid = data["data"].(map[string]interface{})["mid"].(string)
		data, err = w.CommentsDestroy(mid)
		if err != nil {
			log.Println("删除帖子评论"+mid+"异常："+err.Error())
			continue
		}
		if data["ok"].(float64) == 1 {
			log.Println("删除帖子评论"+mid+"成功")
		}else{
			log.Println("删除帖子评论"+mid+"失败:"+data["msg"].(string))
		}
	}
}

// @title         GetComposeList
// @description   获取超话帖子第一页
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @return                   []string                  超话评论
func GetComposeList(w *WeiboClient.WeiboClient, id string) (m []string, err error) {
	data, err := w.ContainerGetIndex(id, "")
	if err != nil {
		return
	}
	cards, ok := data["data"].(map[string]interface{})["cards"].([]interface{})
	if !ok {
		return m, errors.New("获取超话话题列表cards错误")
	}
	for _, card := range cards {
		cardJson := card.(map[string]interface{})
		if _, ok := cardJson["card_group"]; !ok {
			continue
		}
		cardGroup, ok := cardJson["card_group"].([]interface{})
		if !ok {
			continue
		}
		for _, v := range cardGroup {
			item := v.(map[string]interface{})
			if item["card_type"] == "9" {
				m = append(m, item["mblog"].(map[string]interface{})["mid"].(string))
			}
		}
	}
	return
}

// @title         GetFirstSuperTopic
// @description   获取一个关注的超话
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @return                   string                    超话id
func GetFirstSuperTopic(w *WeiboClient.WeiboClient) (string, error) {
	reg := regexp.MustCompile("[0-9a-z]{38}")
	data, err := w.ContainerGetIndex("100803_-_followsuper", "")
	if err != nil {
		return "", err
	}
	cards, ok := data["data"].(map[string]interface{})["cards"].([]interface{})
	if !ok {
		return "", errors.New("获取超话列表cards错误")
	}
	for _, card := range cards {
		cardJson := card.(map[string]interface{})
		if cardJson["card_type"] == "11" {
			cardGroup, ok := cardJson["card_group"].([]interface{})
			if !ok {
				break
			}
			for _, v := range cardGroup {
				item := v.(map[string]interface{})
				if item["card_type"] == "8" {
					id := reg.FindString(item["scheme"].(string))
					return id, nil
				}
			}
		}
	}
	return "", errors.New("没获取到超话")
}
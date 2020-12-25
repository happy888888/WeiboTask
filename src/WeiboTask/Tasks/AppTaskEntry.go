// @Title        AppTaskEntry
// @Description  微博app任务入口
// @Author       星辰
// @Update
package Tasks

import (
	"WeiboClient"
	"log"
	"sync"
)

// @title         AppTaskEntry
// @description   启动微博app任务
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func AppTaskEntry(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	if w.C == "" || w.S == "" || !weiboAppOK(w){
		log.Println("没有正确的S参数和C参数,跳过执行微博app任务")
		return
	}
	if wg != nil {
		wg.Add(3)
	}
	go AppFollowUser(w, wg)
	go AppRepostCommentLike(w, wg)
	go AppRead(w, wg)
}

// @title         weiboAppOK
// @description   判断微博app s参数是否有效
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @return                   bool
func weiboAppOK(w *WeiboClient.WeiboClient) bool {
	data, err := w.UrlSafe()
	if err != nil {
		log.Println("微博app参赛验证异常:"+err.Error())
		return false
	}
	if errmsg, ok := data["errmsg"]; ok {
		log.Println("微博app参数验证失败(请重新填写S参数和C参数):"+errmsg.(string))
		return false
	}else{
		return true
	}
}
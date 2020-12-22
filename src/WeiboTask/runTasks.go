// @Title        runTasks
// @Description  用户任务的加载和执行
// @Author       星辰
// @Update
package main

import (
	"WeiboClient"
	"WeiboTask/Tasks"
	"sync"
)

// @title         runTasks
// @description   启动用户每日任务
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @return
func runTasks(w *WeiboClient.WeiboClient) {
	var mywg sync.WaitGroup
	mywg.Add(3)
	go Tasks.SuperCheckIn(w, &mywg)
	go Tasks.ReceiveScore(w, &mywg)
	go Tasks.RepostAndComment(w, &mywg)
	mywg.Wait()
}

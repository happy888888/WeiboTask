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
// @description   启动用户任务
// @auth          星辰
// @param         w          *WeiboClient.WeiboClient  微博客户端
// @param         wg         *sync.WaitGroup           等待组，保持程序同步
// @return
func runTasks(w *WeiboClient.WeiboClient, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	var mywg sync.WaitGroup
	mywg.Add(1)
	go Tasks.SuperCheckIn(w, &mywg)
	mywg.Wait()
}
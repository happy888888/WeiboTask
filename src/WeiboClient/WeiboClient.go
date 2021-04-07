// @Title        WeiboClient
// @Description  提供一个封装的请求微博的http客户端
// @Author       星辰
// @Update
package WeiboClient

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

// WeiboClient   微博客户端对象，包含一个http客户端及cookieJar
type WeiboClient struct{
	client *http.Client
	C string
	S string
	F string
}

// @title         init
// @description   初始化WeiboClient对象
// @auth          星辰
// @param
// @return
func (w *WeiboClient) init() {
	jar, _ := cookiejar.New(nil)
	w.client = &http.Client{
		Jar: jar,
		Timeout: 30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse /* 不进入重定向 */
		},
	}
}

// @title         New
// @description   创建一个WeiboClient对象并初始化
// @auth          星辰
// @return        WeiboClient对象
func New(c string, s string, f string) *WeiboClient {
	ret := new(WeiboClient)
	ret.init()
	ret.C = c
	ret.S = s
	ret.F = f
	return ret
}

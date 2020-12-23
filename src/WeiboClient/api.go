// @Title        api
// @Description  提供封装的微博请求接口
// @Author       星辰
// @Update
package WeiboClient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// @title         apiConfig
// @description   用于获取st参数，cookie中XSRF-TOKEN也等于cookie
// @auth          星辰
// @param
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) apiConfig() (map[string]interface{}, error) {
	resp, err := w.client.Get("https://m.weibo.cn/api/config")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         getST
// @description   获取st参数
// @auth          星辰
// @param
// @return                   string                   st的值
func (w *WeiboClient) getST() string {
	st := w.getCookie("XSRF-TOKEN", ".m.weibo.cn")
	if st != "" {
		return st
	}
	data, err := w.apiConfig()
	if err != nil {
		return ""
	}
	st, ok := data["data"].(map[string]interface{})["st"].(string)
	if !ok {
		return ""
	}
	return st
}

// @title         GeneralButton
// @description   微博按钮接口
// @auth          星辰
// @param         api        string                  "接口地址"
// @param         id         string                  "请求id"
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) GeneralButton(api string, id string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://weibo.com/p/aj/general/button?api="+api+"&id="+id)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         superCheckin
// @description   超话签到
// @auth          星辰
// @param         id         string                  "超话id"
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) SuperCheckin(id string) (map[string]interface{}, error){
	var data map[string]interface{}
	data, err := w.GeneralButton("http://i.huati.weibo.com/aj/super/checkin", id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         ContainerGetIndex
// @description   获取目录
// @auth          星辰
// @param         id         string                  "容器id"
// @param         sinceId    string                  "容器id"
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) ContainerGetIndex(id string, sinceId string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://m.weibo.cn/api/container/getIndex?containerid="+id+"&since_id="+sinceId)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         SuperReceiveScore
// @description   超话每日积分获取
// @auth          星辰
// @param
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) SuperReceiveScore() (map[string]interface{}, error){
	req, err := http.NewRequest("POST",
		"https://huati.weibo.cn/aj/super/receivescore",
		strings.NewReader("type=REQUEST&user_score=999"),
		)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://huati.weibo.cn/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         ComposeRepost
// @description   转发帖子
// @auth          星辰
// @param         mid        string                   帖子id
// @param         content    string                   转发内容
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) ComposeRepost(mid string, content string) (map[string]interface{}, error){
	req, err := http.NewRequest("POST",
		"https://m.weibo.cn/api/statuses/repost",
		strings.NewReader("id="+mid+"&content="+content+"&mid="+mid+"&st="+w.getST()+"&_spr=screen:1920x1080"),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://m.weibo.cn/compose/repost")
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         DelMyblog
// @description   删除自己的帖子
// @auth          星辰
// @param         mid        string                   帖子id
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) DelMyblog(mid string) (map[string]interface{}, error){
	req, err := http.NewRequest("POST",
		"https://m.weibo.cn/profile/delMyblog",
		strings.NewReader("mid="+mid+"&st="+w.getST()+"&_spr=screen:1920x1080"),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://m.weibo.cn/profile/")
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         CommentsCreate
// @description   帖子下发表评论
// @auth          星辰
// @param         mid        string                   帖子id
// @param         content    string                   转发内容
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) CommentsCreate(mid string, content string) (map[string]interface{}, error){
	req, err := http.NewRequest("POST",
		"https://m.weibo.cn/api/comments/create",
		strings.NewReader("content="+content+"&mid="+mid+"&st="+w.getST()+"&_spr=screen:1920x1080"),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://m.weibo.cn/detail/")
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         CommentsDestroy
// @description   删除自己的评论
// @auth          星辰
// @param         cid        string                   评论id
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) CommentsDestroy(cid string) (map[string]interface{}, error){
	req, err := http.NewRequest("POST",
		"https://m.weibo.cn/comments/destroy",
		strings.NewReader("cid="+cid+"&st="+w.getST()+"&_spr=screen:1920x1080"),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://m.weibo.cn/detail/")
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// @title         checkinSignIn
// @description   微博app签到
// @auth          星辰
// @param
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) CheckinSignIn() (map[string]interface{}, error){
	req, err := http.NewRequest("GET",
		"https://m.weibo.cn/c/checkin/ug/v2/signin/signin?from=10AC395010&hash=task&luicode=20000301&st="+w.getST(),
		nil,
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "https://m.weibo.cn/c/checkin")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android; none; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.136 Mobile Safari/537.36 Weibo (android)")
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
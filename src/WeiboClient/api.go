// @Title        api
// @Description  提供封装的微博请求接口
// @Author       星辰
// @Update
package WeiboClient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
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
// @param         sinceId    string                  "起点id"
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

// @title         UrlSafe
// @description   判断app S和C参数是否有效
// @auth          星辰
// @return                            map[string]interface{}   接口返回值
func (w *WeiboClient) UrlSafe() (map[string]interface{}, error){
	resp, err := w.client.Get("https://api.weibo.cn/2/client/url_safe"+
		"?c="+w.C+
		"&s="+w.S+
		"&from="+w.F+
		"&gsid="+w.getCookie("SUB", ".weibo.cn"),
	)
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

// @title         CardList
// @description   获取帖子列表
// @auth          星辰
// @param         containerid         string                  "容器id"
// @return                            map[string]interface{}   接口返回值
func (w *WeiboClient) CardList(containerid string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://api.weibo.cn/2/cardlist"+
		"?c="+w.C+
		"&s="+w.S+
		"&from="+w.F+
		"&gsid="+w.getCookie("SUB", ".weibo.cn")+
		"&page="+"1"+
		"&count="+"20"+
		"&containerid="+containerid,
		)
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

// @title         AppRepost
// @description   微博app转发帖子
// @auth          星辰
// @param         id             string                   帖子id
// @return                       map[string]interface{}   接口返回值
func (w *WeiboClient) AppRepost(id string) (map[string]interface{}, error){
	bodyMap := map[string]string {
		"c":w.C,
		"s":w.S,
		"id":id,
		"share_source":"0",
		"is_fast":"1",
		"gsid":w.getCookie("SUB", ".weibo.cn"),
		"from":w.F,
	}
	var b bytes.Buffer
	wt := multipart.NewWriter(&b)
	for key, value := range bodyMap {
		fw, err := wt.CreateFormField(key)
		if err != nil {
			wt.Close()
			return nil, err
		}
		_, err = fw.Write(str2bytes(value))
		if err != nil {
			wt.Close()
			return nil, err
		}
	}
	wt.Close()
	req, err := http.NewRequest("POST",
		"https://api.weibo.cn/2/statuses/repost"+
			"?c="+w.C+
			"&s="+w.S+
			"&id="+id+
			"&from="+w.F+
			//"&gsid="+w.getCookie("SUB", ".weibo.cn"),
			"&gsid="+bodyMap["gsid"],
		&b,
	)
	req.Header.Set("Content-Type", wt.FormDataContentType())
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

// @title         AppDestroy
// @description   微博app删除帖子
// @auth          星辰
// @param         id                  string                  "帖子id"
// @return                            map[string]interface{}   接口返回值
func (w *WeiboClient) AppDestroy(id string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://api.weibo.cn/2/statuses/destroy"+
		"?c="+w.C+
		"&s="+w.S+
		"&id="+id+
		"&from="+w.F+
		"&gsid="+w.getCookie("SUB", ".weibo.cn"),
	)
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

// @title         AppCommentsCreate
// @description   微博app评论帖子
// @auth          星辰
// @param         id             string                   帖子id
// @param         content        string                   评论内容
// @return                       map[string]interface{}   接口返回值
func (w *WeiboClient) AppCommentsCreate(id string, content string) (map[string]interface{}, error){
	bodyMap := map[string]string {
		"style":"LIGHT",
		"comment":content,
		"id":id,
	}
	var b bytes.Buffer
	wt := multipart.NewWriter(&b)
	for key, value := range bodyMap {
		fw, err := wt.CreateFormField(key)
		if err != nil {
			wt.Close()
			return nil, err
		}
		_, err = fw.Write(str2bytes(value))
		if err != nil {
			wt.Close()
			return nil, err
		}
	}
	wt.Close()
	req, err := http.NewRequest("POST",
		"https://api.weibo.cn/2/comments/create"+
			"?c="+w.C+
			"&s="+w.S+
			"&fromlog="+"100016356549855"+
			"&featurecode="+"10000001"+
			"&from="+w.F+
			"&lfid="+"guanzhu"+
			"&gsid="+w.getCookie("SUB", ".weibo.cn"),
		&b,
	)
	req.Header.Set("Content-Type", wt.FormDataContentType())
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

// @title         AppCommentsDestroy
// @description   微博app删除评论
// @auth          星辰
// @param         id                  string                  "评论id"
// @return                            map[string]interface{}   接口返回值
func (w *WeiboClient) AppCommentsDestroy(id string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://api.weibo.cn/2/comments/destroy"+
		"?c="+w.C+
		"&s="+w.S+
		"&cid="+id+
		"&from="+w.F+
		"&gsid="+w.getCookie("SUB", ".weibo.cn"),
	)
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

// @title         AppFriendshipsCreate
// @description   微博app关注用户
// @auth          星辰
// @param         id                  string                  "用户id"
// @return                            map[string]interface{}   接口返回值
func (w *WeiboClient) AppFriendshipsCreate(id string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://api.weibo.cn/2/friendships/create"+
		"?c="+w.C+
		"&s="+w.S+
		"&uid="+id+
		"&trim_level="+"1"+
		"&invite="+"0"+
		"&able_recommend="+"0"+
		"&featurecode="+"10000326"+
		"&from="+w.F+
		"&gsid="+w.getCookie("SUB", ".weibo.cn"),
	)
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

// @title         AppFriendshipsDestroy
// @description   微博app取消关注用户
// @auth          星辰
// @param         id                  string                  "用户id"
// @return                            map[string]interface{}   接口返回值
func (w *WeiboClient) AppFriendshipsDestroy(id string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://api.weibo.cn/2/friendships/destroy"+
		"?c="+w.C+
		"&s="+w.S+
		"&uid="+id+
		"&trim_level="+"0"+
		"&trim="+"1"+
		"&from="+w.F+
		"&gsid="+w.getCookie("SUB", ".weibo.cn"),
	)
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

// @title         AppSetLike
// @description   微博app点赞帖子
// @auth          星辰
// @param         id                  string                  "帖子id"
// @return                            map[string]interface{}   接口返回值
func (w *WeiboClient) AppSetLike(id string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://api.weibo.cn/2/like/set_like"+
		"?c="+w.C+
		"&s="+w.S+
		"&id="+id+
		"&sourcetype="+"feed"+
		"&from="+w.F+
		"&gsid="+w.getCookie("SUB", ".weibo.cn"),
	)
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

// @title         AppCancelLike
// @description   微博app取消点赞帖子
// @auth          星辰
// @param         id                  string                  "帖子id"
// @return                            map[string]interface{}   接口返回值
func (w *WeiboClient) AppCancelLike(id string) (map[string]interface{}, error){
	resp, err := w.client.Get("https://api.weibo.cn/2/like/cancel_like"+
		"?c="+w.C+
		"&s="+w.S+
		"&id="+id+
		"&sourcetype="+"feed"+
		"&from="+w.F+
		"&gsid="+w.getCookie("SUB", ".weibo.cn"),
	)
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

// @title         UnreadFriendsTimeline
// @description   刷新未读帖子
// @auth          星辰
// @param
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) UnreadFriendsTimeline() (map[string]interface{}, error){
	req, err := http.NewRequest("POST",
		"https://api.weibo.cn/2/statuses/unread_friends_timeline"+
		"?orifid="+"guanzhu"+
		"&fromlog="+"100016356549855"+
		"&featurecode="+"10000001"+
		"&c="+w.C+
		"&s="+w.S+
		"&from="+w.F+
	    "&lfid="+"guanzhu"+
		"&gsid="+w.getCookie("SUB", ".weibo.cn"),
		strings.NewReader("since_id=0"),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

// @title         ScoreClaim
// @description   微博app领取积分
// @auth          星辰
// @param         actionId   string                   任务id
// @return                   map[string]interface{}   接口返回值
func (w *WeiboClient) ScoreClaim(actionId string) (map[string]interface{}, error){
	req, err := http.NewRequest("GET",
		"https://m.weibo.cn/c/checkin/ug/score/claim"+
		"?action_id="+actionId+
		"&st="+w.getST(),
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
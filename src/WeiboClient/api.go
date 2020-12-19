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
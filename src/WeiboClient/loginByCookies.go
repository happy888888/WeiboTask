// @Title        loginByCookies
// @Description  提供使用cookie登录微博的方法，利用的是新浪的单点登录机制
// @Author       星辰
// @Update
package WeiboClient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"unsafe"
)

// Cookie   定义Cookie结构
type Cookie struct{
	Name string `json:"name"`     // cookie名称
	Value string `json:"value"`   // cookie值
	Domain string `json:"domain"` // cookie有效域名
}

// @title         LoginByCookies
// @description   通过cookie登录微博
// @auth          星辰
// @param         cookies        []Cookie     "用于登录微博的cookie"
// @return                       bool         "是否登录成功"
func (w *WeiboClient) LoginByCookies(cookies []Cookie) bool {
	jar := w.client.Jar
	for _, cookie := range cookies {
		jar.SetCookies(
			&url.URL{
				Scheme: "https",
				Host: cookie.Domain,
			},
			[]*http.Cookie{
				&http.Cookie{
					Name: cookie.Name,
					Value: cookie.Value,
					Domain: cookie.Domain,
				},
			})
	}
	return w.loginWeiboCom() && w.loginWeiboCn()
}

// @title         GetCookies
// @description   获取目前保存的所有微博cookie
// @auth          星辰
// @param
// @return                       []Cookie         "所有微博cookie"
func (w *WeiboClient) GetCookies() []Cookie {
	domains := []string{".login.sina.com.cn", ".m.weibo.cn", "huati.weibo.cn", ".sina.com.cn", ".weibo.com", ".weibo.cn"}
	var rcookies []Cookie
	for _, domain := range domains {
		cookies := w.client.Jar.Cookies(
			&url.URL{
				Scheme: "https",
				Host: domain,
			})
		for _, cookie := range cookies {
			rcookies = append(rcookies,
				Cookie{
					Name: cookie.Name,
					Value: cookie.Value,
					Domain: domain,
				})
		}
	}
	return rcookies
}

// @title         getCookie
// @description   获取指定cookie
// @auth          星辰
// @param         name           string           cookie名称
// @param         domain         string           cookie所在域名
// @return                       string           cookie的value
func (w *WeiboClient) getCookie(name string, domain string) string {
	cookies := w.client.Jar.Cookies(
		&url.URL{
			Scheme: "https",
			Host: domain,
		})
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}

// @title         isloginWeiboCn
// @description   判断是否登录.weibo.cn，手机版一般用m.weibo.cn
// @auth          星辰
// @param
// @return                       bool         "是否登录成功"
func (w *WeiboClient) isloginWeiboCn() bool {
	resp, err := w.client.Get("https://m.weibo.cn/profile/info")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	return ctype == "application/json; charset=utf-8"
}

// @title         loginWeiboCn
// @description   使用单点登录机制登录.weibo.cn，手机版一般用m.weibo.cn
// @auth          星辰
// @param
// @return                       bool         "是否登录成功"
func (w *WeiboClient) loginWeiboCn() bool {
	if w.isloginWeiboCn() {
		return true
	}
	resp, err := w.client.Get("https://login.sina.com.cn/sso/login.php?url=https%3A%2F%2Fm.weibo.cn%2Fprofile%2Finfo&_rand=1607918849.3653&gateway=1&service=sinawap&entry=sinawap&useticket=1&returntype=META&sudaref=&_client_version=0.6.33")
	if err != nil {
		return false
	}
	if !hasCookie(resp.Cookies(), "SUB") {
		resp.Body.Close()
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return false
	}
	html_text := bytes2str(body)
	reg := regexp.MustCompile("location.replace\\(\\\"(https.*?)\\\"\\)")
	turl := reg.FindStringSubmatch(html_text)
	if len(turl) < 2 || turl[1] == "" {
		return false
	}
	resp, err = w.client.Get(turl[1])
	if err != nil {
		return false
	}
	resp.Body.Close()
	return w.isloginWeiboCn()
}

// @title         isloginWeiboCom
// @description   判断是否登录.weibo.com，网页版一般用www.weibo.com
// @auth          星辰
// @param
// @return                       bool         "是否登录成功"
func (w *WeiboClient) isloginWeiboCom() bool {
	resp, err := w.client.Get("https://weibo.com/aj/account/watermark")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	return ctype == "application/json; charset=utf-8"
	// application/json 一般登录成功，text/html登录失败
}

// @title         loginWeiboCom
// @description   使用单点登录机制登录.weibo.com，网页版一般用www.weibo.com
// @auth          星辰
// @param
// @return                       bool         "是否登录成功"
func (w *WeiboClient) loginWeiboCom() bool {
	if w.isloginWeiboCom() {
		return true
	}
	req, err := http.NewRequest("GET", "https://login.sina.com.cn/sso/login.php?url=https%3A%2F%2Fweibo.com%2F6356549855%2Fprofile&_rand=1607840168.9484&gateway=1&service=miniblog&entry=miniblog&useticket=1&returntype=META&_client_version=0.6.36", nil)
	req.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 Chrome/63.0.3239.108")
	resp, err := w.client.Do(req)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return false
	}
	resp.Body.Close()
	html_text := bytes2str(body)
	reg := regexp.MustCompile("location.replace\\(\\\"(https.*?)\\\"\\)")
	turl := reg.FindStringSubmatch(html_text)
	if len(turl) < 2 || turl[1] == "" {
		return false
	}
	resp, err = w.client.Get(turl[1])
	if err != nil {
		return false
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return false
	}
	resp.Body.Close()
	html_text = bytes2str(body)
	reg = regexp.MustCompile("setCrossDomainUrlList\\((.*?)\\);")
	jsonStr := reg.FindStringSubmatch(html_text)
	if len(jsonStr) < 2 || jsonStr[1] == "" {
		return false
	}
	type urllist struct{
		Retcode int `json:"retcode"`
		ArrURL []string `json:"arrURL"`
	}
	var jsonObj urllist
	err = json.Unmarshal(str2bytes(jsonStr[1]), &jsonObj)
	if err != nil {
		return false
	}
	resp, err = w.client.Get(
		jsonObj.ArrURL[0] +
		 "&callback=" + "sinaSSOController.doCrossDomainCallBack" +
		 "&scriptId=" + "ssoscript0" +
		 "&client=" + "ssologin.js(v1.4.2)",
		)
	if err != nil {
		return false
	}
	resp.Body.Close()
	if hasCookie(resp.Cookies(), "SUB") && w.isloginWeiboCom() {
		return true
	}
	return false
}

// @title         hasCookie
// @description   判断cookie数组里是否包含每个名称的cookie
// @auth          星辰
// @param
// @return                       bool         "是否包含cookie"
func hasCookie(cookies []*http.Cookie, name string) bool {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return true
		}
	}
	return false
}

// @title         bytes2str
// @description   字节集转字符串(非安全的指针转换)
// @auth          星辰
// @param         b        []byte     "用于转换的字节数组"
// @return                 string     "转换后的字符串"
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// @title         str2bytes
// @description   字符串转字节集(非安全的指针转换)
// @auth          星辰
// @param         s        string     "用于转换的字符串"
// @return                 []byte     "转换后的字节数组"
func str2bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
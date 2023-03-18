package gHttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type YxHttp struct {
	client   *http.Client
	response *http.Response
}
type YxHttpObj struct {
	Url     string
	Data    []byte
	Heads   map[string]string
	Cookies []*http.Cookie
	Proxy   string
}

// 创建客户端
func (h *YxHttp) cread(Proxy string) *http.Client {
	if len(Proxy) == 0 {
		return &http.Client{}
	} else {
		urlI := url.URL{}
		urlProxy, _ := urlI.Parse(Proxy)
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(urlProxy),
			}}
	}
}

// 设置协议头
func (h *YxHttp) setHead(req *http.Request, heads map[string]string) {
	for k, v := range heads {
		req.Header.Set(k, v)
	}
}
func (h *YxHttp) setCookie(req *http.Request, Cookie []*http.Cookie) {
	for _, v := range Cookie {
		req.AddCookie(v)
	}
	//cookie := http.Cookie{
	//	Name:     "_cookie",
	//	Value:    session.Uuid,
	//	HttpOnly: true,
	//	Path: "/",
	//}
}

// GetBody 获取响应文本
func (h *YxHttp) GetBody() ([]byte, error) {
	body, err := ioutil.ReadAll(h.response.Body)
	return body, err
}
func (h *YxHttp) Cookie2Map(arr []*http.Cookie) map[string]string {
	var cookies = make(map[string]string)
	for _, v := range arr {
		cookies[v.Name] = v.Value
	}
	return cookies
}

// GetAllCookieAarr 获取所有Cookie
func (h *YxHttp) GetAllCookieAarr() []*http.Cookie {
	return h.response.Cookies()
}

// GetAllCookie 获取所有Cookie
func (h *YxHttp) GetAllCookie() string {
	return h.GetHeads("set-cookie")
}

// GetHeads 获取指定协议头
func (h *YxHttp) GetHeads(name string) string {
	return h.response.Header.Get(name)
}

// Close 关闭请求
func (h *YxHttp) Close() {
	h.response.Body.Close()
}

// Send 发送请求
func (h *YxHttp) Send(method string, obj *YxHttpObj) error {
	//创建请求
	request, err := http.NewRequest(method, obj.Url, bytes.NewBuffer(obj.Data))
	if err != nil {
		return err
	}
	h.setHead(request, obj.Heads)          //设置协议头
	h.setCookie(request, obj.Cookies)      //设置Cookie
	h.client = h.cread(obj.Proxy)          //创建客户端
	h.response, err = h.client.Do(request) //发送请求
	return err
}

//cookiejar

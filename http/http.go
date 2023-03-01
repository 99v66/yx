package gHttp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func POST(url, data string) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

// 创建请求
func _creadClient(Proxy string) *http.Client {
	if len(Proxy) == 0 {
		return &http.Client{}
	}
	urlI := url.URL{}
	urlProxy, _ := urlI.Parse(Proxy)
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlProxy),
		}}
	return &client
}

// 设置协议头
func _setHead(req *http.Request, heads map[string]string) {
	for k, v := range heads {
		req.Header.Set(k, v)
	}
}

func POSTV2(url string, data []byte, heads map[string]string, Proxy string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return []byte(nil), err
	}
	_setHead(req, heads)
	client := _creadClient(Proxy)
	resp, err := client.Do(req)
	if err != nil {
		return []byte(nil), err
	}
	defer resp.Body.Close() //关闭请求
	//读取返回包体
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
func GETV2(url string, data []byte, heads map[string]string, Proxy string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(data))
	if err != nil {
		return []byte(nil), err
	}
	_setHead(req, heads)
	client := _creadClient(Proxy)
	resp, err := client.Do(req)
	if err != nil {
		return []byte(nil), err
	}
	defer resp.Body.Close() //关闭请求
	//读取返回包体
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
func Test() {
	s, e := POST("https://oapi.dingtalk.com/robot/send?access_token=91457c1c4fa1bad1f03f951382558e814f5f729f7d87b5f3741ebdaff2cb96b2", "aaaa")
	fmt.Println(s)
	fmt.Println(e)
}

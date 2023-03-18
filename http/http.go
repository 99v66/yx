package gHttp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	glog "yxProject/log"
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
	//s, e := POST("https://oapi.dingtalk.com/robot/send?access_token=91457c1c4fa1bad1f03f951382558e814f5f729f7d87b5f3741ebdaff2cb96b2", "aaaa")
	//fmt.Println(s)
	//fmt.Println(e)
	hp := &YxHttp{}
	o := &YxHttpObj{Url: "https://twitter.com/i/lists/1449337936890191872", Proxy: "http://127.0.0.1:7890"}
	err := hp.Send("get", o)
	if err != nil {
		glog.Log().Info().Err(err).Msg("创建请求失败")
		return
	}
	byte, err := hp.GetBody()
	glog.Log().Info().Bytes("响应数据", byte).Err(err).Send()
	glog.Log().Info().Str("Cookie", hp.GetAllCookie()).Send()
	glog.Log().Info().Str("Heads", hp.GetHeads("x-connection-hash")).Send()
	list := hp.GetAllCookieAarr()
	glog.Log().Info().Any("GetAllCookieAarr", list).Send()
	glog.Log().Info().Any("GetAllCookieAarr", hp.Cookie2Map(list)).Send()
	glog.Log().Info().Str("gt", hp.Cookie2Map(list)["gt"]).Send()
	hp.Close()
}

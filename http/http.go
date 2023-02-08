package gHttp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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

func Test() {
	s, e := POST("https://oapi.dingtalk.com/robot/send?access_token=91457c1c4fa1bad1f03f951382558e814f5f729f7d87b5f3741ebdaff2cb96b2", "aaaa")
	fmt.Println(s)
	fmt.Println(e)
}

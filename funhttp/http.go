package funhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var Client http.Client = clients()

// Res 模拟响应结构
// @Description:
type Res struct {
	Have string `json:"Have"`
}

// Get 获取GET请求
// @Description:
// @param uri
// @param args
// @return *http.Request
func Get(uri string, args map[string]interface{}) *http.Request {
	uri = uri + "?" + ToValues(args)
	request, _ := http.NewRequest("GET", uri, nil)
	return request
}

// PostForm POST接口form表单
// @Description:
// @param path
// @param args
// @return *http.Request
func PostForm(path string, args map[string]interface{}) *http.Request {
	request, _ := http.NewRequest("POST", path, strings.NewReader(ToValues(args)))
	return request
}

// PostJson POST请求,JSON参数
// @Description:
// @param path
// @param args
// @return *http.Request
func PostJson(path string, args map[string]interface{}) *http.Request {
	marshal, _ := json.Marshal(args)
	request, _ := http.NewRequest("POST", path, bytes.NewReader(marshal))

	return request
}

// ToValues 将map解析成HTTP参数,用于GET和POST form表单
// @Description:
// @param args
// @return string
func ToValues(args map[string]interface{}) string {
	if args != nil && len(args) > 0 {
		params := url.Values{}
		for k, v := range args {
			params.Set(k, fmt.Sprintf("%v", v))
		}
		return params.Encode()
	}
	return ""
}

// Response 获取响应详情,默认[]byte格式
// @Description:
// @param request
// @return []byte
func Response(request *http.Request) []byte {
	res, err := Client.Do(request)
	if err != nil {
		log.Println("响应出错", err)
		return nil
	}
	body, _ := ioutil.ReadAll(res.Body) // 读取响应 body, 返回为 []byte
	defer res.Body.Close()
	return body
}

// clients 初始化请求客户端
// @Description:
// @return http.Client
func clients() http.Client {
	return http.Client{
		Timeout: time.Duration(5) * time.Second, //超时时间
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   5,   //单个路由最大空闲连接数
			MaxConnsPerHost:       100, //单个路由最大连接数
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

// ParseRes 解析响应
// @Description:
// @receiver r
// @param res
func (r *Res) ParseRes(res []byte) {
	json.Unmarshal(res, r)
}

// ParseRes 解析响应,将[]byte转成传入对象
// @Description:
// @param res
// @param r
//
func ParseRes(res []byte, r interface{}) {
	json.Unmarshal(res, r)
}

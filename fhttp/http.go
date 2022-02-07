package fhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"funtester/futil"
	"io/ioutil"
	"log"
	"net"
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
// @return *fhttp.Request
func Get(uri string, args map[string]interface{}) *http.Request {
	if args != nil {
		uri = uri + "?" + ToValues(args)
	}
	request, _ := http.NewRequest("GET", uri, nil)
	return request
}

// PostForm POST接口form表单
// @Description:
// @param path
// @param args
// @return *fhttp.Request
func PostForm(path string, args map[string]interface{}) *http.Request {
	request, _ := http.NewRequest("POST", path, strings.NewReader(ToValues(args)))
	return request
}

// PostJson POST请求,JSON参数
// @Description:
// @param path
// @param args
// @return *fhttp.Request
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
			params.Set(k, futil.ToString(v))
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
// @return fhttp.Client
func clients() http.Client {
	dialer := &net.Dialer{
		Timeout: 1 * time.Second,
	}
	dialer.Resolver = &net.Resolver{
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.DialContext(ctx, "tcp", "114.114.114.114:53") // 通过tcp请求nameserver解析域名
		},
	}

	return http.Client{
		Timeout: time.Duration(5) * time.Second, //超时时间
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   200,   //单个路由最大空闲连接数
			MaxConnsPerHost:       10000, //单个路由最大连接数
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			//DialContext:           dialer.DialContext,
			DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
				host, port, err := net.SplitHostPort(address)
				if err != nil {
					return nil, err
				}
				//通过自定义nameserver获取域名解析的IP
				//ips, _ := dialer.Resolver.LookupHost(ctx, host)
				//for _, s := range ips {
				//	log.Println(s)
				//}

				// 创建链接
				if host == "fun.tester" {
					ips := []string{"127.0.0.1", "0.0.0.0"}
					ip := futil.RandomStrs(ips)
					log.Println(ip)
					conn, err := dialer.DialContext(ctx, network, ip+":"+port)
					if err == nil {
						return conn, nil
					}
				}
				return dialer.DialContext(ctx, network, address)
			},
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

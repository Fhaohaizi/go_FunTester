package task

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

var FastClient fasthttp.Client = fastClient()

// FastGet 获取GET请求对象,没有进行资源回收
// @Description:
// @param url
// @param args
// @return *fasthttp.Request
func FastGet(url string, args map[string]interface{}) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	values := ToValues(args)
	req.SetRequestURI(url + "?" + values)
	return req
}

// FastPostJson POST请求JSON参数,没有进行资源回收
// @Description:
// @param url
// @param args
// @return *fasthttp.Request
func FastPostJson(url string, args map[string]interface{}) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	marshal, _ := json.Marshal(args)
	req.SetBody(marshal)
	return req
}

// FastPostForm POST请求表单传参,没有进行资源回收
// @Description:
// @param url
// @param args
// @return *fasthttp.Request
func FastPostForm(url string, args map[string]interface{}) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	// 默认是application/x-www-form-urlencoded
	//req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	marshal, _ := json.Marshal(args)
	req.BodyWriter().Write([]byte(ToValues(args)))
	req.BodyWriter().Write(marshal)
	return req
}

// FastResponse 获取响应,保证资源回收
// @Description:
// @param request
// @return []byte
// @return error
func FastResponse(request *fasthttp.Request) ([]byte, error) {
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)
	defer fasthttp.ReleaseRequest(request)
	if err := FastClient.Do(request, response); err != nil {
		log.Println("响应出错了")
		log.Println(err.Error())
		return nil, err
	}
	return response.Body(), nil
}

// DoGet 发送GET请求,获取响应
// @Description:
// @param url
// @param args
// @return []byte
// @return error
func DoGet(url string, args map[string]interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	req.Header.SetMethod("GET")
	values := ToValues(args)
	req.SetRequestURI(url + "?" + values)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源
	if err := FastClient.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return nil, err
	}
	return resp.Body(), nil
}

// DoPosJson 发送POST请求JSON参数,获取响应
// @Description:
// @param url
// @param args
// @return []byte
// @return error
func DoPosJson(url string, args map[string]interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	marshal, _ := json.Marshal(args)
	req.SetBody(marshal)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源
	if err := FastClient.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return nil, err
	}
	return resp.Body(), nil
}

// DoPosForm 发送POST请求form参数,获取响应
// @Description:
// @param url
// @param args
// @return []byte
// @return error
func DoPosForm(url string, args map[string]interface{}) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	// 默认是application/x-www-form-urlencoded
	//req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	marshal, _ := json.Marshal(args)
	req.BodyWriter().Write([]byte(ToValues(args)))
	req.BodyWriter().Write(marshal)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := FastClient.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return nil, err
	}
	return resp.Body(), nil
}

// fastClient 获取fast客户端
// @Description:
// @return fasthttp.Client
func fastClient() fasthttp.Client {
	return fasthttp.Client{
		Name:                     "FunTester",
		NoDefaultUserAgentHeader: true,
		TLSConfig:                &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost:          2000,
		MaxIdleConnDuration:      5 * time.Second,
		MaxConnDuration:          5 * time.Second,
		ReadTimeout:              5 * time.Second,
		WriteTimeout:             5 * time.Second,
		MaxConnWaitTimeout:       5 * time.Second,
	}
}

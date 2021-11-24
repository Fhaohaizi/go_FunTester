package test

import (
	"funtester/ft"
	"funtester/funhttp"
	"log"
	"testing"
)

const testurl = "http://localhost:12345/test"

func args() map[string]interface{} {
	return map[string]interface{}{
		"code": 32,
		"fun":  32,
		"msg":  "324",
	}
}

func TestGet(t *testing.T) {
	get := funhttp.FastGet(testurl, args())
	res, err := funhttp.FastResponse(get)
	if err != nil {
		t.Fail()
	}
	v := string(res)
	log.Println(v)
	if v != "get请求" {
		t.Fail()
	}
}

func TestPostJson(t *testing.T) {
	post := funhttp.FastPostJson(testurl, args())
	res, err := funhttp.FastResponse(post)
	if err != nil {
		t.Fail()
	}
	v := string(res)
	log.Println(v)
	if v != "post请求json表单" {
		t.Fail()
	}
}

func TestPostForm(t *testing.T) {
	post := funhttp.FastPostForm(testurl, args())
	res, err := funhttp.FastResponse(post)
	if err != nil {
		t.Fail()
	}
	v := string(res)
	log.Println(v)
	if v != "post请求form表单" {
		t.Fail()
	}
}

func TestGetNor(t *testing.T) {
	res, err := ft.FastGet(testurl, args())
	if err != nil {
		t.Fail()
	}
	v := string(res)
	log.Println(v)
	if v != "get请求" {
		t.Fail()
	}
}

func TestPostJsonNor(t *testing.T) {
	res, err := ft.FastPostJson(testurl, args())
	if err != nil {
		t.Fail()
	}
	v := string(res)
	log.Println(v)
	if v != "post请求json表单" {
		t.Fail()
	}
}

func TestPostFormNor(t *testing.T) {
	res, err := ft.FastPostForm(testurl, args())
	if err != nil {
		t.Fail()
	}
	v := string(res)
	log.Println(v)
	if v != "post请求form表单" {
		t.Fail()
	}
}

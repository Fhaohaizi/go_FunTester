package fwebsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	websocket2 "golang.org/x/net/websocket"
	"log"
	"testing"
)

// TestWebSocket
// @Description: 测试WebSocket脚本
// @param t
func TestWebSocket() {
	url := "ws://localhost:1234/fwebsocket"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	log.Printf("响应:%s", fmt.Sprint(res))
	defer c.Close()
	done := make(chan struct{})
	err = c.WriteMessage(websocket.TextMessage, []byte("你好,我是FunTester"))
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
		log.Printf("收到消息: %s", message)

	}
	<-done
}

// TestWebSocket2
// @Description: 第二种测试WebSocket的方法
// @param t
func TestWebSocket2(t *testing.T) {
	url := "ws://localhost:1234/fwebsocket"
	dial, er := websocket2.Dial(url, "", "/fwebsocket")
	if er != nil {
		fmt.Println(er)
		return
	}
	err := websocket2.Message.Send(dial, "你好,我是FunTester - Go ,Have Fun ~ Tester ！")

	if err != nil {
		fmt.Println(err)
	}
	ints := make(chan int)
	go func() {
		defer close(ints)
		for {
			var m []byte
			err2 := websocket2.Message.Receive(dial, &m)
			if err2 == nil {
				fmt.Println(string(m))
			}
		}

	}()
	<-ints
}

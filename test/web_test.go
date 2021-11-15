package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
	"time"
)

// TestWebSocket
// @Description: 测试WebSocket脚本
// @param t
func TestWebSocket(t *testing.T) {

	url := "ws://localhost:1234/websocket"
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

// TestWEBs 创建一个WebSocket服务
// @Description:
// @param t
func TestWEBs(t *testing.T) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
	}

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("%s receive: %s\n", conn.RemoteAddr(), string(msg))

			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":1234", nil)
}

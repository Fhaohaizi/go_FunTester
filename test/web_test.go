package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestWebSocket(t *testing.T) {

	url := "wss://wspri.coinall.ltd:8443/ws/v5/public"
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
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Fatal(err)
				break
			}
			log.Printf("收到消息: %s", message)

		}
	}()
	s := <-done
	fmt.Println(s)

}

func TestWEBs(t *testing.T) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":1234", nil)
}

package test

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestMn2(t *testing.T) {
	url := "ws://localhost:1234/websocket"
	dial, er := websocket.Dial(url, "", "/websocket")
	if er != nil {
		fmt.Println(er)
		return
	}
	err := websocket.Message.Send(dial, "你好,我是FunTester - Go ,Have Fun ~ Tester ！")

	if err != nil {
		fmt.Println(err)
	}
	ints := make(chan int)
	go func() {
		defer close(ints)
		for {
			var m []byte
			err2 := websocket.Message.Receive(dial, &m)
			if err2 == nil {
				fmt.Println(string(m))
			}
		}

	}()
	<-ints
}

// Echo
// @Description:WebSocket接口handle
// @param ws
func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}
		log.Printf("收到消息:%s", reply)
		msg := string(time.Now().String())
		websocket.Message.Send(ws, msg)
	}

}

// TestSer
// @Description: 创建一个WebSocket接口
// @param t
func TestSer(t *testing.T) {
	//接受websocket的路由地址
	http.HandleFunc("/websocket", func(w http.ResponseWriter, req *http.Request) {
		s := websocket.Server{Handler: websocket.Handler(Echo)}
		s.ServeHTTP(w, req)
	})
	if err := http.ListenAndServe(":1234", nil); err != nil {

		log.Fatal("ListenAndServe:", err)

	}
}

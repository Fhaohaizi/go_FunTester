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
	url := "wss://wspri.coinall.ltd:8443/ws/v5/public"
	dial, er := websocket.Dial(url, "", "/ws/v5/public")
	if er != nil {
		fmt.Println(er)
		fmt.Println(324)
		return
	}
	err := websocket.Message.Send(dial, "{\"op\":\"subscribe\",\"args\":[{\"channel\":\"candle1m\",\"instId\":\"LTC-USDT\"}]}")

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

var cons = make(map[int]*websocket.Conn)
var i int = 1

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}
		for k, con := range cons {
			sendMessage("你发错了", k, con)
		}
		log.Printf("收到消息:%s", reply)
		msg := string(time.Now().String())
		websocket.Message.Send(ws, msg)
	}

}

func sendMessage(msg string, k int, s *websocket.Conn) {
	if err := websocket.Message.Send(s, msg); err != nil {
		fmt.Println("Can't send")
		delete(cons, k)
	}

}

func TestSer(t *testing.T) {
	//接受websocket的路由地址
	http.Handle("/websocket", websocket.Handler(Echo))
	http.HandleFunc("/t", func(w http.ResponseWriter, req *http.Request) {
		s := websocket.Server{Handler: websocket.Handler(Echo)}
		s.ServeHTTP(w, req)
	})
	if err := http.ListenAndServe(":1234", nil); err != nil {

		log.Fatal("ListenAndServe:", err)

	}
}

package test

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"testing"
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

func Echo(ws *websocket.Conn) {

	var err error

	for {

		var reply string
		//websocket接受信息
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("receive failed:", err)
			break
		}

		fmt.Println("reveived from client: " + reply)
		msg := "received:" + reply
		fmt.Println("send to client:" + msg)
		//这里是发送消息
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("send failed:", err)
			break
		}

	}

}

func TestSer(t *testing.T) {
	//接受websocket的路由地址
	http.Handle("/websocket", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {

		log.Fatal("ListenAndServe:", err)

	}
}

package fwebsocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	websocket2 "golang.org/x/net/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

// CreateServer
// @Description: 重建一个WebSocket服务
// @param port 端口
// @param path 路径
func CreateServer(port int, path string) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
	}

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		conn.WriteMessage(websocket.TextMessage, []byte("msg"))

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%s receive: %s\n", conn.RemoteAddr(), string(msg))

			if err = conn.WriteMessage(msgType, msg); err != nil {
				log.Println("ffahv")
				return
			}
		}
	})

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func CreateServer2(port int, path string) {
	//接受websocket的路由地址
	http.HandleFunc("/"+path, func(w http.ResponseWriter, req *http.Request) {
		s := websocket2.Server{Handler: websocket2.Handler(func(conn *websocket2.Conn) {
			var err error
			for {
				var reply string
				if err = websocket2.Message.Receive(conn, &reply); err != nil {
					fmt.Println("receive failed:", err)
					break
				}
				log.Printf("收到消息:%s", reply)
				msg := string(time.Now().String())
				websocket2.Message.Send(conn, msg)
			}
		})}
		s.ServeHTTP(w, req)
	})
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {

		log.Fatal("ListenAndServe:", err)

	}
}

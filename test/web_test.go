package test

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
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
func connHandler(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	c.Write([]byte("你好,我是FunTester!"))
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			return
		}

		c.Write([]byte(input))

		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			continue
		}

		fmt.Print(string(buf[0:cnt]))
	}
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
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":1234", nil)
}

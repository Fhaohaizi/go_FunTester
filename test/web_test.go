package test

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"
)

func TestWebSocket(t *testing.T) {
	flag.Parse()

	// 用来接收命令行的终止信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := "wss://wspri.coinall.ltd:8443/ws/v5/public"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	log.Printf("响应:%s", fmt.Sprint(res))
	defer c.Close()
	done := make(chan struct{})
	err = c.WriteMessage(websocket.TextMessage, []byte("{\"op\":\"subscribe\",\"args\":[{\"channel\":\"candle1m\",\"instId\":\"LTC-USDT\"}]}"))
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		defer close(done)
		for {
			// 从接收服务端message
			_, message, _ := c.ReadMessage()
			log.Printf("收到消息: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	s := <-done
	fmt.Println(s)
	//for {
	//	select {
	//	case <-done:
	//		return
	//	//case <-ticker.C:
	//	//	log.Println("定时器")
	//	case <-interrupt:
	//		log.Println("中断")
	//		// 收到命令行终止信号，通过发送close message关闭连接。
	//		c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "FunTester"))
	//		// 收到接收协程完成的信号或者超时，退出
	//		select {
	//		case <-done:
	//		case <-time.After(time.Second):
	//		}
	//		return
	//	}
	//}
}
func connHandler(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	c.Write([]byte("{\"op\":\"subscribe\",\"args\":[{\"channel\":\"candle1m\",\"instId\":\"LTC-USDT\"}]}"))
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
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}

// TestF 测试外部输入
// @Description:
// @param t
func TestF(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println(input)
		if input == "quit" {
			return
		}

	}
}

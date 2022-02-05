package main

import (
	"fmt"
	"funtester/fhttp"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	os.Mkdir("./log/", 0766)
	//os.Mkdir("./long/", 0666)
	file := "./log/" + string(time.Now().Format("20060102")) + ".log"
	println(file)
	os.Create(file)
	openFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	writer := io.MultiWriter(os.Stdout, openFile)
	log.SetOutput(writer)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate)
}

var done = make(chan struct{})

func main() {
	url := "http://localhost:12345/test/fun"
	get := fhttp.Get(url, nil)
	response := fhttp.Response(get)
	log.Println(string(response))

	fastGet := fhttp.FastGet(url, nil)
	fastResponse, _ := fhttp.FastResponse(fastGet)
	log.Printf(string(fastResponse))
}

func ManySocket() {
	args := os.Args
	var n, t int = 1, 1
	if len(args) > 1 {
		n, _ = strconv.Atoi(args[1])
		log.Printf("创建 %d倍连接", n)
	}

	if len(args) > 2 {
		t, _ = strconv.Atoi(args[2])
		log.Printf("创建 %d倍连接", t)
	}
	for i := 0; i < n; i++ {
		for i := 0; i < t; i++ {
			time.Sleep(100 * time.Millisecond)
		}
		go getSoecket()
		log.Printf("创建第%d个链接", i+1)
	}
	<-done
	log.Println("结束了")
}

func getSoecket() {
	//url := "wss://wspre.okex.com:8443/ws/v5/public"
	url := "wss://wspri.coinall.ltd:8443/ws/v5/public"
	c, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	log.Printf("响应:%s", fmt.Sprint(res))
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, []byte("{\"op\":\"subscribe\",\"args\":[{\"channel\":\"candle1m\",\"instId\":\"LTC-USDT\"}]}"))
	if err != nil {
		fmt.Println(err)
	}
	err = c.WriteMessage(websocket.PingMessage, []byte("ping"))
	if err != nil {
		fmt.Println(err)
	}
	//go func() {
	//	defer close(done)
	for {
		err := c.WriteMessage(websocket.PingMessage, []byte("ping"))
		if err != nil {
			log.Println(err)
		}
		_, m, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
			//break
		}
		log.Printf("收到消息: %s", m)

	}
	//}()
	//<-done
}

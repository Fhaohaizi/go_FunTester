package funhttp

import (
	"bufio"
	r1 "crypto/rand"
	"funtester/base"
	r2 "math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func Intput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Random(data []byte) {
	r1.Read(data)
}

func RandomInt(bound int) int {
	r := r2.New(r2.NewSource(time.Now().UnixNano()))
	return r.Intn(bound)
}
func RandomFloat() float32 {
	r := r2.New(r2.NewSource(time.Now().UnixNano()))
	return r.Float32()
}

func Milli() int64 {
	return time.Now().UnixMilli()
}

func Nano() int64 {
	return time.Now().UnixNano()
}

func Workspace() string {
	getwd, err := os.Getwd()
	if err != nil {
		return getwd
	}
	return base.Empty
}

func Once(drive sync.Once, f func()) {
	drive.Do(f)
}

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

//ticker := time.NewTicker(time.Second)
//defer ticker.Stop()

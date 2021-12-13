package futil

import (
	"bufio"
	r1 "crypto/rand"
	r2 "math/rand"
	"os"
	"strings"
	"time"
)

var strs = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func Intput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func Random(data []byte) {
	r1.Read(data)
}

// RandomInt
// @Description:
// @param bound
// @return int 返回0~int-1
func RandomInt(bound int) int {
	r := r2.New(r2.NewSource(time.Now().UnixNano()))
	return r.Intn(bound)
}

func RandomStr(bound int) string {
	var build strings.Builder
	for i := 0; i < bound; i++ {
		build.WriteString(RandomStrs(strs))
	}
	return build.String()
}

func RandomSlice(s []interface{}) interface{} {
	randomInt := RandomInt(len(s))
	return s[randomInt]
}

func RandomStrs(s []string) string {
	randomInt := RandomInt(len(s))
	return s[randomInt]
}

func RandomInts(s []int) int {
	randomInt := RandomInt(len(s))
	return s[randomInt]
}

func RandomFloat() float32 {
	r := r2.New(r2.NewSource(time.Now().UnixNano()))
	return r.Float32()
}

func Workspace() string {
	getwd, _ := os.Getwd()
	return getwd
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
//		c.WriteMessage(fwebsocket.CloseMessage, fwebsocket.FormatCloseMessage(fwebsocket.CloseNormalClosure, "FunTester"))
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

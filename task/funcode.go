package task

import (
	"bufio"
	r1 "crypto/rand"
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
	return Empty
}

func Once(drive sync.Once, f func()) {
	drive.Do(f)
}

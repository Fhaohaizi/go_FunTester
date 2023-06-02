package ftool

import (
	"bufio"
	r1 "crypto/rand"
	"funtester/base"
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

// RangInt
// @Description: 通过范围
// @param start 起始值
// @param end 最大值,不会到达
// @return int  返回值
func RangInt(start, end int) int {
	if end <= start {
		return base.ErrorInt
	}
	return RandomInt(end-start) + start
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

func RandomArray(s []struct{}) struct{} {
	randomInt := RandomInt(len(s))
	return s[randomInt]
}

// RandomStrs
// @Description: 从String切片中随机一个
// @param s
// @return string
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

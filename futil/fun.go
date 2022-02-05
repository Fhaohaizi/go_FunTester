package futil

import (
	"fmt"
	"log"
	"strings"
)

// ToString
// @Description: 将对象转换成string
// @param v
// @return string
func ToString(v interface{}) string {
	return fmt.Sprint(v)
}

// PrintTime
// @Description: 打印方法执行时间
// @param f
// @param time
// @param name
func PrintTime(f func(), times int, name string) {
	start := Milli()
	f()
	end := Milli()
	log.Printf("%s执行%d次耗时:%s s", name, times, NumberFormat(ToString(float32(end-start)/1000)))

}

func NumberFormat(str string) string {
	length := len(str)
	if length < 4 {
		return str
	}
	arr := strings.Split(str, ".") //用小数点符号分割字符串,为数组接收
	length1 := len(arr[0])
	if length1 < 4 {
		return str
	}
	count := (length1 - 1) / 3
	for i := 0; i < count; i++ {
		arr[0] = arr[0][:length1-(i+1)*3] + "," + arr[0][length1-(i+1)*3:]
	}
	return strings.Join(arr, ".") //将一系列字符串连接为一个字符串，之间用sep来分隔。
}

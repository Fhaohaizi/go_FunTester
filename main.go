package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	os.Mkdir("./log/", 0666)
	os.Mkdir("./long/", 0666)
	file := "./log/" + string(time.Now().Format("20060102")) + ".log"
	openFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	writer := io.MultiWriter(os.Stdout, openFile)
	log.SetOutput(writer)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate)
}

func main() {
	now := time.Now()
	milli := now.UnixMilli()
	unix := now.Unix()
	log.Println(milli)
	log.Println(unix)
	log.Println(now.UnixNano())
	log.Println(now.UnixMicro())
	fmt.Println(strings.Index("123", "2"))
}

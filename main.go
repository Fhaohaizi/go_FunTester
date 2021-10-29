package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func init() {
	os.Mkdir("./log/", 0666)
	//os.Mkdir("./long/", 0666)
	file := "./log/" + string(time.Now().Format("20060102")) + ".log"
	openFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	writer := io.MultiWriter(os.Stdout, openFile)
	log.SetOutput(writer)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate)
}

func main() {
	args := os.Args
	fmt.Println(args)
	now := time.Now()
	milli := now.UnixMilli()
	unix := now.Unix()
	log.Println(milli)
	log.Println(unix)
	log.Println(now.UnixNano())
	log.Println(now.UnixMicro())
	fmt.Println(strings.Index("123", "1112"))
	getwd, _ := os.Getwd()
	fmt.Println(getwd)
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	var wait sync.WaitGroup
	wait.Add(10)
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
			wait.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	wait.Wait()
	//reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	//for {
	//	input, _ := reader.ReadString('\n')
	//	input = strings.TrimSpace(input)
	//	if input == "quit" {
	//		return
	//	}
	//}

}

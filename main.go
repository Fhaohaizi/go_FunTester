package main

import (
	"fmt"
	"funtester/task"
	"io"
	"log"
	"os"
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
	fmt.Println(task.RandomFloat())

}

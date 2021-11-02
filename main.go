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
	log.Printf("32")
	args := os.Args
	fmt.Println(args)
	fmt.Println(task.RandomFloat())
	ints := make(chan int)
	go ps(ints)
	go p(ints)
	time.Sleep(2 * time.Second)
}

func p(c chan int) {
	for {
		log.Printf("324")
		i := <-c
		fmt.Println(i)
		close(c)
		break
	}
}
func ps(c chan<- int) {
	for {
		c <- 23333333
		break
	}
}

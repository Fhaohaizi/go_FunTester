package test

import (
	"funtester/execute"
	"log"
	"sync"
	"testing"
	"time"
)

var tasks = make(chan func(), 1000)
var lock sync.Once

// 这里是 worker，我们将并发执行多个 worker。
// worker 将从 `jobs` 通道接收任务，并且通过 `results` 发送对应的结果。
// 我们将让每个任务间隔 1s 来模仿一个耗时的任务。
func worker() {
	for {
		f := <-tasks
		f()
	}
}

func TestPool1(t *testing.T) {
	//StartPool()
	Execute(func() {
		log.Println(3232)
	})
	ch := make(chan int, 100)
	ch <- 32
	//close(ch)
	ticker := time.NewTicker(5 * time.Second)
	//timer := time.NewTimer(time.Second)
	for {
		select {
		case msg := <-ch:
			log.Println(msg)
			//return
		case <-ticker.C:
			log.Println("定时器")
			//case <-interrupt:
			//	log.Println("中断")
			//select {
			//case <-done:
			//case <-time.After(time.Second):
			//}
			//return
			//}

		}
		//msg, ok := <-ch
		//fmt.Println(msg)
		//fmt.Println(ok)
		//ftool.Sleep(1000)
	}
	time.Sleep(time.Second)
}

func StartPool() {
	for i := 0; i < 10; i++ {
		go worker()
	}
}

func Execute(task func()) {
	execute.Once(lock, StartPool)
	tasks <- task
}

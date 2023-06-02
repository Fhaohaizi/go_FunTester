package pool

import (
	"fmt"
	"funtester/execute"
	"funtester/ftool"
	"log"
	"sync"
	"testing"
	"time"
)

func TestPool12(t *testing.T) {

	// 为了使用 worker 线程池并且收集他们的结果，我们需要 2 个通道。
	jobs := make(chan func() int, 100)
	results := make(chan int, 100)

	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务。
	for w := 1; w <= 3; w++ {
		go worker1(w, jobs, results)
	}

	// 这里我们发送 9 个 `jobs`，然后 `close` 这些通道
	// 来表示这些就是所有的任务了。
	for j := 1; j <= 9; j++ {
		log.Println(j)
		jobs <- func() int {
			return j
		}
	}
	//close(jobs)

	// 最后，我们收集所有这些任务的返回值。
	for a := 1; a <= 9; a++ {
		<-results
	}

}
func TestPool122(t *testing.T) {
	pool := execute.GetPool(6, 2, 200, 1)
	for i := 0; i < 15; i++ {
		err := pool.Execute(func() {
			ftool.Sleep(1000)
			log.Println(32432)

		})
		if err != nil {
			log.Println(err.Error())
		}
	}
	ftool.Sleep(1000)
	pool.Wait()
}

// 这里是 worker，我们将并发执行多个 worker。
// worker 将从 `jobs` 通道接收任务，并且通过 `results` 发送对应的结果。
// 我们将让每个任务间隔 1s 来模仿一个耗时的任务。
func worker1(id int, jobs <-chan func() int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j())
		time.Sleep(time.Second)
		results <- j()
	}
}

var poolLock sync.Mutex = sync.Mutex{}

func do() {
	poolLock.Lock()
	defer poolLock.Unlock()
	log.Println("解锁成功!")
}

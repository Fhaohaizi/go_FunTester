package execute

import (
	"log"
	"sync"
	"time"
)

func Execute(fun func(), times, thread int) {
	c := make(chan int) //确认所有线程都结束
	key := false        //用于控制所有线程一起结束
	start := time.Now().UnixMilli()
	for i := 0; i < thread; i++ {
		go func() {
			sum := 0
			for i := 0; i < times; i++ {
				if key {
					break
				}
				fun()
				sum++
			}
			key = true
			c <- sum
		}()
	}
	total := 0
	for i := 0; i < thread; i++ {
		num := <-c
		total += num
	}
	end := time.Now().UnixMilli()
	diff := end - start
	//total := thread * times
	log.Printf("总耗时: %f", float64(diff)/1000)

	log.Printf("请求总数: %d", total)
	log.Printf("QPS: %f", float64(total)/float64(diff)*1000.0)
}

func ExecuteTask(task Task) {
	c := make(chan int) //确认所有线程都结束
	key := false        //用于控制所有线程一起结束
	start := time.Now().UnixMilli()
	for i := 0; i < task.Thread; i++ {
		go func() {
			sum := 0
			for i := 0; i < task.Times; i++ {
				if key {
					break
				}
				task.Run()
				sum++
			}
			key = true
			c <- sum
		}()
	}
	total := 0
	for i := 0; i < task.Thread; i++ {
		num := <-c
		total += num
	}
	end := time.Now().UnixMilli()
	diff := end - start
	//total := thread * times
	log.Printf("总耗时: %f", float64(diff)/1000)

	log.Printf("请求总数: %d", total)
	log.Printf("QPS: %f", float64(total)/float64(diff)*1000.0)
}

func Once(drive sync.Once, f func()) {
	drive.Do(f)
}

type Task struct {
	Times, Thread int
}

func (t *Task) Run() {
	log.Println("最初的梦想,需要实现!")
}

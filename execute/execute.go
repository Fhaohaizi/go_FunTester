package execute

import (
	"funtester/ftool"
	"log"
	"sync"
	t "time"
)

// ExecuteRoutineTimes
// @Description: FunTester性能测试执行框架
// @param fun 待执行方法
// @param times 次数
// @param routine 线程数
func ExecuteRoutineTimes(fun func(), times, routine int) {
	c := make(chan int) //确认所有线程都结束
	key := false        //用于控制所有线程一起结束
	start := ftool.Milli()
	for i := 0; i < routine; i++ {
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
	for i := 0; i < routine; i++ {
		num := <-c
		total += num
	}
	end := ftool.Milli()
	diff := end - start
	//total := thread * times
	log.Printf("总耗时: %f", float64(diff)/1000)

	log.Printf("请求总数: %d", total)
	log.Printf("QPS: %f", float64(total)/float64(diff)*1000.0)
}

func ExecuteRoutineTime(fun func(), time, routine int) {
	c := make(chan int) //确认所有线程都结束
	key := false        //用于控制所有线程一起结束
	start := t.Now().UnixMilli()
	go func() {
		ftool.Sleep(time * 1000)
		key = true
	}()
	for i := 0; i < routine; i++ {
		go func() {
			sum := 0
			for {
				if key {
					break
				}
				fun()
				sum++
			}
			c <- sum
		}()
	}
	total := 0
	for i := 0; i < routine; i++ {
		num := <-c
		total += num
	}
	end := t.Now().UnixMilli()
	diff := end - start
	//total := thread * time
	log.Printf("总耗时: %f", float64(diff)/1000)

	log.Printf("请求总数: %d", total)
	log.Printf("QPS: %f", float64(total)/float64(diff)*1000.0)
}

func ExecuteTask(task Task) {
	c := make(chan int) //确认所有线程都结束
	key := false        //用于控制所有线程一起结束
	start := t.Now().UnixMilli()
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
	end := t.Now().UnixMilli()
	diff := end - start
	//total := thread * times
	log.Printf("总耗时: %f", float64(diff)/1000)

	log.Printf("请求总数: %d", total)
	log.Printf("QPS: %f", float64(total)/float64(diff)*1000.0)
}

// ExecuteQPS 以固定QPS执行某方法
//  @Description:
//  @param task
//  @param qps
//
func ExecuteQPS(task func(), qps int) {
	for {
		ftool.Sleep(1000)
		go func() {
			for i := 0; i < qps; i++ {
				task()
			}
		}()
	}
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

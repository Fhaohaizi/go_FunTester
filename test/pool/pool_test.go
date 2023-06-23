package pool

import (
	"funtester/execute"
	"funtester/ftool"
	"log"
	"sync"
	"testing"
)

func TestPool122(t *testing.T) {
	pool := execute.GetPool(1000, 2, 200, 1)
	var in string
	go ftool.HandleInput(func(input string) bool {
		for {
			in = input
		}
	})
	for {
		pool.ExecuteQps(func() {
			log.Println(in)
		}, 4)
		ftool.Sleep(1000)
	}
	pool.Wait()
}

func TestPool(t *testing.T) {
	pool := execute.GetPool(1000, 1, 200, 1)
	for i := 0; i < 3; i++ {
		pool.Execute(func() {
			log.Println(i)
			ftool.Sleep(1000)
		})
	}
	ftool.Sleep(3000)
	pool.Wait()
	log.Printf("T : %d", pool.ExecuteTotal)
	log.Printf("R : %d", pool.ReceiveTotal)
	log.Printf("max : %d", pool.Max)
	log.Printf("min : %d", pool.Min)
}

var poolLock sync.Mutex

func TestDa(t *testing.T) {
	poolLock.Lock()
	defer poolLock.Unlock()
	log.Println("解锁成功!")
}

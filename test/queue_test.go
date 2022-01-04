package test

import (
	"funtester/fhttp"
	"funtester/futil"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	url       = "http://localhost:12345/funtester?name=3242423&age=416516515"
	token     = "FunTesterFunTesterFunTesterFunTesterFunTesterFunTesterFunTester"
	total     = 100_0000
	size      = 10
	threadNum = 100
	piece     = total / size
)

func TestQueue(t *testing.T) {
	var index int32 = 0
	rs := make(chan *http.Request, total+10000)
	var group sync.WaitGroup
	group.Add(threadNum)
	milli := futil.Milli()
	funtester := func() {
		go func() {
			for {
				l := atomic.AddInt32(&index, 1)
				if l%piece == 0 {
					m := futil.Milli()
					log.Println(m - milli)
					milli = m
				}
				if l > total {
					break
				}
				get := fhttp.Get(url, nil)
				get.Header.Add("token", token)
				get.Header.Add("Connection", "keep-alive")
				rs <- get
			}
			group.Done()
		}()
	}
	start := futil.Milli()
	for i := 0; i < threadNum; i++ {
		funtester()
	}
	group.Wait()
	end := futil.Milli()

	log.Println(atomic.LoadInt32(&index))
	log.Printf("平均每毫秒速率%d", total/(end-start))
}

func TestConsumer(t *testing.T) {
	rs := make(chan *http.Request, total+10000)
	var group sync.WaitGroup
	group.Add(10)
	funtester := func() {
		go func() {
			for {
				if len(rs) > total {
					break
				}
				get := fhttp.Get(url, nil)
				get.Header.Add("token", token)
				get.Header.Add("Connection", "keep-alive")
				rs <- get
			}
			group.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		funtester()
	}
	group.Wait()
	log.Printf("造数据完成! 总数%d", len(rs))
	totalActual := int64(len(rs))
	var conwait sync.WaitGroup
	conwait.Add(threadNum)
	consumer := func() {
		go func() {
		FUN:
			for {
				select {
				case <-rs:
				case <-time.After(100 * time.Millisecond):
					break FUN
				}
			}
			conwait.Done()
		}()
	}
	start := futil.Milli()
	for i := 0; i < threadNum; i++ {
		consumer()
	}
	conwait.Wait()
	end := futil.Milli()
	log.Printf("平均每毫秒速率%d", totalActual/(end-start))

}

func TestBoth(t *testing.T) {
	var index int32 = 0
	rs := make(chan *http.Request, total+10000)
	funtester := func() {
		go func() {
			for {
				l := atomic.AddInt32(&index, 1)
				if l > total {
					break
				}
				get := fhttp.Get(url, nil)
				get.Header.Add("token", token)
				get.Header.Add("Connection", "keep-alive")
				rs <- get
			}
		}()
	}
	var conwait sync.WaitGroup
	conwait.Add(threadNum)
	consumer := func() {
		go func() {
		FUN:
			for {
				select {
				case <-rs:
				case <-time.After(100 * time.Millisecond):
					break FUN
				}
			}
			conwait.Done()
		}()
	}
	start := futil.Milli()
	for i := 0; i < threadNum; i++ {
		consumer()
		funtester()
	}
	conwait.Wait()
	end := futil.Milli()
	log.Printf("平均每毫秒速率%d", int64(index)/(end-start))

}

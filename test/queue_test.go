package test

import (
	"funtester/base"
	"funtester/execute"
	"funtester/ftool"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	url       = "http://localhost:12345/funtester"
	token     = "FunTesterFunTesterFunTesterFunTesterFunTesterFunTesterFunTester"
	total     = 100_0000
	length    = 20_0000
	size      = 10
	threadNum = 1
	piece     = total / size
)

func TestQueue(t *testing.T) {
	var index int32 = 0
	rs := make(chan *http.Request, total+10000)
	var group sync.WaitGroup
	group.Add(threadNum)
	milli := ftool.Milli()
	funtester := func() {
		go func() {
			for {
				l := atomic.AddInt32(&index, 1)
				if l%piece == 0 {
					m := ftool.Milli()
					log.Println(m - milli)
					milli = m
				}
				if l > total {
					break
				}
				get := getRequest()
				rs <- get
			}
			group.Done()
		}()
	}
	start := ftool.Milli()
	for i := 0; i < threadNum; i++ {
		funtester()
	}
	group.Wait()
	end := ftool.Milli()

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
				get := getRequest()

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
				case <-time.After(10 * time.Millisecond):
					break FUN
				}
			}
			conwait.Done()
		}()
	}
	start := ftool.Milli()
	for i := 0; i < threadNum; i++ {
		consumer()
	}
	conwait.Wait()
	end := ftool.Milli()
	log.Printf("平均每毫秒速率%d", totalActual/(end-start))

}

func TestBoth(t *testing.T) {
	var index int32 = 0
	rs := make(chan *http.Request, length)
	for i := 0; i < length; i++ {
		rs <- getRequest()
	}
	funtester := func() {
		go func() {
			for {
				l := atomic.AddInt32(&index, 1)
				if l > total {
					break
				}
				get := getRequest()

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
				case <-time.After(10 * time.Millisecond):
					break FUN
				}
			}
			conwait.Done()
		}()
	}
	start := ftool.Milli()
	for i := 0; i < threadNum; i++ {
		consumer()
		funtester()
	}
	conwait.Wait()
	end := ftool.Milli()
	log.Printf("平均每毫秒速率:%d", int64(index+length)/(end-start))

}

// TestBase
// @Description: 基准测试
// @param t
func TestBase(t *testing.T) {
	execute.ExecuteRoutineTimes(func() {
		getFastRequest()
	}, total, threadNum*5)
}

func getRequest() *http.Request {
	//get, _ := http.NewRequest("GET", base.Empty, nil)

	//get,_ := http.NewRequest("GET",url, nil)
	//get.Header.Add("token", token)
	//get.Header.Add("Connection", base.ConnectionAlive)
	//get.Header.Add("User-Agent", base.UserAgent)

	get, _ := http.NewRequest("GET", url, nil)
	get.Header.Add("token", token)
	get.Header.Add("token1", token)
	get.Header.Add("token2", token)
	get.Header.Add("token3", token)
	get.Header.Add("token4", token)
	get.Header.Add("token5", token)
	get.Header.Add("Connection", base.ConnectionAlive)
	get.Header.Add("User-Agent", base.UserAgent)

	return get
}

func getFastRequest() *fasthttp.Request {
	get := fasthttp.AcquireRequest()
	get.Header.SetMethod("GET")
	//get.SetRequestURI(base.Empty)
	get.SetRequestURI(url)
	//get.Header.Add("token", token)
	//get.Header.Add("Connection", base.ConnectionAlive)
	//get.Header.Add("User-Agent", base.UserAgent)

	get.Header.Add("token", token)
	get.Header.Add("token1", token)
	get.Header.Add("token2", token)
	get.Header.Add("token3", token)
	get.Header.Add("token4", token)
	get.Header.Add("token5", token)
	get.Header.Add("Connection", base.ConnectionAlive)
	get.Header.Add("User-Agent", base.UserAgent)

	return get
}

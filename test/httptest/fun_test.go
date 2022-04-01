package httptest

import (
	"funtester/fhttp"
	"github.com/go-resty/resty/v2"
	"log"
	"testing"
	"time"
)

var key bool = false
var r = resty.New()

const (
	url      = "http://localhost:12345/test/fun"
	thread   = 1
	times    = 200000
	workers  = thread
	maxCount = times
)

func TestPer(t *testing.T) {
	//get := fhttp.Get(url, nil)
	c := make(chan int)

	start := time.Now().UnixMilli()
	for i := 0; i < thread; i++ {
		go func() {
			sum := 0
			for i := 0; i < times; i++ {
				if key {
					break
				}
				//fhttp.Response(get)
				get1()
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

func TestPerFast(t *testing.T) {
	c := make(chan int)
	start := time.Now().UnixMilli()
	for i := 0; i < thread; i++ {
		go func() {
			sum := 0
			for i := 0; i < times; i++ {
				if key {
					break
				}
				get := fhttp.FastGet(url, nil)
				fhttp.FastResponse(get)
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

// TestFaast
// @Description: 测试自定义DNS解析功能
// @param t
func TestFaast(t *testing.T) {
	url := "http://fun.tester:12345/test"
	get := fhttp.Get(url, nil)
	for i := 0; i < 10; i++ {
		go func() {
			log.Println(string(fhttp.Response(get)))
		}()
	}
	response := fhttp.Response(get)
	log.Println(string(response))

}

func TestFaa1st(t *testing.T) {
	jobs := make(chan func(), 1000)

	start := time.Now().UnixMilli()

	for w := 0; w < workers; w++ {
		go worker(jobs)
	}

	for j := 1; j <= maxCount; j++ {
		jobs <- get1
	}

	end := time.Now().UnixMilli()
	diff := end - start
	log.Printf("总耗时: %f", float64(diff)/1000)
	log.Printf("请求总数: %d", maxCount)
	log.Printf("QPS: %f", float64(maxCount)/float64(diff)*1000.0)
}

func worker(jobs <-chan func()) {
	for j := range jobs {
		j()
	}
}

func get1() {
	r.R().Get(url)
}

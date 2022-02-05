package httptest

import (
	"funtester/fhttp"
	"log"
	"testing"
	"time"
)

var key bool = false

const (
	url    = "fhttp://localhost:12345/test/fun"
	thread = 20
	times  = 5000
)

func TestPer(t *testing.T) {
	get := fhttp.Get(url, nil)
	c := make(chan int)

	start := time.Now().UnixMilli()
	for i := 0; i < thread; i++ {
		go func() {
			sum := 0
			for i := 0; i < times; i++ {
				if key {
					break
				}
				fhttp.Response(get)
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

func TestFaast(t *testing.T) {
	url := "http://fun.tester:12345/test"
	get := fhttp.Get(url, nil)
	for i := 0; i < 10; i++ {
		//go log.Println(string(fhttp.Response(get)))
		go func() {
			log.Println(string(fhttp.Response(get)))
		}()
	}
	response := fhttp.Response(get)
	log.Println(string(response))
}

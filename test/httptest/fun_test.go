package httptest

import (
	"funtester/funtester"
	"log"
	"testing"
	"time"
)

var key bool = false

const (
	url    = "http://localhost:8001/test/fun"
	thread = 20
	times  = 10000
)

func TestPer(t *testing.T) {
	get := funtester.Get(url, nil)
	c := make(chan int)

	start := time.Now().UnixMilli()
	for i := 0; i < thread; i++ {
		go func() {
			sum := 0
			for i := 0; i < times; i++ {
				if key {
					break
				}
				funtester.Response(get)
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
				get := funtester.FastGet(url, nil)
				funtester.FastResponse(get)
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

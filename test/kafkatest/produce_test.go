package kafkatest

import (
	"fmt"
	"funtester/execute"
	"funtester/ftool"
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"testing"
	"time"
)

func TestProduce(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.Compression = sarama.CompressionLZ4
	config.Producer.Timeout = time.Duration(50) * time.Millisecond
	config.Producer.Retry.Max = 3
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		// 关闭生产者
		if err = producer.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}()
	// 定义需要发送的消息
	headers := []sarama.RecordHeader{sarama.RecordHeader{
		Key:   []byte("funtest"),
		Value: []byte("have fun ~"),
	}}

	msg := &sarama.ProducerMessage{
		Topic:   "topic_test",
		Key:     sarama.StringEncoder("test"),
		Value:   sarama.StringEncoder("ddddddddddddddddd"),
		Headers: headers,
	}
	// 发送消息，并获取该消息的分片、偏移量
	for i := 0; i < 100; i++ {
		ftool.Sleep(1000)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("partition:%d offset:%d\n", partition, offset)
	}

	execute.ExecuteRoutineTimes(func() {
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("partition:%d offset:%d\n", partition, offset)

	}, 100, 10)
}
func TestPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new goroutine")
			return make(chan int, 100)
		},
	}

	for i := 0; i < 5; i++ {
		go func() {
			ch := pool.Get().(chan int)
			log.Println(len(ch))
			defer pool.Put(ch)
			ch <- i
			log.Printf("ssss %d ", i)
		}()
	}
	ftool.Sleep(1000)
	for i := 0; i < 5; i++ {
		log.Println(233333)
		fmt.Printf("main: %d", <-pool.Get().(chan int))
	}

}

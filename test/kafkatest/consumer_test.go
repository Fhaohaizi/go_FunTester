package kafkatest

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.Retry.Max = 3
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	topic := "topic_test"
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	defer consumer.Close()
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		log.Println(partition)
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
		}
		for msg := range pc.Messages() {
			log.Println(string(msg.Value))
			//log.Println(string(msg.Headers[0].Value))
			break
		}
		for {
			msg := <-pc.Messages()
			log.Println(string(msg.Value))
		}
	}
}

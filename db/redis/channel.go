package redis

import (
	"funtester/base"
	"funtester/ftool"
	"github.com/go-redis/redis"
	"log"
	"time"
)

// Subscribe
//  @Description: 订阅channel
//  @receiver r
//
func (r RedisBase) Subscribe(channels ...string) *redis.PubSub {
	subscribe := r.pool.Subscribe(channels...)
	initPublish(subscribe)
	//channel := subscribe.Channel()
	//for message := range channel {
	//	payload := message.Payload
	//	log.Printf("收到消息:%s", payload)
	//}
	return subscribe
}

// PSubcribe
//  @Description: 用法跟Subscribe一样，区别是PSubscribe订阅通道(channel)支持模式匹配。
//  @receiver r
//  @param channels
//  @return *PubSub
//
func (r RedisBase) PSubcribe(channels ...string) *redis.PubSub {
	subscribe := r.pool.PSubscribe(channels...)
	initPublish(subscribe)
	return subscribe
}

// Publish
//  @Description: 将消息发送到指定的channel
//  @receiver r
//  @param channel
//  @param message
//  @return int64
//
func (r RedisBase) Publish(channel string, message interface{}) int64 {
	result, err := r.pool.Publish(channel, message).Result()
	if err != nil {
		log.Printf("publish channel:%s,message:%s", channel, message)
		log.Println(err)
		return base.TestError
	}
	return result
}

// HandleMessage
//  @Description: 处理订阅消息
//  @param sub
//  @param msgFunc
//
func HandleMessage(sub *redis.PubSub, msgFunc func(msg string) bool) {
FUN:
	for {
		channel := sub.Channel()
		for message := range channel {
			b := msgFunc(message.Payload)
			if b {
				break FUN
			}
		}
	}
}

// HandleMessage2
//  @Description: 处理订阅消息
//  @param sub
//  @param msgFunc
//
func HandleMessage2(sub *redis.PubSub, msgFunc func(msg string, err error) bool) {
FUN:
	for {
		receive, err := sub.Receive()
		switch receive.(type) {
		case *redis.Subscription:
			log.Println("订阅成功")
		case redis.Message:
			message := receive.(redis.Message)
			b := msgFunc(message.Payload, err)
			if b {
				break FUN
			}
		case redis.Pong:
			log.Println("心跳消息")
		default:
			log.Printf("未知消息:%s", ftool.ToString(receive))
		}
	}

}

// initReceive
//  @Description: 订阅时候进行消息处理
//  @param client
//
func initPublish(client *redis.PubSub) {
	receive, err := client.ReceiveTimeout(time.Second)
	if err != nil {
		log.Printf("订阅失败")
		log.Fatal(err)
	}
	switch receive.(type) {
	case *redis.Subscription:
		log.Println("订阅成功")
	case *redis.Message:
		message := receive.(redis.Message)
		log.Printf("收到消息:%s", message.Payload)
	case *redis.Pong:
		log.Println("心跳消息")
	default:
		log.Printf("未知消息:%s", ftool.ToString(receive))
	}
}

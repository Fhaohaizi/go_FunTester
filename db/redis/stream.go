package redis

import (
	"funtester/base"
	"funtester/ftool"
	"github.com/go-redis/redis"
	"log"
)

// XAdd
//  @Description: 添加消息到末尾
//	args := &redis.XAddArgs{
//		Stream:       "fun2",
//		MaxLen:       0,
//		MaxLenApprox: 0,
//		ID:           "",
//		Values: map[string]interface{}{
//			"id":   "FunTester",
//			"name": 32,
//		},
//	}
//  @receiver r
//  @param a
//
func (r RedisBase) XAdd(args *redis.XAddArgs) string {
	result, err := r.pool.XAdd(args).Result()
	if err != nil {
		log.Printf("xadd args:%s fail\n", ftool.ToString(args))
		log.Println(err)
		return base.Empty
	}
	return result
}

// XRead
//  @Description: 以阻塞或非阻塞方式获取消息列表
//args := &redis.XReadArgs{
//Streams: []string{"fun2", "1659793290440"},
//Count:   10,
//Block:   time.Second,
//}
//	for _, result := range result {
//		for _, message := range result.Messages {
//			log.Println(message.ID)
//			log.Println(message.Values["name"])
//		}
//	}
//  @receiver r
//  @param a
//
func (r RedisBase) XRead(args *redis.XReadArgs) []redis.XStream {
	result, err := r.pool.XRead(args).Result()
	if err != nil {
		log.Printf("XRead args:%s fail\n", ftool.ToString(args))
		log.Println(err)
		return nil
	}
	return result
}

// XDel
//  @Description: 删除消息
//  @receiver r
//  @param stream
//  @param ids
//
func (r RedisBase) XDel(stream string, ids ...string) int64 {
	result, err := r.pool.XDel(stream, ids...).Result()
	if err != nil {
		log.Printf("XDel stream:%s ids:%s fail\n", stream, ftool.ToString(ids))
		log.Println(err)
		return base.TestError
	}
	return result
}

// XTrim
//  @Description: 对流进行修剪，限制长度
//  @receiver r
//  @param key
//  @param maxLen
//
func (r RedisBase) XTrim(key string, maxLen int64) int64 {
	result, err := r.pool.XTrim(key, maxLen).Result()
	if err != nil {
		log.Printf("XTrim stream:%s maxLen:%d fail\n", key, maxLen)
		log.Println(err)
		return base.TestError
	}
	return result
}

// XLen
//  @Description: 获取流包含的元素数量，即消息长度
//  @receiver r
//  @param key
//
func (r RedisBase) XLen(key string) int64 {
	result, err := r.pool.XLen(key).Result()
	if err != nil {
		log.Printf("XLen key:%s fail\n", key)
		log.Println(err)
		return base.TestError
	}
	return result
}

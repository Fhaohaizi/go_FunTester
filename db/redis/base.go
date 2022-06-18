package redis

import (
	"fmt"
	"funtester/base"
	"github.com/go-redis/redis"
	"log"
	"time"
)

//redis.Options 默认池大小为10×cpu核数

type RedisBase struct {
	Host string
	db   int
	pool *redis.Client
}

// init 初始化类,创建连接池
//  @Description:
//
func NewRdisPool(host string, db int) RedisBase {
	redisBase := RedisBase{Host: host, db: db}
	redisBase.pool = redis.NewClient(&redis.Options{
		//Password: "",
		Addr:            host,
		DB:              0,
		MaxRetries:      3,
		MinRetryBackoff: 100 * time.Millisecond,
		DialTimeout:     5 * time.Second,
		WriteTimeout:    1 * time.Second,
		PoolSize:        200,
		MaxConnAge:      10 * time.Second,
		IdleTimeout:     8 * time.Second,
	})
	_, err := redisBase.pool.Ping().Result()
	if err != nil {
		log.Fatal("连接失败", err)
	}
	log.Println("Redis 连接成功")
	ping := redisBase.Ping()
	if ping == "PONG" {
		log.Println("确认连接成功!")
	}
	return redisBase
}

func (r RedisBase) Ping() string {
	ping := r.pool.Ping()
	result, err := ping.Result()
	if err != nil {
		log.Println("确认连接失败")
	}
	return result
}

// Keys 获取所有的服务条件的keys列表
//  @Description:
//  @param patten
//  @return []string
//
func (r RedisBase) Keys(patten string) []string {
	result, err := r.pool.Keys(patten).Result()
	if err != nil {
		log.Printf("获取keys: %s 失败%s\n", patten, err.Error())
		return nil
	}
	return result
}

// Set 设置一个key的值
//  @Description:
//  @param key
//  @param value
//  @param expiration
//  @return string
//
func (r RedisBase) Set(key string, value interface{}, second time.Duration) string {
	result, err := r.pool.Set(key, value, time.Duration(second)*time.Second).Result()
	if err != nil {
		log.Printf("set:%s value: %s 失败\n", key, value)
		return base.Empty
	}
	return result
}

// Get 查询key的值
//  @Description:
//  @param key
//  @return string
//
func (r RedisBase) Get(key string) string {
	result, err := r.pool.Get(key).Result()
	if err != nil {
		log.Printf("get:%s 失败\n", key)
		return base.Empty
	}
	return result
}

// GetSet 设置一个key的值，并返回这个key的旧值
//  @Description:
//  @param key
//  @param value
//  @return string
//
func (r RedisBase) GetSet(key string, value interface{}) string {
	result, err := r.pool.GetSet(key, value).Result()
	if err != nil {
		log.Printf("set:%s value: %s 失败\n", key, value)
		return base.Empty
	}
	return result
}

// SetNX 如果key不存在，则设置这个key的值
//  @Description:
//  @param key
//  @param value
//  @param expiration
//  @return bool
//
func (r RedisBase) SetNX(key string, value interface{}, second int64) bool {
	result, err := r.pool.SetNX(key, value, time.Duration(second)*time.Second).Result()
	if err != nil {
		log.Printf("set:%s value: %s 失败\n", key, value)
		return false
	}
	return result
}

// MGet 批量查询key的值
//  @Description:
//  @param key
//  @param value
//  @param expiration
//  @return bool
//
func (r RedisBase) MGet(keys ...string) []interface{} {
	result, err := r.pool.MGet(keys...).Result()
	if err != nil {
		log.Printf("获取 keys : %s 失败 %s", fmt.Sprint(keys), err.Error())
		return nil
	}
	return result
}

// MSet 批量设置key的值
//  @Description:
//  @param keys
//  @return string
//
func (r RedisBase) MSet(keys ...string) string {
	result, err := r.pool.MSet(keys).Result()
	if err != nil {
		log.Printf("设置 keys : %s 失败 %s", fmt.Sprint(keys), err.Error())
		return base.Empty
	}
	return result
}

// Incr 针对一个key的数值进行递增操作
//  @Description:
//  @param key
//  @return string
//
func (r RedisBase) Incr(key string) int64 {
	result, err := r.pool.Incr(key).Result()
	if err != nil {
		log.Printf("自增 key: %s 失败 %s", key, err.Error())
		return base.TEST_ERROR
	}
	return result
}

// IncrBy 针对一个key的数值进行递增操作
//  @Description:
//  @param key
//  @param value
//  @return string
//
func (r RedisBase) IncrBy(key string, value int64) int64 {
	result, err := r.pool.IncrBy(key, value).Result()
	if err != nil {
		log.Printf("自增 key: %s 失败 %s", key, err.Error())
		return -1
	}
	return result
}

// Decr 针对一个key的数值进行递减操作
//  @Description:
//  @param key
//  @return string
//
func (r RedisBase) Decr(key string) int64 {
	result, err := r.pool.Decr(key).Result()
	if err != nil {
		log.Printf("自减 key: %s 失败 %s", key, err.Error())
		return base.TEST_ERROR
	}
	return result
}

// DecrBy 针对一个key的数值进行递减操作
//  @Description:
//  @param key
//  @param value
//  @return string
//
func (r RedisBase) DecrBy(key string, value int64) int64 {
	result, err := r.pool.DecrBy(key, value).Result()
	if err != nil {
		log.Printf("自减 key: %s 失败 %s", key, err.Error())
		return base.TEST_ERROR
	}
	return result
}

// Del 删除key操作，支持批量删除
//  @Description:
//  @param keys
//  @return int64
//
func (r RedisBase) Del(keys ...string) int64 {
	result, err := r.pool.Del(keys...).Result()
	if err != nil {
		log.Printf("删除 key: %s 失败 %s", fmt.Sprintln(keys), err.Error())
		return base.TEST_ERROR
	}
	return result
}

// Expire 设置key的过期时间,单位秒
//  @Description:
//  @param key
//  @param second
//  @return bool
//
func (r RedisBase) Expire(key string, second int64) bool {
	result, err := r.pool.Expire(key, time.Duration(second)*time.Second).Result()
	if err != nil {
		log.Printf("设置 key: %s 过期时间失败 %s", fmt.Sprintln(key), err.Error())
		return false
	}
	return result
}

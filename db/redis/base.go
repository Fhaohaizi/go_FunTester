package redis

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

//redis.Options 默认池大小为10×cpu核数

type RedisBase struct {
	Host string
	Db   int
	Pwd  string
	pool *redis.Client
}

// init 初始化类,创建连接池
//  @Description:
//
func NewRdisPool(host string, db int) RedisBase {
	redisBase := RedisBase{Host: host, Db: db}
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

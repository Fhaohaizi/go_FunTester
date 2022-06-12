package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	log.Println("Publisher started")
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Unable to connect to Redis", err)
	}
	log.Println("Connected to Redis server")
}

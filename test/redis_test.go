package test

import (
	"funtester/db/redis"
	"funtester/ftool"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

var pool = redis.GetClient()

func TestRedis(t *testing.T) {
	var str = "FunTester"
	set := redis.Set("fun", str, 0)
	log.Print(set)
	get := redis.Get("fun")
	assert.Equal(t, get, str)
	getSet := redis.GetSet("fun", str+ftool.RandomStr(3))
	log.Println(getSet)
	redis.Set("aa", "32342", 0)
	mGet := redis.MGet("fun", "aa")
	for i, i2 := range mGet {
		log.Printf("index :%d  value : %s\n", i, i2)
	}
	redis.Incr("sum")
	redis.Expire("fun", 300)
	keys := redis.Keys("fu*")
	for i, key := range keys {
		log.Printf("index : %d, key : %s", i, key)
	}
}

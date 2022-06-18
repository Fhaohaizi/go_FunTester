package test

import (
	"funtester/db/redis"
	"funtester/ftool"
	"github.com/go-playground/assert/v2"
	"log"
	"strconv"
	"testing"
)

var pool = redis.NewRdisPool("127.0.0.1:6379", 1)

func TestRedis(t *testing.T) {
	var str = "FunTester"
	set := pool.Set("fun", str, 0)
	log.Print(set)
	get := pool.Get("fun")
	assert.Equal(t, get, str)
	getSet := pool.GetSet("fun", str+ftool.RandomStr(3))
	log.Println(getSet)
	pool.Set("aa", "32342", 0)
	mGet := pool.MGet("fun", "aa")
	for i, i2 := range mGet {
		log.Printf("index :%d  value : %s\n", i, i2)
	}
	pool.Expire("fun", 300)
	keys := pool.Keys("fu*")
	for i, key := range keys {
		log.Printf("index : %d, key : %s", i, key)
	}
	key := str + strconv.Itoa(ftool.RandomInt(1000))
	pool.SetNX(key, "32432", 100)
	log.Println(pool.Get(key))
	log.Println("22222")
	i := pool.MGet(str, key)
	for _, i3 := range i {
		log.Println(i3)
	}
	pool.MSet("aa", "111", "aabbbb", "22222")
	pool.Incr("sum")
	pool.IncrBy("sum", 10)
	log.Println(pool.Get("sum"))
	strings := pool.Keys("aa*")
	for _, s := range strings {
		log.Println(s)
	}
	pool.Decr("aa")
	pool.Expire("sum", 100)
}

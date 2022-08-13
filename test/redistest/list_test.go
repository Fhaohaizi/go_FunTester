package redistest

import (
	"funtester/base"
	"funtester/db/redis"
	"github.com/go-playground/assert/v2"
	"log"
	"strconv"
	"testing"
)

func TestPushPop(t *testing.T) {
	var pool = redis.NewRdisPool("127.0.0.1:6379", base.Empty, 1)
	var key = "fun"
	var value = "abc"
	pool.LPush(key, value)
	pop := pool.LPop(key)
	assert.Equal(t, pop, value)
}

func TestPushPop2(t *testing.T) {
	var pool = redis.NewRdisPool("127.0.0.1:6379", base.Empty, 1)
	var key = "fun"
	var value = "abc"
	pool.RPush(key, value)
	pop := pool.RPop(key)
	assert.Equal(t, pop, value)
}

func TestLen(t *testing.T) {
	var pool = redis.NewRdisPool("127.0.0.1:6379", base.Empty, 1)
	var key = "fun"
	var value = "FunTester"
	len1 := pool.LLen(key)
	pool.LPush(key, value)
	len2 := pool.LLen(key)
	pool.LPop(key)
	len3 := pool.LLen(key)
	assert.Equal(t, len1, len3)
	assert.Equal(t, len1, len2-1)
}

func TestLRange(t *testing.T) {
	var pool = redis.NewRdisPool("127.0.0.1:6379", base.Empty, 1)
	var key = "fun"
	var value = "FunTester"
	pool.Del(key)
	for i := 0; i < 10; i++ {
		pool.RPush(key, value+strconv.Itoa(i))
	}
	lRange := pool.LRange(key, 1, 3)
	for _, s := range lRange {
		log.Println(s)
	}

}

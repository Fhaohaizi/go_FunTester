package redistest

import (
	"funtester/base"
	"funtester/db/redis"
	"funtester/execute"
	"funtester/ftool"
	"log"
	"strings"
	"testing"
)

func TestCollect(t *testing.T) {
	var pool = redis.NewRdisPool("127.0.0.1:6379", base.Empty, 1)
	var key = "32"
	//for i := 0; i < 10; i++ {
	//	pool.HSet(key, strconv.Itoa(i), ftool.RandomStr(30))
	//}
	//get := pool.HGet(key, "3")
	//log.Println(get)
	//all := pool.HGetAll(key)
	//for k, v := range all {
	//	log.Printf("k: %s , v: %s \n", k, v)
	//}
	//by := pool.HIncrBy("fun", "1", 1)
	//log.Println(by)
	keys := pool.HKeys(key)
	log.Println(len(keys))
	for i := range keys {
		println(i)
	}
	get := pool.HMGet(key, "1", "2")
	for i, i2 := range get {
		log.Println(i)
		log.Println(i2)
	}
	m := make(map[string]interface{})
	m["f"] = 3
	m["f32"] = 343
	m["2f32"] = 343e2
	set := pool.HMSet(key, m)
	set = strings.ToLower(set)
	log.Println(set == "ok")
}

func TestSet(t *testing.T) {
	var pool = redis.NewRdisPool("127.0.0.1:6379", base.Empty, 1)
	var key = "funset"
	add := pool.SAdd(key, ftool.RandomStr(10), ftool.RandomStr(3), "FunTester")
	log.Println(add)
	log.Printf("FunTester 是否存在 %t", pool.SIsMember(key, "FunTester"))
	log.Println(pool.SCard(key))
	members := pool.SMembers(key)
	for _, s := range members {
		log.Println(s)
		if ftool.RandomInt(2) == 1 && s != "FunTester" {
			pool.SRem(key, s)
		}
	}
	log.Println(pool.SRem(key, "FunTester", "000000000"))
	log.Println(pool.SPop(key))
	log.Println(len(pool.SPopN(key, 1000)))
}

func TestSetPer(t *testing.T) {
	var pool = redis.NewRdisPool("127.0.0.1:6379", base.Empty, 1)
	var key = "funsetper"
	execute.ExecuteRoutineTime(func() {
		add := int(pool.SAdd(key, ftool.RandomStr(10), ftool.RandomStr(3)))
		for i := 0; i < add; i++ {
			pool.SPop(key)
		}

	}, 20, 5)
}

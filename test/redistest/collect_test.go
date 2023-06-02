package redistest

import (
	"funtester/base"
	"funtester/db/redis"
	"funtester/execute"
	"funtester/ftool"
	r "github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v5"
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

type Funtester struct {
	Name string
	Age  int
}

//func (f Funtester) MarshalBinary() (data []byte, err error) {
//	return f.Marshal(), nil
//}

func (f Funtester) Marshal() []byte {
	marshal, err := msgpack.Marshal(f)
	if err != nil {
		log.Println(err)
		return nil
	}
	return marshal
}

func (f *Funtester) UnMarshal(data []byte) {
	msgpack.Unmarshal(data, f)
}

func (f *Funtester) UnMarshalStr(s string) {
	msgpack.Unmarshal([]byte(s), f)
}

func TestSorted(t *testing.T) {
	var pool = redis.NewRdisPool("127.0.0.1:6379", base.Empty, 2)
	var key = "redis sorted"
	ob := Funtester{
		Name: "我是FunTester" + ftool.RandomStr(10),
		Age:  ftool.RandomInt(100),
	}
	z := r.Z{
		Score:  float64(ftool.RandomInt(100)),
		Member: ob.Marshal(),
	}
	pool.ZAdd(key, z)
	log.Printf("总数量:%d", pool.ZCard(key))
	log.Printf("50 分以上总数量:%d", pool.ZCount(key, "50", "99"))
	//pool.ZIncrBy(key, 1000, string(ob.Marshal()))
	zRange := pool.ZRange(key, 0, 1)
	for i := range zRange {
		v := zRange[i]
		s := &Funtester{}
		s.UnMarshalStr(v)
		log.Println(s.Age)
	}
	log.Println("------------")
	scores := pool.ZRangeByScoreWithScores(key, r.ZRangeBy{
		Min:    "2",    // 最小分数
		Max:    "1000", // 最大分数
		Offset: 0,      // 表示开始偏移量
		Count:  2,      // 一次返回多少数据
	})
	for i, score := range scores {
		log.Println(i)
		log.Println(score.Score)
		member := score.Member
		var f Funtester
		f.UnMarshalStr(ftool.ToString(member))
		log.Println(f.Name)
	}
	pool.ZRemRangeByRank(key)
}

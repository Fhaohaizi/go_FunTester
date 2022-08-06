package redis

import (
	"funtester/base"
	"funtester/ftool"
	"github.com/go-redis/redis"
	"log"
)

// ZAdd
//  @Description: 添加一个或者多个元素到集合，如果元素已经存在则更新分数
//  @receiver r
//  @param key
//  @param members
//  @return int64
//
func (r RedisBase) ZAdd(key string, members ...redis.Z) int64 {
	result, err := r.pool.ZAddCh(key, members...).Result()
	if err != nil {
		log.Printf("zadd key:%s, members:%s", key, ftool.ToString(members))
		log.Println(err)
		return base.TestError
	}
	return result
}

// ZCard
//  @Description: 返回集合元素个数
//  @receiver r
//  @param key
//  @return int64
//
func (r RedisBase) ZCard(key string) int64 {
	result, err := r.pool.ZCard(key).Result()
	if err != nil {
		log.Printf("ZCard key:%s fail\n", key)
		log.Println(err)
		return base.TestError
	}
	return result
}

// ZCount
//  @Description: 统计某个分数范围内的元素个数
// 返回： 1<分数<=5 的元素个数
// 说明：默认第二，第三个参数是大于等于和小于等于的关系。
// 如果加上（ 则表示大于或者小于，相当于去掉了等于关系。
// size, err := client.ZCount("key", "(1","5").Result()
//  @receiver r
//  @param key
//  @param min
//  @param max
//  @return int64
//
func (r RedisBase) ZCount(key string, min, max string) int64 {
	result, err := r.pool.ZCount(key, min, max).Result()
	if err != nil {
		log.Printf("ZCard key:%s fail\n", key)
		log.Println(err)
		return base.TestError
	}
	return result
}

// ZIncrBy
//  @Description:增加元素的分数
//  @receiver r
//  @param key
//  @param increment
//  @param member
//  @return float64
//
func (r RedisBase) ZIncrBy(key string, increment float64, member string) float64 {
	result, err := r.pool.ZIncrBy(key, increment, member).Result()
	if err != nil {
		log.Printf("ZIncrBy key:%s increment:%f ,member:%s fail\n", key, increment, member)
		log.Println(err)
		return base.TestError
	}
	return result
}

// ZRange
//  @Description:返回集合中某个索引范围的元素，根据分数从小到大排序
// 返回从0到-1位置的集合元素， 元素按分数从小到大排序
// 0到-1代表则返回全部数据
//  @receiver r
//  @param key
//  @param start
//  @param stop
//  @return []string
//
func (r RedisBase) ZRange(key string, start, stop int64) []string {
	result, err := r.pool.ZRange(key, start, stop).Result()
	if err != nil {
		log.Printf("ZRange key:%s start:%d ,stop:%d fail\n", key, start, stop)
		log.Println(err)
		return nil
	}
	return result
}

// ZRevRange
//  @Description:返回集合中某个索引范围的元素，根据分数从大到小排序
//  @receiver r
//  @param key
//  @param start
//  @param stop
//  @return []string
//
func (r RedisBase) ZRevRange(key string, start, stop int64) []string {
	result, err := r.pool.ZRevRange(key, start, stop).Result()
	if err != nil {
		log.Printf("ZRevRange key:%s start:%d ,stop:%d fail\n", key, start, stop)
		log.Println(err)
		return nil
	}
	return result
}

// ZRangeByScore
//  @Description:根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
//op := redis.ZRangeBy{
//	Min:"2", // 最小分数
//	Max:"10", // 最大分数
//	Offset:0, // 类似sql的limit, 表示开始偏移量
//	Count:5, // 一次返回多少数据
//}
//  @receiver r
//  @param key
//  @param opt
//  @return []string
//
func (r RedisBase) ZRangeByScore(key string, opt redis.ZRangeBy) []string {
	result, err := r.pool.ZRangeByScore(key, opt).Result()
	if err != nil {
		log.Printf("ZRangeByScore key:%s opt:%s fail\n", key, ftool.ToString(opt))
		log.Println(err)
		return nil
	}
	return result
}

// ZRevRangeByScore
//  @Description:用法类似ZRangeByScore，区别是元素根据分数从大到小排序。
//  @receiver r
//  @param key
//  @param opt
//  @return []string
//
func (r RedisBase) ZRevRangeByScore(key string, opt redis.ZRangeBy) []string {
	result, err := r.pool.ZRevRangeByScore(key, opt).Result()
	if err != nil {
		log.Printf("ZRevRangeByScore key:%s opt:%s fail\n", key, ftool.ToString(opt))
		log.Println(err)
		return nil
	}
	return result
}

// ZRangeByScoreWithScores
//  @Description:用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
//  @receiver r
//  @param key
//  @param opt
//  @return []string
//
func (r RedisBase) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) []redis.Z {
	result, err := r.pool.ZRangeByScoreWithScores(key, opt).Result()
	if err != nil {
		log.Printf("ZRangeByScoreWithScores key:%s opt:%s fail\n", key, ftool.ToString(opt))
		log.Println(err)
		return nil
	}
	return result
}

// ZRem
//  @Description:删除集合元素
//  @receiver r
//  @param key
//  @param members
//  @return int64
//
func (r RedisBase) ZRem(key string, members ...interface{}) int64 {
	result, err := r.pool.ZRem(key, members).Result()
	if err != nil {
		log.Printf("ZRem key:%s members:%s fail\n", key, ftool.ToString(members))
		log.Println(err)
		return base.TestError
	}
	return result
}

// ZRemRangeByRank
//  @Description:根据索引范围删除元素
// 集合元素按分数排序，从最低分到高分，删除第0个元素到第5个元素。
// 这里相当于删除最低分的几个元素
//client.ZRemRangeByRank("key", 0, 5)
// 位置参数写成负数，代表从高分开始删除。
// 这个例子，删除最高分数的两个元素，-1代表最高分数的位置，-2第二高分，以此类推。
//client.ZRemRangeByRank("key", -1, -2)
//  @receiver r
//  @param key
//  @param start
//  @param stop
//  @return int64
//
func (r RedisBase) ZRemRangeByRank(key string, start, stop int64) int64 {
	result, err := r.pool.ZRemRangeByRank(key, start, stop).Result()
	if err != nil {
		log.Printf("ZRemRangeByRank key:%s start:%d ,stop:%d fail\n", key, start, stop)
		log.Println(err)
		return base.TestError
	}
	return result
}

// ZRemRangeByScore
//  @Description:根据分数范围删除元素
// 删除范围： 2<=分数<=5 的元素
//client.ZRemRangeByScore("key", "2", "5")
// 删除范围： 2<=分数<5 的元素
//client.ZRemRangeByScore("key", "2", "(5")
//  @receiver r
//  @param key
//  @param start
//  @param stop
//  @return int64
//
func (r RedisBase) ZRemRangeByScore(key string, start, stop string) int64 {
	result, err := r.pool.ZRemRangeByScore(key, start, stop).Result()
	if err != nil {
		log.Printf("ZRemRangeByScore key:%s start:%s ,stop:%s fail\n", key, start, stop)
		log.Println(err)
		return base.TestError
	}
	return result
}

// ZScore
//  @Description:查询元素对应的分数
//  @receiver r
//  @param key
//  @param member
//  @return float64
//
func (r RedisBase) ZScore(key, member string) float64 {
	result, err := r.pool.ZScore(key, member).Result()
	if err != nil {
		log.Printf("ZScore key:%s member:%s fail\n", key, member)
		log.Println(err)
		return base.TestError
	}
	return result
}

// ZRank
//  @Description:根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
//  @receiver r
//  @param key
//  @param member
//  @return float64
//
func (r RedisBase) ZRank(key, member string) int64 {
	result, err := r.pool.ZRank(key, member).Result()
	if err != nil {
		log.Printf("ZRank key:%s member:%s fail\n", key, member)
		log.Println(err)
		return base.TestError
	}
	return result
}

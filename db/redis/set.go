package redis

import (
	"funtester/base"
	"funtester/ftool"
	"log"
)

// SAdd
//  @Description: 添加集合元素
//  @receiver r
//  @param key
//  @param members
//  @return int64
//
func (r RedisBase) SAdd(key string, members ...interface{}) int64 {
	result, err := r.pool.SAdd(key, members...).Result()
	if err != nil {
		log.Printf("sadd key:%s,members:%s fail\n", key, ftool.ToString(members))
		log.Println(err)
		return base.TestError
	}
	return result
}

// SCard
//  @Description: 获取集合元素个数
//  @receiver r
//  @param key
//  @return int64
//
func (r RedisBase) SCard(key string) int64 {
	result, err := r.pool.SCard(key).Result()
	if err != nil {
		log.Printf("scard key:%s fail\n", key)
		log.Println(err)
		return base.TestError
	}
	return result
}

// SIsMember
//  @Description: 判断元素是否在集合中
//  @receiver r
//  @param key
//  @param member
//  @return bool
//
func (r RedisBase) SIsMember(key string, member interface{}) bool {
	result, err := r.pool.SIsMember(key, member).Result()
	if err != nil {
		log.Printf("sismember key:%s,member:%s fail\n", key, ftool.ToString(member))
		log.Println(err)
		return false
	}
	return result
}

// SMembers
//  @Description: 获取集合中所有的元素
//  @receiver r
//  @param key
//  @return []string
//
func (r RedisBase) SMembers(key string) []string {
	result, err := r.pool.SMembers(key).Result()
	if err != nil {
		log.Printf("smember key:%s fail\n", key)
		log.Println(err)
		return nil
	}
	return result
}

// SRem
//  @Description: 删除集合元素
//  @receiver r
//  @param key
//  @param members
//
func (r RedisBase) SRem(key string, members ...interface{}) int64 {
	result, err := r.pool.SRem(key, members...).Result()
	if err != nil {
		log.Printf("srem key:%s members:%s fail\n", key, ftool.ToString(members))
		log.Println(err)
		return base.TestError
	}
	return result
}

// SPop
//  @Description: 随机返回集合中的元素，并且删除返回的元素
//  @receiver r
//  @param key
//  @return string
//
func (r RedisBase) SPop(key string) string {
	result, err := r.pool.SPop(key).Result()
	if err != nil {
		log.Printf("spop key:%s fail\n", key)
		log.Println(err)
		return base.Empty
	}
	return result
}

// SPopN
//  @Description: 随机返回集合中的元素，并且删除返回的元素
//  @receiver r
//  @param key
//  @param count
//  @return []string
//
func (r RedisBase) SPopN(key string, count int64) []string {
	result, err := r.pool.SPopN(key, count).Result()
	if err != nil {
		log.Printf("spop key:%s fail\n", key)
		log.Println(err)
		return nil
	}
	return result
}

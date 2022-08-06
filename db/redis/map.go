package redis

import (
	"funtester/base"
	"funtester/ftool"
	"log"
)

// HSet
//  @Description: hashset 值
//  @receiver r redis基础对象
//  @param key HashMap key
//  @param field hash内部key
//  @param value hash内部value
//  @return bool 是否成功
//
func (r RedisBase) HSet(key, field string, value interface{}) bool {
	result, err := r.pool.HSet(key, field, value).Result()
	if err != nil {
		log.Printf("hset key : %s filed: %s value: f%s fail\n", key, field, value)
		log.Println(err)
		return false
	}
	return result
}

// HGet
//  @Description: hashmap 获取值
//  @receiver r
//  @param key
//  @param field
//  @return string
//
func (r RedisBase) HGet(key, field string) string {
	result, err := r.pool.HGet(key, field).Result()
	if err != nil {
		log.Printf("hget key: %s,files:%s fail\n", key, field)
		log.Println(err)
		return base.Empty
	}
	return result
}

// HGetAll
//  @Description: 获取所有key和value
//  @receiver r
//  @param key
//  @return map[string]string
//
func (r RedisBase) HGetAll(key string) map[string]string {
	result, err := r.pool.HGetAll(key).Result()
	if err != nil {
		log.Printf("hgetall key: %s fail\n", key)
		log.Println(err)
		return nil
	}
	return result
}

// HIncrBy 根据key和field字段，累加字段的数值
//  @Description:
//  @receiver r
//  @param key
//  @param field
//  @param incr
//  @return int64
//
func (r RedisBase) HIncrBy(key, field string, incr int64) int64 {
	result, err := r.pool.HIncrBy(key, field, incr).Result()
	if err != nil {
		log.Printf("hincrby key:%s field:%s,incr: %d fail\n", key, field, incr)
		log.Println(err)
		return base.TestError
	}
	return result
}

// HKeys
//  @Description: 根据key返回所有字段名
//  @receiver r
//  @param key
//  @return []string
//
func (r RedisBase) HKeys(key string) []string {
	result, err := r.pool.HKeys(key).Result()
	if err != nil {
		log.Printf("hkeys key:%s fail\n", key)
		log.Println(err)
		return nil
	}
	return result
}

// HLen
//  @Description: 根据key，查询hash的字段数量
//  @receiver r
//  @param key
//  @return int64
//
func (r RedisBase) HLen(key string) int64 {
	result, err := r.pool.HLen(key).Result()
	if err != nil {
		log.Printf("HLen key:%s fail\n", key)
		log.Println(err)
		return base.TestError
	}
	return result
}

// HMGet
//  @Description: 根据key和多个字段名，批量查询多个hash字段值
//  @receiver r
//  @param key
//  @param fields
//
func (r RedisBase) HMGet(key string, fields ...string) []interface{} {
	result, err := r.pool.HMGet(key, fields...).Result()
	if err != nil {
		log.Printf("HMGet key:%s,field:%s fail:%d\n", key, ftool.ToString(fields))
		log.Println(err)
		return nil
	}
	return result
}

// HMSet
//  @Description:根据key和多个字段名和字段值，批量设置hash字段值
//  @receiver r
//  @param key
//  @param fields
//  @return string ok
//
func (r RedisBase) HMSet(key string, fields map[string]interface{}) string {
	result, err := r.pool.HMSet(key, fields).Result()
	if err != nil {
		log.Printf("hmset key:%s, fields:%s\n", key, ftool.ToString(fields))
		log.Println(err)
		return base.Empty
	}
	return result
}

// HSetNX
//  @Description:如果field字段不存在，则设置hash字段值
//  @receiver r
//  @param key
//  @param field
//  @param value
//  @return bool
//
func (r RedisBase) HSetNX(key, field string, value interface{}) bool {
	result, err := r.pool.HSetNX(key, field, value).Result()
	if err != nil {
		log.Printf("hsetnx key:%s ,field: %s ,value:%s fail\n", key, field, value)
		log.Println(err)
		return false
	}
	return result
}

// HDel
//  @Description: 根据key和字段名，删除hash字段，支持批量删除hash字段
//  @receiver r
//  @param key
//  @param field
//  @return int64
//
func (r RedisBase) HDel(key string, field ...string) int64 {
	result, err := r.pool.HDel(key, field...).Result()
	if err != nil {
		log.Printf("hdel key:%s ,field: %s fail\n", key, ftool.ToString(field))
		log.Println(err)
		return base.TestError
	}
	return result
}

// HExists
//  @Description: 检测hash字段名是否存在。
//  @receiver r
//  @param key
//  @param field
//  @return bool
//
func (r RedisBase) HExists(key, field string) bool {
	result, err := r.pool.HExists(key, field).Result()
	if err != nil {
		log.Printf("hexists key:%s , field:%s fail\n", key, field)
		println(err)
		return false
	}
	return result
}

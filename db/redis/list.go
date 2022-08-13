package redis

import (
	"funtester/base"
	"log"
)

// LPush
//  @Description: 从列表左边插入数据
//  @receiver r
//  @param key
//  @param value
//  @return int64
//
func (r RedisBase) LPush(key string, value interface{}) int64 {
	result, err := r.pool.LPush(key, value).Result()
	if err != nil {
		log.Printf("LPush:%s value: %s 失败\n", key, value)
		log.Println(err)
		return base.TestError
	}
	return result
}

// RPush
//  @Description:从列表右边插入数据
//  @receiver r
//  @param key
//  @param value
//  @return int64
//
func (r RedisBase) RPush(key string, value interface{}) int64 {
	result, err := r.pool.RPush(key, value).Result()
	if err != nil {
		log.Printf("RPush:%s value: %s 失败\n", key, value)
		log.Println(err)
		return base.TestError
	}
	return result
}

// LPop
//  @Description:  从列表左边删除第一个数据，并返回删除的数据
//  @receiver r
//  @param key
//  @param second
//  @return string
//
func (r RedisBase) LPop(key string) string {
	result, err := r.pool.LPop(key).Result()
	if err != nil {
		log.Printf("LPop:%s 失败\n", key)
		log.Println(err)
		return base.Empty
	}
	return result
}

// RPop
//  @Description:  从列表的右边删除第一个数据，并返回删除的数据
//  @receiver r
//  @param key
//  @param second
//  @return string
//
func (r RedisBase) RPop(key string) string {
	result, err := r.pool.RPop(key).Result()
	if err != nil {
		log.Printf("RPop:%s 失败\n", key)
		log.Println(err)
		return base.Empty
	}
	return result
}

// LLen
//  @Description: 返回列表的大小
//  @receiver r
//  @param key
//  @return int64
//
func (r RedisBase) LLen(key string) int64 {
	result, err := r.pool.LLen(key).Result()
	if err != nil {
		log.Printf("LLen :%s 失败\n", key)
		log.Println(err)
		return base.TestError
	}
	return result
}

// LRange
//  @Description:返回列表的一个范围内的数据，也可以返回全部数据
//  @receiver r
//  @param key
//  @param start
//  @param end
//  @return []string
//
func (r RedisBase) LRange(key string, start, end int64) []string {
	result, err := r.pool.LRange(key, start, end).Result()
	if err != nil {
		log.Printf("LRange :%s 失败\n", key)
		log.Println(err)
		return nil
	}
	return result
}

// LRem
//  @Description:删除列表中的数据
//从列表左边开始删除， 如果出现重复元素，仅删除1次，也就是删除第一个(key,1,"FunTester")
// 如果存在多个，则从列表左边开始删除2个100(key,2,"FunTester")
// 如果存在多个，则从列表右边开始删除2个100(key,-2,"FunTester")
// 删除所有(key,0,value)
//  @receiver r
//  @param key
//  @param count
//  @param value
//  @return int64
//
func (r RedisBase) LRem(key string, count int64, value interface{}) int64 {
	result, err := r.pool.LRem(key, count, value).Result()
	if err != nil {
		log.Printf("LRem :%s count: %d value: %s 失败\n", key, count, value)
		log.Println(err)
		return base.TestError
	}
	return result
}

// LIndex
//  @Description:根据索引坐标，查询列表中的数据
//  @receiver r
//  @param key
//  @param index 索引,从0开始算
//  @return string
//
func (r RedisBase) LIndex(key string, index int64) string {
	result, err := r.pool.LIndex(key, index).Result()
	if err != nil {
		log.Printf("LIndex :%s index: %d 失败\n", key, index)
		log.Println(err)
		return base.Empty
	}
	return result
}

// LInsertBefore
//  @Description:在指定位置插入数据
//  @receiver r
//  @param key
//  @param pivot
//  @param value
//  @return int64
//
func (r RedisBase) LInsertBefore(key string, pivot, value interface{}) int64 {
	result, err := r.pool.LInsertBefore(key, pivot, value).Result()
	if err != nil {
		log.Printf("LInsertBefore :%s pivot: %s value:%s 失败\n", key, pivot, value)
		log.Println(err)
		return base.TestError
	}
	return result
}

// LInsertAfter
//  @Description:在指定位置插入数据
//  @receiver r
//  @param key
//  @param pivot
//  @param value
//  @return int64
//
func (r RedisBase) LInsertAfter(key string, pivot, value interface{}) int64 {
	result, err := r.pool.LInsertAfter(key, pivot, value).Result()
	if err != nil {
		log.Printf("LInsertAfter :%s pivot: %s value:%s 失败\n", key, pivot, value)
		log.Println(err)
		return base.TestError
	}
	return result
}

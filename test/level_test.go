package test

import (
	"funtester/task"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"testing"
	"time"
)

func TestLeveldb(t *testing.T) {
	db, err := leveldb.OpenFile("funtester", nil)
	if err != nil {
		log.Println("创建出错!", err)
	}
	fun := []byte("Have Fun ~ Tester ！")
	db.Put([]byte(task.FunTester), fun, nil)
	time.Sleep(1 * time.Second)
	get, _ := db.Get([]byte(task.FunTester), nil)
	log.Printf(string(get))

	db.Delete([]byte("test"), nil)
	iterator := db.NewIterator(nil, nil)
	for iterator.Next() {
		key := iterator.Key()
		value := iterator.Value()
		log.Printf("数据key:%s,value:%s", key, value)
	}
	db.Close()

}

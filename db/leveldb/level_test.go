package leveldb

import (
	"container/list"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func OpenDB(path string) *leveldb.DB {
	db, err := leveldb.OpenFile("funhttp", nil)
	if err != nil {
		log.Fatal("创建出错!", err)
	}
	return db
}

func Get(db *leveldb.DB, key string) string {
	get, err := db.Get([]byte(key), nil)
	if err != nil {
		log.Fatal("获取值出错了", err)
	}
	return string(get)

}

func Put(db *leveldb.DB, key, value string) {
	err := db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		log.Fatal("存储出错了", err)
	}
}

func Del(db *leveldb.DB, key string) {
	err := db.Delete([]byte(key), nil)
	if err != nil {
		log.Fatal("删除出错了", err)
	}
}

func AllKey(db *leveldb.DB) *list.List {
	l := list.New()
	iterator := db.NewIterator(nil, nil)
	for iterator.Next() {
		key := iterator.Key()
		l.PushBack(key)
	}
	return l
}

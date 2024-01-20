package pool

import (
	"log"
	"reflect"
	"testing"
)

type ObjectPool2 struct {
	objects chan *PooledObject
	Name    string
}

// NewPool
//
//	@Description: 创建对象池
//	@param size 对象池大小
//	@return *ObjectPool 对象类型
func NewPool2(size int) *ObjectPool2 {
	return &ObjectPool2{
		objects: make(chan *PooledObject, size),
		Name:    "FunTester测试",
	}
}

// Get
//
//	@Description: 获取对象
//	@receiver p 对象池
//	@return *PooledObject 对象
func (p *ObjectPool2) Get2() *PooledObject {
	select {
	case obj := <-p.objects:
		return obj
	default:
		log.Println("额外创建对象")
		return NewObject()
	}
}

// Back
//
//	@Description: 回收对象
//	@receiver p 对象池
//	@param obj 回收的对象
func (p *ObjectPool2) Back(obj *PooledObject) {
	obj.Reset()
	select {
	case p.objects <- obj:
	default:
		obj = nil
		log.Println("丢弃对象")
	}
}

func TestPool2(t *testing.T) {
	pool := NewPool2(1)
	get := pool.Get2()
	object := pool.Get2()
	log.Printf("%T", get)
	log.Println(reflect.TypeOf(get))
	pool.Back(get)
	pool.Back(object)

}

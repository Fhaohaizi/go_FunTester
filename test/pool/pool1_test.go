package pool

import (
	"log"
	"reflect"
	"testing"
)

// PooledObject
// @Description: 对象池对象
type PooledObject struct {
	Name    string
	Age     int
	Address string
}

// NewObject
//
//	@Description: 创建对象
//	@return *PooledObject
func NewObject() *PooledObject {
	return &PooledObject{
		Name:    "",
		Age:     0,
		Address: "",
	}

}

// Reset
//
//	@Description: 重置对象
//	@receiver m 对象
func (m *PooledObject) Reset() {
	m.Name = ""
	m.Age = 0
	m.Address = ""
}

type ObjectPool struct {
	objects chan *PooledObject
	Name    string
}

// NewPool
//
//	@Description: 创建对象池
//	@param size 对象池大小
//	@return *ObjectPool 对象类型
func NewPool(size int) *ObjectPool {
	return &ObjectPool{
		objects: make(chan *PooledObject, size),
		Name:    "FunTester测试",
	}
}

// Get
//
//	@Description: 获取对象
//	@receiver p 对象池
//	@return *PooledObject 对象
func (p *ObjectPool) Get() *PooledObject {
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
func (p *ObjectPool) Back(obj *PooledObject) {
	obj.Reset()
	select {
	case p.objects <- obj:
	default:
		log.Println("丢弃对象")
	}
}

func TestPool1(t *testing.T) {
	pool := NewPool(1)
	get := pool.Get()
	object := pool.Get()
	log.Printf("%T", get)
	log.Println(reflect.TypeOf(get))
	pool.Back(get)
	pool.Back(object)

}

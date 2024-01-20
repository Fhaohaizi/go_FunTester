package pool

import (
	"funtester/ftool"
	"log"
	"sync"
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
	log.Println("创建对象")
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
	log.Println("重置对象")
}

type ObjectPool struct {
	ObjPool sync.Pool
	Name    string
}

// NewPool
//
//	@Description: 创建对象池
//	@param size 对象池大小
//	@return *ObjectPool 对象类型
func NewPool(size int) *ObjectPool {
	return &ObjectPool{
		Name:    "FunTester测试",
		ObjPool: sync.Pool{New: func() interface{} { return NewObject() }},
	}
}

// Get
//
//	@Description: 获取对象
//	@receiver p 对象池
//	@return *PooledObject 对象
func (p *ObjectPool) Get() *PooledObject {
	return p.ObjPool.Get().(*PooledObject)
}

// Back
//
//	@Description: 回收对象
//	@receiver p 对象池
//	@param obj 回收的对象
func (p *ObjectPool) Back(obj *PooledObject) {
	obj.Reset()
	p.ObjPool.Put(obj)
}

func TestPool1(t *testing.T) {
	pool := NewPool(1)
	get := pool.Get()
	get.Name = "FunTester"
	get.Age = 18
	get.Address = "地球"
	log.Printf("%T %s", get, ftool.ToString(get))
	pool.Back(get)
	get2 := pool.Get()
	log.Printf("%T %s", get, ftool.ToString(get2))

}

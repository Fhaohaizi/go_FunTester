package execute

import (
	"errors"
	"funtester/ftool"
	"log"
	"sync/atomic"
	"time"
)

type GorotinesPool struct {
	Max          int
	Min          int
	tasks        chan func() taskType
	status       bool
	active       int32
	ReceiveTotal int32
	ExecuteTotal int32
	SingleTimes  int
	addTimeout   time.Duration
}

type taskType int

const (
	normal taskType = 0
	reduce taskType = 1
)

// GetPool
//  @Description: 创建线程池
//  @param max 最大协程数
//  @param min 最小协程数
//  @param maxWaitTask 最大任务等待长度
//  @param timeout 添加任务超时时间，单位s
//  @return *GorotinesPool
//
func GetPool(max, min, maxWaitTask, timeout int) *GorotinesPool {
	p := &GorotinesPool{
		Max:          max,
		Min:          min,
		tasks:        make(chan func() taskType, maxWaitTask),
		status:       true,
		active:       0,
		ReceiveTotal: 0,
		ExecuteTotal: 0,
		SingleTimes:  10,
		addTimeout:   time.Duration(timeout) * time.Second,
	}
	for i := 0; i < min; i++ {
		atomic.AddInt32(&p.active, 1)
		go p.worker()
	}
	go func() {
		for {
			if !p.status {
				break
			}
			ftool.Sleep(1000)
			p.balance()
		}
	}()
	return p
}

// worker
//  @Description: 开始执行协程
//  @receiver pool
//
func (pool *GorotinesPool) worker() {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("execute task fail: %v", p)
		}
	}()
Fun:
	for t := range pool.tasks {
		atomic.AddInt32(&pool.ExecuteTotal, 1)
		switch t() {
		case normal:
			atomic.AddInt32(&pool.active, -1)
		case reduce:
			if pool.active > int32(pool.Min) {
				break Fun
			}
		}
	}
}

// Execute
//  @Description: 执行任务
//  @receiver pool
//  @param t
//  @return error
//
func (pool *GorotinesPool) Execute(t func()) error {
	if pool.status {
		select {
		case pool.tasks <- func() taskType {
			t()
			return normal
		}:
			atomic.AddInt32(&pool.ReceiveTotal, 1)
			return nil
		case <-time.After(pool.addTimeout):
			return errors.New("add tasks timeout")
		}
	} else {
		return errors.New("pools is down")
	}
}

// Wait
//  @Description: 结束等待任务完成
//  @receiver pool
//
func (pool *GorotinesPool) Wait() {
	pool.status = false
Fun:
	for {
		if len(pool.tasks) == 0 || pool.active == 0 {
			break Fun
		}
		ftool.Sleep(1000)
	}
	defer close(pool.tasks)
	log.Printf("recieve: %d,execute: %d", pool.ReceiveTotal, pool.ExecuteTotal)
}

// AddWorker
//  @Description: 添加worker,协程数加1
//  @receiver pool
//
func (pool *GorotinesPool) AddWorker() {
	atomic.AddInt32(&pool.active, 1)
	go pool.worker()
}

// ReduceWorker
//  @Description: 减少worker,协程数减1
//  @receiver pool
//
func (pool *GorotinesPool) ReduceWorker() {
	atomic.AddInt32(&pool.active, -1)
	pool.tasks <- func() taskType {
		return reduce
	}
}

// balance
//  @Description: 平衡活跃协程数
//  @receiver pool
//
func (pool *GorotinesPool) balance() {
	if pool.status {
		if len(pool.tasks) > 0 && pool.active < int32(pool.Max) {
			pool.AddWorker()
		}
		if len(pool.tasks) == 0 && pool.active > int32(pool.Min) {
			pool.ReduceWorker()
		}
	}
}

// ExecuteQps
//  @Description: 执行任务固定次数
//  @receiver pool
//  @param t
//  @param qps
//
func (pool *GorotinesPool) ExecuteQps(t func(), qps int) {
	mutiple := qps / pool.SingleTimes
	remainder := qps % pool.SingleTimes
	for i := 0; i < pool.SingleTimes; i++ {
		pool.Execute(func() {
			for i := 0; i < mutiple; i++ {
				t()
			}
		})
	}
	pool.Execute(func() {
		for i := 0; i < remainder; i++ {
			t()
		}
	})
}

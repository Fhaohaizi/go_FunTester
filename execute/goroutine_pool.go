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
	tasks        chan func()
	status       bool
	Active       int32
	ExecuteTotal int32
	addTimeout   time.Duration
	MaxIdle      time.Duration
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
func GetPool(max, min, maxWaitTask, timeout, maxIdle int) *GorotinesPool {
	p := &GorotinesPool{
		Max:          max,
		Min:          min,
		tasks:        make(chan func(), maxWaitTask),
		status:       true,
		Active:       0,
		ExecuteTotal: 0,
		addTimeout:   time.Duration(timeout) * time.Second,
		MaxIdle:      time.Duration(maxIdle) * time.Second,
	}
	for i := 0; i < min; i++ {
		p.AddWorker()
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
	for {
		select {
		case t := <-pool.tasks:
			atomic.AddInt32(&pool.ExecuteTotal, 1)
			t()
		case <-time.After(pool.MaxIdle):
			if pool.Active > int32(pool.Min) {
				atomic.AddInt32(&pool.Active, -1)
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
		case pool.tasks <- func() {
			t()
		}:
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
		if len(pool.tasks) == 0 || pool.Active == 0 {
			break Fun
		}
		ftool.Sleep(1000)
	}
	defer close(pool.tasks)
	log.Printf("execute: %d", pool.ExecuteTotal)
}

// AddWorker
//  @Description: 添加worker,协程数加1
//  @receiver pool
//
func (pool *GorotinesPool) AddWorker() {
	atomic.AddInt32(&pool.Active, 1)
	go pool.worker()
}

// balance
//  @Description: 平衡活跃协程数
//  @receiver pool
//
func (pool *GorotinesPool) balance() {
	if pool.status {
		if len(pool.tasks) > 0 && pool.Active < int32(pool.Max) {
			for i := 0; i < len(pool.tasks); i++ {
				if int(pool.Active) < pool.Max {
					pool.AddWorker()
				}
			}
		}
	}
}

// ExecuteQps
//  @Description: 执行任务固定次数
//  @receiver pool 线程池
//  @param t 任务
//  @param times 运行次数
//
func (pool *GorotinesPool) ExecuteQps(t func(), times int) {
	for i := 0; i < times; i++ {
		atomic.AddInt32(&pool.ExecuteTotal, 1)
		pool.Execute(func() {
			t()
		})
	}
}

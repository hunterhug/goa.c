package main

import (
	"fmt"
	"sync"
	"time"
)

type Money struct {
	lock   sync.Mutex // 锁
	amount int64
}

// 加钱
func (m *Money) Add(i int64) {
	// 加锁
	m.lock.Lock()

	// 在该函数结束后执行
	defer m.lock.Unlock()
	m.amount = m.amount + i
}

// 减钱
func (m *Money) Minute(i int64) {
	// 加锁
	m.lock.Lock()

	// 在该函数结束后执行
	defer m.lock.Unlock()

	// 钱足才能减
	if m.amount >= i {
		m.amount = m.amount - i
	}
}

// 查看还有多少钱
func (m *Money) Get() int64 {
	return m.amount
}

func main() {
	m := new(Money)
	m.Add(10000)

	for i := 0; i < 1000; i++ {
		go func() {
			time.Sleep(500 * time.Millisecond)
			m.Minute(5)
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println(m.Get())

}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"

)

// // Counter 计数器结构体
// type Counter struct {
// 	// mu    sync.Mutex
// 	mu sync.Mutex
// 	value int
// }

// // Increment 递增计数器
// func (c *Counter) Increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.value++
// }

// // GetValue 获取计数器值
// func (c *Counter) GetValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.value
// }

// func main() {
// 	// 创建计数器实例
// 	counter := &Counter{}

// 	// 创建等待组，用于等待所有goroutine完成
// 	var wg sync.WaitGroup

// 	// 启动10个goroutine
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(goroutineID int) {
// 			defer wg.Done()
// 			// 每个goroutine执行1000次递增操作
// 			for j := 0; j < 1000; j++ {
// 				counter.Increment()
// 			}
// 			fmt.Printf("Goroutine %d 完成\n", goroutineID)
// 		}(i)
// 	}

// 	// 等待所有goroutine完成
// 	wg.Wait()

// 	// 输出最终结果
// 	fmt.Printf("期望值: %d\n", 10*1000)
// 	fmt.Printf("实际值: %d\n", counter.GetValue())
// }

type AtomicCounter struct {
	value int64
}

// Increment 递增计数器，不用加锁
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// GetValue 获取计数器值，不用加锁
func (c *AtomicCounter) GetValue() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() { 

	counter := &AtomicCounter{}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
			fmt.Printf("Goroutine %d 完成\n", goroutineID)
		}(i)
	}

	wg.Wait()

	fmt.Printf("期望值: %d\n", 10*1000)
	fmt.Printf("实际值: %d\n", counter.GetValue())

}
package help

import (
    "fmt"
    "sync"
)

// 共享资源结构体
// Mutex其实就是互斥锁
type SafeCounter struct {
    mu      sync.Mutex
    counter map[string]int
}

// 增加计数器的方法
func (sc *SafeCounter) Increment(key string, wg *sync.WaitGroup) {
    defer wg.Done()

    sc.mu.Lock()         // 上锁
    sc.counter[key]++    // 修改共享资源
    sc.mu.Unlock()       // 解锁
}

// 获取计数器的方法
func (sc *SafeCounter) Value(key string) int {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    return sc.counter[key]
}

func Test_wg() {
    sc := SafeCounter{
        counter: make(map[string]int),
    }

    var wg sync.WaitGroup
    keys := []string{"apple", "banana", "cherry"}
    numIncrements := 1000

    // 启动多个 Goroutine 增加计数器
    for _, key := range keys {
        for i := 0; i < numIncrements; i++ {
            wg.Add(1)
            go sc.Increment(key, &wg)
        }
    }

    wg.Wait() // 等待所有 Goroutine 完成

    // 打印结果
    for _, key := range keys {
        fmt.Printf("%s: %d\n", key, sc.Value(key))
    }
}

package help

import (
    "fmt"
    "sync"
)

// 共享计数器
var counter int
var mu sync.Mutex // 定义一个全局的 Mutex

// 增加计数器的函数
func increment(wg *sync.WaitGroup) {
    defer wg.Done()

    mu.Lock()         // 上锁
    counter++         // 修改共享资源
    mu.Unlock()       // 解锁
}

func Test_mutex() {
    var wg sync.WaitGroup
    numGoroutines := 1000

    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait() // 等待所有 Goroutine 完成
    fmt.Printf("Final counter value: %d\n", counter)
}

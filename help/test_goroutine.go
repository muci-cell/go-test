package help

import (
    "fmt"
    "sync"
    "time"
)

// worker 函数模拟一个耗时的任务
// ch chan<- string 是 Go 语言中 Channel 的一种声明方式，具体含义如下：
// ch: 变量名，表示一个 Channel。
// chan<-: 这是 Channel 的方向标识符。它表示这个 Channel 是一个只发送（send-only）Channel。
// string: 这个 Channel 的数据类型，表示它只能发送字符串类型的数据。
func worker(id int, wg *sync.WaitGroup, ch chan<- string) {
    defer wg.Done() // 标记当前 Goroutine 完成
    fmt.Printf("Worker %d starting\n", id)

    // 模拟耗时工作
    time.Sleep(time.Second * 2)

    // 将结果发送到 Channel
    result := fmt.Sprintf("Worker %d finished", id)
    ch <- result
}

func Test_goroutine() {
    var wg sync.WaitGroup
    ch := make(chan string) // 创建一个 Channel 用于接收结果

    // 启动多个 Goroutines
    for i := 1; i <= 5; i++ {
        wg.Add(1) // 增加等待的 Goroutine 数量
        go worker(i, &wg, ch) // 启动 Goroutine
    }

    // 启动一个 Goroutine 用于等待所有工作完成后关闭 Channel
    go func() {
        wg.Wait() // 等待所有 Goroutines 完成
        close(ch) // 关闭 Channel
    }()

    // 从 Channel 接收结果
    for result := range ch {
        fmt.Println(result)
    }

    fmt.Println("All workers completed.")
}



package help

import (
    "fmt"
    "sync"
    "time"
)

// 共享队列结构体
type SafeQueue struct {
    mu    sync.Mutex
    queue []int
}

// 入队方法
func (sq *SafeQueue) Enqueue(item int) {
    sq.mu.Lock()
    defer sq.mu.Unlock()
    sq.queue = append(sq.queue, item)
}

// 出队方法
func (sq *SafeQueue) Dequeue() (int, bool) {
    sq.mu.Lock()
    defer sq.mu.Unlock()
    if len(sq.queue) == 0 {
        return 0, false
    }
    item := sq.queue[0]
    sq.queue = sq.queue[1:]
    return item, true
}

func producer(id int, sq *SafeQueue, wg *sync.WaitGroup, items int) {
    defer wg.Done()
    for i := 0; i < items; i++ {
        item := id*1000 + i
        sq.Enqueue(item)
        fmt.Printf("Producer %d enqueued %d\n", id, item)
        time.Sleep(time.Millisecond * 10) // 模拟生产延迟
    }
}

func consumer(id int, sq *SafeQueue, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        item, ok := sq.Dequeue()
        if !ok {
            // 队列为空，等待一段时间后重试
            time.Sleep(time.Millisecond * 50)
            continue
        }
        fmt.Printf("Consumer %d dequeued %d\n", id, item)
        time.Sleep(time.Millisecond * 20) // 模拟消费延迟
    }
}

func Test_mu_wg() {
    var wg sync.WaitGroup
    sq := SafeQueue{
        queue: make([]int, 0),
    }

    numProducers := 3
    numConsumers := 2
    itemsPerProducer := 5

    // 启动生产者
    for p := 1; p <= numProducers; p++ {
        wg.Add(1)
        go producer(p, &sq, &wg, itemsPerProducer)
    }

    // 启动消费者
    for c := 1; c <= numConsumers; c++ {
        wg.Add(1)
        go consumer(c, &sq, &wg)
    }

    // 等待生产者完成
    wg.Wait()

    // 由于消费者是无限循环，主程序需要手动结束
    fmt.Println("All producers have finished. Exiting.")
}

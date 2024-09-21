package help

import (
    "fmt"
)

// 只发送的函数
func sendData(ch chan<- string) {
    ch <- "Hello, World!" // 发送数据到 Channel
}

// 只接收的函数
func receiveData(ch <-chan string) {
    msg := <-ch // 从 Channel 接收数据
    fmt.Println(msg)
}

func Test_channel() {
    ch := make(chan string)

    go sendData(ch)     // 启动 Goroutine 发送数据
    go receiveData(ch)  // 启动 Goroutine 接收数据

    // 为了让主程序等待 Goroutines 完成
    var input string
    fmt.Scanln(&input) // 等待用户输入
}

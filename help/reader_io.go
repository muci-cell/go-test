package help

import (
    "fmt"
    "io"
    "os"
)

func Test_reader_io() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	// defer 关键字用于延迟执行某个函数，直到包含它的函数返回为止。
	// 结合 file.Close()，这通常用于资源管理，确保文件在不再需要时被正确关闭，避免资源泄漏。
    defer file.Close()

    buffer := make([]byte, 100)
    n, err := file.Read(buffer)
    if err != nil && err != io.EOF {
        fmt.Println("Read Error:", err)
        return
    }
    fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
}

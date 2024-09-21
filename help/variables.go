// variables.go
package help

import "fmt"

func Variables() {
    // 显式声明
    var a int = 10
    var b string = "Hello, Go!"

    // 类型推断
    var c = 3.14

    // 短变量声明（仅在函数内部）
    d := true

    fmt.Println("a:", a)
    fmt.Println("b:", b)
    fmt.Println("c:", c)
    fmt.Println("d:", d)
}

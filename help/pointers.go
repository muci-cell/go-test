// pointers.go
package help

import "fmt"

func Pointers() {
    var a int = 10
    var p *int = &a // p 是指向 a 的指针

    fmt.Println("a 的值:", a)
    fmt.Println("p 的值（a 的地址）:", p)
    fmt.Println("*p 的值（通过指针访问 a 的值）:", *p)

    // 通过指针修改 a 的值
    *p = 20
    fmt.Println("修改后 a 的值:", a)
}

// control_structures.go
package help

import "fmt"

func Control_structures() {
    // if-else 结构
    number := 10
    if number%2 == 0 {
        fmt.Println(number, "是偶数")
    } else {
        fmt.Println(number, "是奇数")
    }

    // for 循环
    fmt.Println("for 循环示例:")
    for i := 1; i <= 5; i++ {
        fmt.Println("i =", i)
    }

    // while 风格的 for 循环
    fmt.Println("while 风格的 for 循环示例:")
    j := 1
    for j <= 3 {
        fmt.Println("j =", j)
        j++
    }

    // switch 结构
    day := "星期二"
    switch day {
    case "星期一":
        fmt.Println("今天是星期一")
    case "星期二":
        fmt.Println("今天是星期二")
    default:
        fmt.Println("今天是其他日子")
    }
}

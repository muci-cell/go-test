// functions.go
package help

import "fmt"

// 无返回值函数
func greet(name string) {
    fmt.Println("Hello,", name)
}

// 有返回值函数
func add(a int, b int) int {
    return a + b
}

// 多返回值函数
func swap(a, b string) (string, string) {
    return b, a
}

func Functions() {
    greet("Alice")

    sum := add(5, 7)
    fmt.Println("5 + 7 =", sum)

    x, y := swap("first", "second")
    fmt.Println("Swapped:", x, y)
}

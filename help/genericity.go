package help

import (
    "fmt"
)

func Map[T any, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

func Test_genericity() {
    ints := []int{1, 2, 3, 4}
    strs := Map(ints, func(i int) string {
        return fmt.Sprintf("Number %d", i)
    })
    fmt.Println(strs) // 输出: [Number 1 Number 2 Number 3 Number 4]
}

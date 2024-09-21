package help

import (
    "fmt"
)

func Filter[T any](s []T, predicate func(T) bool) []T {
    var result []T
    for _, v := range s {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

func Test_genericity_filter() {
    nums := []int{1, 2, 3, 4, 5, 6}
    evens := Filter(nums, func(n int) bool {
        return n%2 == 0
    })
    fmt.Println(evens) // 输出: [2 4 6]
}

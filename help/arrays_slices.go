// arrays_slices.go
package help

import "fmt"

func Arrays_slices() {
    // 数组
    var arr [5]int
    arr[0] = 10
    arr[1] = 20
    arr[2] = 30
    arr[3] = 40
    arr[4] = 50
    fmt.Println("数组:", arr)

    // 切片
    slice := []string{"Go", "Python", "Java"}
    fmt.Println("切片:", slice)

    // 添加元素
    slice = append(slice, "JavaScript")
    fmt.Println("添加元素后的切片:", slice)

    // 切片的长度和容量
    fmt.Println("切片长度:", len(slice))
    fmt.Println("切片容量:", cap(slice))

    // 切片的子切片
    subSlice := slice[1:3]
    fmt.Println("子切片:", subSlice)
}

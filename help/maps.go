// maps.go
package help

import "fmt"

func Maps() {
    // 创建一个空映射
    colors := make(map[string]string)

    // 添加键值对
    colors["red"] = "#FF0000"
    colors["green"] = "#00FF00"
    colors["blue"] = "#0000FF"

    // 访问映射中的值
    fmt.Println("Red 的十六进制代码是", colors["red"])

    // 删除键值对
    delete(colors, "green")
    fmt.Println("删除 green 后的映射:", colors)

    // 遍历映射
    fmt.Println("遍历映射:")
    for key, value := range colors {
        fmt.Println(key, ":", value)
    }

    // 检查键是否存在
    value, exists := colors["blue"]
    if exists {
        fmt.Println("blue 存在，值为", value)
    } else {
        fmt.Println("blue 不存在")
    }
}

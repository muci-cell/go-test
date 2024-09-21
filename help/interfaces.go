// interfaces.go
package help

import "fmt"

// 定义接口
type Animal interface {
    Speak() string
}

// 定义结构体 Dog
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return d.Name + " says Woof!"
}

// 定义结构体 Cat
type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return c.Name + " says Meow!"
}

func Interfaces() {
    var animal Animal

    animal = Dog{Name: "Buddy"}
    fmt.Println(animal.Speak())

    animal = Cat{Name: "Whiskers"}
    fmt.Println(animal.Speak())

    // 切片存储不同类型的 Animal
    animals := []Animal{
        Dog{Name: "Rex"},
        Cat{Name: "Mittens"},
    }

    fmt.Println("动物们说:")
    for _, a := range animals {
        fmt.Println(a.Speak())
    }
}

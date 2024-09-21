// structs.go
package help

import "fmt"

// 定义一个结构体
type Person struct {
    Name string
    Age  int
}

func Structs() {
    // 创建结构体实例
    var person1 Person
    person1.Name = "Alice"
    person1.Age = 30

    // 使用字面量创建结构体
    person2 := Person{Name: "Bob", Age: 25}

    // 指针方式创建结构体
    person3 := &Person{Name: "Charlie", Age: 28}

    fmt.Println("Person1:", person1)
    fmt.Println("Person2:", person2)
    fmt.Println("Person3:", *person3)

    // 结构体嵌套
    type Address struct {
        City    string
        ZipCode string
    }

    type Employee struct {
        Person
        Address
        Position string
    }

    employee := Employee{
        Person: Person{
            Name: "David",
            Age:  35,
        },
        Address: Address{
            City:    "Shanghai",
            ZipCode: "200000",
        },
        Position: "Engineer",
    }

    fmt.Println("Employee:", employee)
}

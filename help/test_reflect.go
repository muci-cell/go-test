package help

import (
    "fmt"
    "reflect"
)

type Person2 struct {
    Name string
    Age  int
}

type Greeter struct {
    Name string
}

// 保持方法集一致性：如果您的类型有方法需要使用指针接收者，建议所有的方法都使用指针接收者，以避免方法集的不一致性
// 不修改原始结构体：由于接收者是副本，对接收者的任何修改都不会影响原始结构体。
func (g Greeter) Greet() {
    fmt.Printf("Hello, %s!\n", g.Name)
}

func Test_reflect() {
    p := Person2{"Alice", 30}
    t := reflect.TypeOf(p)

    // 确保是结构体类型
    if t.Kind() == reflect.Struct {
        for i := 0; i < t.NumField(); i++ {
            field := t.Field(i)
            fmt.Printf("Field %d: %s %s\n", i, field.Name, field.Type)
        }
    }
	// 要通过反射修改值，必须确保 reflect.Value 是可设置的。这通常意味着需要传递指针。
	x := 10
    v := reflect.ValueOf(&x).Elem() // 获取指针指向的值

    fmt.Println("Before:", x)
    v.SetInt(20)
    fmt.Println("After:", x)

	g := Greeter{"Bob"}
    z := reflect.ValueOf(g)

    method := z.MethodByName("Greet")
    if method.IsValid() {
        method.Call(nil) // 调用方法，无参数
    }
}

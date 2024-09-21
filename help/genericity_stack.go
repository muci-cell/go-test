package help

import (
    "fmt"
)
/*
泛型应用
4.1 数据结构库
实现通用的数据结构，如列表、栈、队列、树、图等，不需要为每种数据类型编写单独的实现。

4.2 通用算法
实现通用的算法，如排序、搜索、过滤、映射等，适用于不同的数据类型。

4.3 序列化与反序列化
编写通用的序列化库，可以处理多种数据类型，而无需为每种类型编写特定的序列化逻辑。

注意：
6.4 提供类型约束的文档说明
在定义泛型函数或类型时，提供清晰的文档说明类型参数的约束和预期行为，帮助其他开发者理解和使用。

6.5 避免与非泛型代码混用
尽量避免在泛型代码中混用非泛型代码，保持代码的一致性和类型安全。
*/

//泛型集如下
type Ordered interface {
    ~int | ~float64 | ~string
}

func Min[T Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// 泛型方法如下
type Container[T any] struct {
    value T
}

func (c *Container[T]) Set(val T) {
    c.value = val
}

// 小型结构体：对于非常小且复制成本低的结构体（例如只有几个字段），值接收者的性能差异可以忽略不计。
// 这里考虑到 是 只读  所以弄个副本也没事
func (c Container[T]) Get() T {
    return c.value
}

// 使用泛型实现栈数据结构
type Stack[T any] struct {
    elements []T
}

func (s *Stack[T]) Push(elem T) {
    s.elements = append(s.elements, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }
    index := len(s.elements) - 1
    elem := s.elements[index]
    s.elements = s.elements[:index]
    return elem, true
}

func Test_genericity_stack() {
    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    elem, ok := intStack.Pop()
    if ok {
        fmt.Println(elem) // 输出: 2
    }

    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")
    s, ok := stringStack.Pop()
    if ok {
        fmt.Println(s) // 输出: world
    }
}

package help

import "fmt"

// 定义 sentinel error
var ErrInvalidInput = fmt.Errorf("invalid input")

// 定义自定义错误类型
type OperationError struct {
    Op  string
    Err error
}
// 如果你希望 *OperationError 实现 error 接口，而不仅仅是 OperationError，那么你需要使用指针接收者。
// %s：字符串
// %d：十进制整数
// %v：默认格式，适用于多种类型
// %f：浮点数
// %t：布尔值
// %+v：结构体字段名和值
func (e *OperationError) Error() string {
    return fmt.Sprintf("operation %s failed: %v", e.Op, e.Err)
}

func (e *OperationError) Unwrap() error {
    return e.Err
}

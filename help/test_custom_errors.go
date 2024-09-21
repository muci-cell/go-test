package help

import (
    "errors"
    "fmt"
)
// 模拟一个函数，可能返回 sentinel error
func ValidateInput(input int) error {
    if input <= 0 {
        return ErrInvalidInput
    }
    return nil
}

// 模拟一个函数，返回自定义错误类型
func PerformOperation(input int) error {
    err := ValidateInput(input)
    if err != nil {
        return &OperationError{
            Op:  "PerformOperation",
            Err: err,
        }
    }
    // 执行操作...
    return nil
}

func Test_errors() {
    err := PerformOperation(-1)
    if err != nil {
        fmt.Println("Error occurred:", err)
        
        // 使用 errors.Is 检查 sentinel error
        if errors.Is(err, ErrInvalidInput) {
            fmt.Println("The input provided is invalid.")
        }
        
        // 使用 errors.As 进行类型断言
        var opErr *OperationError
        if errors.As(err, &opErr) {
            fmt.Printf("Operation: %s, Underlying error: %v\n", opErr.Op, opErr.Err)
        }
    } else {
        fmt.Println("Operation completed successfully.")
    }
}

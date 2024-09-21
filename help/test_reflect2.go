package help

import (
    "fmt"
    "reflect"
)

type User struct {
    Name  string
    Email string
    Age   int
}

type Admin struct {
    Name    string
    Email   string
    Privledges []string
}

func CopyFields(src interface{}, dst interface{}) error {
    srcVal := reflect.ValueOf(src)
    dstVal := reflect.ValueOf(dst)

    // 确保 dst 是指针
    if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
        return fmt.Errorf("dst must be a pointer to a struct")
    }

    srcVal = srcVal.Elem()
    dstVal = dstVal.Elem()

    srcType := srcVal.Type()

    for i := 0; i < srcVal.NumField(); i++ {
        field := srcType.Field(i)
        dstField := dstVal.FieldByName(field.Name)
        if dstField.IsValid() && dstField.CanSet() && dstField.Type() == srcVal.Field(i).Type() {
            dstField.Set(srcVal.Field(i))
        }
    }

    return nil
}

// 不使用反射的坏处（反射的性能开销 相对直接赋值大）：
// 重复代码：如果有多个结构体需要类似的字段复制，可能会导致代码重复。
// 不够灵活：需要在编译时明确指定字段，无法处理动态类型或未知结构体。
func Test_reflect2() {
    user := User{
        Name:  "Alice",
        Email: "alice@example.com",
        Age:   30,
    }

    admin := Admin{
        Privledges: []string{"read", "write"},
    }

    err := CopyFields(&user, &admin)
    if err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Printf("Admin: %+v\n", admin)
}

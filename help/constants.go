// constants.go
package help

import "fmt"

const Pi = 3.14159

func Constants() {
    const Greeting = "Hello, Constants!"
    fmt.Println(Greeting)
    fmt.Println("Pi =", Pi)
}

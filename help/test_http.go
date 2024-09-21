package help

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

//自定义 HTTP 客户端
var client = &http.Client{
    Timeout: time.Second * 10,
}

// 可以自定义中间件
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Request received:", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// 处理表单数据
func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        r.ParseForm()
        name := r.FormValue("name")
        fmt.Fprintf(w, "Hello, %s!", name)
    }
}
// 关于json编码
type Post struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    post := Post{ID: 1, Title: "Hello", Body: "World"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(post)
}
// 监听到端口的请求后的处理函数
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func Test_http() {
	http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)

    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading body:", err)
        return
    }

    fmt.Println("Response:", string(body))
}

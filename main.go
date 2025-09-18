package main

import (
	"fmt"
	"net/http"
)
func main() {
	// 定义一个处理函数
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, KVStore learner!")
    })

    fmt.Println("Server running at http://localhost:8080/hello")
    // 启动 HTTP 服务，监听 8080 端口
    http.ListenAndServe(":8080", nil)
}
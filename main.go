package main

import (
	"fmt"
	"net/http"
)

// 定义一个全局 map 作为简单 KV Store
var store = make(map[string]string)

func main() {
	// 设置 key-value
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		if key == "" || value == "" {
			fmt.Fprintf(w, "Missing key or value\n")
			return
		}
		store[key] = value
		fmt.Fprintf(w, "Set %s = %s\n", key, value)
	})

	// 获取 key
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			fmt.Fprintf(w, "Missing key\n")
			return
		}
		value, ok := store[key]
		if !ok {
			fmt.Fprintf(w, "Key not found\n")
			return
		}
		fmt.Fprintf(w, "Value = %s\n", value)
	})

	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

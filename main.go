package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type KVRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 定义一个全局 map 作为简单 KV Store
var store = make(map[string]string)

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
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
	http.HandleFunc("/getjson", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			respondJSON(w, map[string]string{"error": "Missing key"})
			return
		}
		value, ok := store[key]
		if !ok {
			respondJSON(w, map[string]string{"error": "Key not found"})
			return
		}
		respondJSON(w, map[string]string{"key": key, "value": value})
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
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, "Delete %s = %s\n", key, value)
		delete(store, key)
		respondJSON(w, map[string]string{"message": "Key deleted successfully"})
	})
	http.HandleFunc("/setjson", func(w http.ResponseWriter, r *http.Request) {
		var req KVRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Key == "" || req.Value == "" {
			respondJSON(w, map[string]string{"error": "Invalid JSON"})
			return
		}
		store[req.Key] = req.Value
		respondJSON(w, map[string]string{"message": "Set successful"})
	})

	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

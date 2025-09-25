package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type KVRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 全局 KV Store
var store = make(map[string]string)

// 统一响应格式
func respondJSON(w http.ResponseWriter, status string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": status,
		"data":   data,
	})
}

// 提取 URL 路径中的 key，例如 /kv/name -> name
func getKeyFromPath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) >= 3 {
		return parts[2]
	}
	return ""
}

func main() {
	// 创建（POST /kv）
	http.HandleFunc("/kv", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}

		var req KVRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Key == "" || req.Value == "" {
			respondJSON(w, "error", "Invalid JSON")
			return
		}

		store[req.Key] = req.Value
		respondJSON(w, "success", map[string]string{"key": req.Key, "value": req.Value})
	})

	// 读取 / 更新 / 删除（GET, PUT, DELETE /kv/{key}）
	http.HandleFunc("/kv/", func(w http.ResponseWriter, r *http.Request) {
		key := getKeyFromPath(r.URL.Path)
		if key == "" {
			respondJSON(w, "error", "Missing key in URL")
			return
		}

		switch r.Method {
		case http.MethodGet: // 读取
			value, ok := store[key]
			if !ok {
				respondJSON(w, "error", "Key not found")
				return
			}
			respondJSON(w, "success", map[string]string{"key": key, "value": value})

		case http.MethodPut: // 更新
			var req KVRequest
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil || req.Value == "" {
				respondJSON(w, "error", "Invalid JSON")
				return
			}

			_, exists := store[key]
			if !exists {
				respondJSON(w, "error", "Key not found")
				return
			}

			store[key] = req.Value
			respondJSON(w, "success", map[string]string{"key": key, "value": req.Value})

		case http.MethodDelete: // 删除
			_, ok := store[key]
			if !ok {
				respondJSON(w, "error", "Key not found")
				return
			}
			delete(store, key)
			respondJSON(w, "success", fmt.Sprintf("Key %s deleted", key))

		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

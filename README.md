# 🚀 C++ & Go 学习日记

你好，这里是我的 **C++ & Go 学习之路**！  
我会在这个仓库里一步一步留下学习痕迹 📚，让未来的自己看到现在的努力。  

---

## 📅 学习进度

### ✅ Day 1
- 实现了一个最小版 **C++ KVStore**（键值存储）  
  - 支持 `Set / Get / Delete`  
  - 使用 `unordered_map` 存放数据  
  - 体会到 C++ STL 的方便  
- [代码入口](./kvstore.cpp)

---

### ✅ Day 2
- 学习了 Go 的基础语法（变量、函数、map、struct）  
- 写了人生第一个 **Go HTTP 服务** 🎉  
  - 访问 [http://localhost:8080/hello](http://localhost:8080/hello) 会返回：  
    ```
    Hello, KVStore learner!
    ```
- [代码入口](./main.go)
- ### ✅ Day 3
一个简单的 **键值存储（KV Store）Web 服务**，用于练习 Go 语言和 HTTP 编程。  
这是我 C++ KVStore 项目的延伸版本，把内存中的 KV Store 暴露为 Web API，支持 GET/SET 操作。

## 功能

- **设置键值对**  
GET /set?key=<key>&value=<value>

复制代码
示例：
http://localhost:8080/set?key=name&value=Alice

复制代码
返回：
Set name = Alice

markdown
复制代码

- **获取键值对**  
GET /get?key=<key>

复制代码
示例：
http://localhost:8080/get?key=name

复制代码
返回：
Value = Alice

bash
复制代码

## 技术栈

- Go 语言（Golang）  
- 标准库 `net/http`  

## 使用方法

1. 克隆仓库：
 ```bash
 git clone https://github.com/YourName/kvstore-go.git
 cd kvstore-go
运行服务器：

bash
复制代码
go run main.go
打开浏览器或用 curl 访问：

bash
复制代码
curl "http://localhost:8080/set?key=name&value=Alice"
curl "http://localhost:8080/get?key=name"
今日收获
学会 Go 的 匿名函数及如何作为回调函数传给 http.HandleFunc。

理解 HTTP URL 的组成，能从 URL 获取查询参数。

将一个简单的内存 KV Store 暴露为 Web 服务接口。

yaml
复制代码

---
---

## 🌟 小感想
- C++：感觉很像“沉重的铁锤”，功能强大，但需要小心使用。  
- Go：感觉就像“轻便的小刀”，上手简单，很快就能切开问题。  
- 今天在浏览器里看到 `"Hello, KVStore learner!"` 的那一刻，超级有成就感！🚀

---

## 🏃 目标
- 把 C++ 写的 KVStore 和 Go 的 Web 服务连接起来。  
- 做一个能存取数据的迷你数据库服务器！  

---

_持续更新中…_

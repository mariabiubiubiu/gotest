package main

import (
	"fmt"
	"net/http"
)

//③处理请求，返回结果
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	//①路由注册
	http.HandleFunc("/", Hello)
	//②服务监听
	http.ListenAndServe(":8080", nil)
}
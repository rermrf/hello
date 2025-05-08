package main

import (
	"fmt"
	"net/http"
)

func main() {
	//使用 http 包创建一个6060端口的http服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world! 测试CI/CD")
	})
	http.ListenAndServe(":6060", nil)
}

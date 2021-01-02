package main // 要求这里报名必须是 main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()                                // 解析参数，默认是不会解析的
	fmt.Println("request parameters : ", request.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path : ", request.URL.Path)
	fmt.Println("scheme : ", request.URL.Scheme)
	fmt.Println("request parameter : url_long = ", request.Form["url_long"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ","))
	}
	fmt.Fprintf(responseWriter, "你好，浏览器!") // 这个写入到 responseWriter 的是输出到客户端的
}

/**
访问测试 :
1. 启动本程序
2. 访问地址 http://localhost:9090/?a=123&a=32&url_long=10000
*/
func main() {
	http.HandleFunc("/", sayhelloName)       // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

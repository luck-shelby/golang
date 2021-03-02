/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  简单web服务器
 * @Date: 2020/12/15 16:13
 */
package main

import (
	"fmt"
	"net/http"
)

func handler(resp http.ResponseWriter, req *http.Request)  {
	// _, _ = resp.Write([]byte("hello word"))
	fmt.Println("hello consul")
}
func main() {
	// 注册回调函数, 用于服务器被访问时,自动被调用 . 回调函数 作为另一个函数的参数的函数，叫回调函数
	http.HandleFunc("/health",handler)  // 每一个请求在都有一个对应的 goroutine 去处理
	// 绑定服务器监听地址
	err := http.ListenAndServe("172.21.165.61:8099",nil)
	if err != nil {
		return
	}
}
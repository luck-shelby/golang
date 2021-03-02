/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  go协程
 * @Date: 2020/12/3 18:20
 */
package main

import (
	"fmt"
	"time"
)

/*
	go并发: 不需要在程序中设计并发, go语言天生支持并发
		在go中实现并发只要的两种手段: channel和 goroutine

	goroutine: 协程 , 通过通信来共享内存
	goroutine特性: 主goroutine执行结果(相当于进程结束), 其他子goroutine随之退出,
*/
func Dance()  {
	for i:=0;i<5;i++{
		fmt.Println("任务函数...")
		time.Sleep(time.Second)
	}
}

func main() {
	go Dance()

	// 在main中默认会有一个主 goroutine
	time.Sleep(time.Second * 5)
	fmt.Println("main...")
}
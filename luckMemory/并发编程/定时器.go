/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  定时器
 * @Date: 2020/12/6 17:21
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	Timer: 定时器
*/

func main9() {
	// 创建一个定时器, 当到达Duration后.系统自动会在返回的Timer结构体成员 C 写入当前时间
	now := time.Now()
	fmt.Println("当前时间:",now)
	myTimer := time.NewTimer(time.Second*2)
	duration := <- myTimer.C  // 读取定时后的时间,并完成一次chan的读操作
	fmt.Println("到达duration后:",duration)
}

// 完成定时的其他方法 time.after(duration)
func main10() {
	now := time.Now()
	fmt.Println("当前时间:",now)
	duration := <- time.After(time.Second*2)
	fmt.Println("到达duration后:",duration)
}

// 停止定时器与重置定时器
func main11() {
	myTimer := time.NewTimer(time.Second*5)
	myTimer.Reset(1) // 重置定时器
	go func() {
		<- myTimer.C
		fmt.Println("子go程完成读取定时后的时间")
	}()
	myTimer.Stop()  // 设置定时器停止
	// 让主go程等待
	time.Sleep(time.Minute)
}

// 周期定时
func main() {
	quit := make(chan bool)
	fmt.Println("当前时间: ",time.Now())
	myTicker := time.NewTicker(time.Second*1)
	count := 0
	go func() {
		for  {
			// 每隔5秒就会运行一次
			nowTime := <- myTicker.C
			fmt.Println("子go程循环读取定时时间：",nowTime)
			if count == 8{
				quit <- true
				runtime.Goexit()
			}
			count ++
		}
	}()
	// 让主go程等待
	<- quit
}
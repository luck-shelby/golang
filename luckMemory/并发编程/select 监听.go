/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  select 监听
 * @Date: 2020/12/6 19:33
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	通过select监听channel中的数据流动,它与switch语法非常类似.
	它的每个case必须是一个IO操作
	注意事项:
		监听的case中,没有满足条件则会阻塞
		监听的case中.有多个满足监听的条件,会任意选择一个执行
		可以使用default来处理所以case都不满足条件的状况,会产生忙轮询(不常用)
		select自身不带循环机制,所以需要借助for循环来循环监听
		break只能跳出当前case,这个与switch中的用法一致
*/

func main12() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 1; i < 11; i++{
			ch <- i
		}
		close(ch)
		quit <- true
		runtime.Goexit()
	}()

	for  {
		select {
		case data := <- ch:
			if data == 0{
				fmt.Println("数据全部读取完成")
			}else {
				fmt.Println("data = ",data)
			}
		case <-quit:
			// break break只是退出当前的case,所以不能使用break
			// runtime.Goexit() 在主go程中使用: fatal error: no goroutines (main called runtime.Goexit) - deadlock!
			return
		}
		fmt.Println("--------------------------")
	}
}

// exercise
func Fibonacci(ch <-chan int,q <-chan bool)  {
	// 如果没有使用for报错: 1 fatal error: all goroutines are asleep - deadlock!
	for  {
		select {
		case data:= <- ch:
			fmt.Print(data ," ")
		case <-q:
			// runtime.Goexit() 在子go程中使用不会报错,
			return
		}
	}
}

func main13() {
	ch := make(chan int)
	quit := make(chan bool)
	go Fibonacci(ch,quit)
	x,y := 1,1
	for i := 0;i<20;i++{
		ch <- x
		x,y = y,x+y
	}
	quit <- true
}

// select 超时处理

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for  {
			select {
			case num := <-ch:
				fmt.Println("num = ",num)
			case <- time.After(time.Second*3):
				fmt.Println("读取数据超时")
				quit <- true
				// return
				goto label // 可以使用goto
			}
		}
	label:
	}()
	for i :=0;i<2;i++{
		ch <- i
		time.Sleep(time.Second*2)
	}
	<-quit
	fmt.Println("-------------------------")
}
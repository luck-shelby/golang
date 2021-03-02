/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  死锁
 * @Date: 2020/12/6 22:32
 */
package main

import "fmt"
/*
	死锁: fatal error: all goroutines are asleep - deadlock!
*/
// 死锁案例一: 单go程死锁
func main14() {
	ch := make(chan int)
	ch <- 100  // 阻塞
	num := <- ch
	fmt.Println(num)
}

// 死锁案例二: go程间channel访问顺序导致死锁,使用channel的一端进行读写,要保证另一端写读操作,同时有机会运行,否则死锁
func main15() {
	ch := make(chan int)
	<- ch    // 阻塞
	go func() {
		ch <- 100
	}()
}
// 案例二改进
func main16() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data:=<- ch
	fmt.Println("data = ",data)
}

// 死锁案例三: 多go程,多channel交叉死锁,是指两个或者两个以上的go程执行运行的过程中互相等待的现象
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for  {
			select {
			case num:= <-ch1:
				ch2 <- num
			}
		}
	}()
	for  {
		select {
		case num:= <-ch2:
			ch1 <- num
		}
	}
}
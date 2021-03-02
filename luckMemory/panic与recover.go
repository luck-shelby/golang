/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  panic与recover
 * @Date: 2020/12/16 21:51
 */
package main

import (
	"fmt"
	"time"
)

/*
	recover是恢复程序panic,让程序继续执行,可以接收panic的异常信息
		它一般用在defer内部
	发生panic后，程序会从调用panic的函数位置或发生panic的地方立即返回，逐层向上执行函数的defer语句，
		然后逐层打印函数调用堆栈，直到被recover捕获或运行到最外层函数。
*/
func catchErr(i int)  {
	defer func() {
		if err := recover();err != nil{
			fmt.Println("catchErr recover",err,i)
		}
	}()
	fmt.Println("子go程,",i)
	panic("出现panic啦")
}
func main100() {
	defer func() {
		if err := recover();err != nil{
			fmt.Println("main recover",err)
		}
	}()
	for i := 0; i <3;i++{
		fmt.Println("main ",i)
		go catchErr(i)
		time.Sleep(time.Second)
	}
}

// 示例二
func handlePanic()  {
	defer handlePanic()
	if err := recover();err != nil{
		fmt.Println("捕获到panic啦")
	}
}
func myFunc2()  {
	fmt.Println("welcome to myFunc2")
	panic("出现panic啦")
}
func myFunc1()  {
	fmt.Println("welcome to func1")
	go myFunc2()
	time.Sleep(time.Second*10)
}
func main() {
	myFunc1()
	fmt.Println("main 函数")
}
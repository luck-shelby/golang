/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  runtime包常用函数
 * @Date: 2020/12/5 17:31
 */
package main

import (
	"fmt"
	"runtime"
)

func main1() {
	/*
		Gosched: 让出CPU执行时间片,让出当前goroutine执行权限,把权限给其他任务运行,在下次在获得cpu时间片后,
				 从该让出的地方继续向下执行
	*/
	go func() {
		for  {
			fmt.Println("test func")
		}
	}()
	for  {
		runtime.Gosched()  // 执行几率变得很小
		fmt.Println("main goroutine ")
	}
}

func test()  {
	defer fmt.Println("C")
	runtime.Goexit()
	defer fmt.Println("D")
}
func main2() {
	/*
		Goexit: 结束当前函数的go程, 会执行Goexit 之前的defer语句
		return: 返回当前函数到调用者那里去,return之前的defer会执行
	*/
	go func() {
		fmt.Println("A")
		test()
		defer fmt.Println("B")
	}()

	for  {
		;
	}
}

func main3() {
	/*
		GOMAXPROCS: 设置当前进程可以并行计算CPU核数的最大值
	*/
	n := runtime.GOMAXPROCS(-1)  // 返回上次最大CPU核数
	fmt.Println(n)
	n = runtime.GOMAXPROCS(-1)
	fmt.Println(n)
}

func main() {
	// 其他
	fmt.Println(runtime.GOROOT())
	_,file,line,ok := runtime.Caller(0)  //  0-当前函数，1-上一层函数
	if ok{
		fmt.Println(file,line)
	}
}
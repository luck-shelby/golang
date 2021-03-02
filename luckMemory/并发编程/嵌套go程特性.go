/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  嵌套go程特性
 * @Date: 2020/12/14 17:14
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func aa()  {
	for  {
		time.Sleep(time.Second)
		fmt.Println("i am aa")
	}
}
func main() {
	// 当在其他函数中创建的go程中,函数退出不会导致它对应的子go程退出
	go func() {
		fmt.Println("一")
		go aa()
		fmt.Println("二")
		return
	}()

	for  {
		runtime.GC()
	}
}

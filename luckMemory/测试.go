/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  测试
 * @Date: 2020/12/14 14:53
 */
package main

import (
	"fmt"
	"time"
)

var c = make(chan bool)

func ab()  {
	go func() {
		for i :=0;i<5;i++{
			if i ==3 {
				fmt.Println(i)
				time.Sleep(time.Second * 5)
				c <- true
				return
			}
		}
	}()
	for {
		select {
		case <-c:
			fmt.Println("哈哈")
			return
		}
	}
}
func main() {
	for  {
		println("1111111")
		time.Sleep(time.Second*2)
		go ab()
	}
}

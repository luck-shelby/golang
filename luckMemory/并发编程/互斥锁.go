/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  互斥锁
 * @Date: 2020/12/7 15:02
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	互斥锁: 访问共享数据之前加锁,锁只有一把
*/
// 除了可以用channel完成同步, 使用传统的互斥锁完成同步示例
var mutex sync.Mutex
func printer1(s string)  {
	mutex.Lock()
	for _,v := range s{
		fmt.Printf("%c",v)
		time.Sleep(time.Second)
	}
	mutex.Unlock()
}
func person3(group *sync.WaitGroup)  {  // 一定要通过指针传值，不然进程会进入死锁状态
	printer1("hello")
	group.Done()
}
func person4(group *sync.WaitGroup)  {
	printer1("world")
	group.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)  // 写多了就会出现死锁
	go person3(&wg)
	go person4(&wg)
	wg.Wait()
}
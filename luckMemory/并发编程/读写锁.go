/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  读写锁
 * @Date: 2020/12/7 15:36
 */
package main
/*
	读写锁: 读的时候共享,写的时候独占, 写的优先级高于读的,锁只有一把(实现方式: 读写锁+全局变量)
*/
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var rwLock sync.RWMutex
var value int   // 定义全局变量模拟共享数据
func read(index int)  {
	for  {
		rwLock.RLock()
		num := value
		fmt.Printf("%dth go程 读取 %d \n",index,num)
		rwLock.RUnlock()
	}
}
func write(index int)  {
	for  {
		num := rand.Intn(1000)
		rwLock.Lock()
		value = num
		fmt.Printf("%dth go程 写入 %d \n",index,num)
		time.Sleep(time.Millisecond*3000)
		rwLock.Unlock()
	}
}

// 读写锁 同步
func main17() {
	rand.Seed(time.Now().UnixNano())
	for i:=1;i<6;i++{
		go read(i)
	}
	for i:=1;i<6;i++{
		go write(i)
	}
	for  {
		;
	}
}

// channel同步, 但是没有读写锁效率高
func read1(c <- chan int,index int)  {
	for  {
		num := <- c
		fmt.Printf("%dth go程 读取 %d \n",index,num)
	}
}
func write1(c chan<- int,index int)  {
	for  {
		num := rand.Intn(1000)
		c <- num
		fmt.Printf("%dth go程 写入 %d \n",index,num)
		time.Sleep(time.Millisecond*3000)
	}
}
func main() {
	ch := make(chan int)
	for i:=1;i<6;i++{
		go read1(ch,i)
	}
	for i:=1;i<6;i++{
		go write1(ch,i)
	}
	for  {
		;
	}
}
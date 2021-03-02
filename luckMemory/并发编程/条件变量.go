/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  条件变量
 * @Date: 2020/12/7 19:22
 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
/*
	条件变量不是锁,但是通常与锁一起使用
	使用步骤:
		1. 判断条件变量
		2. 加锁
		3. 访问公共区
		4. 解锁
		5. 唤醒阻塞在条件变量上的对端
	详细流程:
		1. 创建条件变量 var cond sync.Cond
		2. 指定条件变量的需要的锁 cond.L = new(sync.Mutex)
		3. 加锁  给公共区加锁 cond.L.Lock()
		4. 判断是否达到阻塞条件  注意是要循环判断  for len(ch) == cap(ch)  cond.wait()
			 cond.wait()做了3件事情:
				1. 阻塞
				2. 解锁 (1,2是原子操作)
				3. 重新加锁
		5. 访问公共区 (读，写，输出等)
		6. 解锁  cond.L.Unlock
		7. 唤醒阻塞在条件变量上的对端
*/
func production(out chan <- int,q chan bool)  {
	for i:=1;i<101;i++ {
		num := rand.Intn(1000)
		fmt.Println("生产者: ",num)
		out <- num
	}
	close(out)
	q <- true
}
func consumption(in <- chan int)  {
	for data := range in{
		fmt.Println("消费者获得: ",data)
	}
	fmt.Println("消费者数据读取完成！")
}
func main18() {
	noChan := make(chan int)
	quit := make(chan bool)
	rand.Seed(time.Now().UnixNano())
	go production(noChan,quit)
	consumption(noChan)
	<-quit
}

// 多go程 让打印出来的顺序不会错乱,使用条件变量解决示例
var cond sync.Cond
func production2(out chan <- int,index int)  {
	for  {
		cond.L.Lock() // 先加锁
		// 判断缓冲区是否满,必须是循环判断
		for len(out) == 5{
			cond.Wait()
		}
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("生产者%d th 生产了%d \n: ",index,num)
		// 访问公共区结束,解锁
		cond.L.Unlock()
		//  唤醒阻塞在条件变量上的对端
		cond.Signal()
		time.Sleep(time.Millisecond*300)
	}
}
func consumption2(in <- chan int,index int)  {
	for  {
		cond.L.Lock() // 先加锁
		// 判断缓冲区是否满,必须是循环判断
		for len(in) == 0{
			cond.Wait()
		}
		num := <- in
		fmt.Printf("消费者%d th 消费了%d \n: ",index,num)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond*300)
	}
}
func main19()  {
	ch := make(chan int,5)
	rand.Seed(time.Now().UnixNano())
	// 指定条件变量的需要的锁
	cond.L = new(sync.Mutex)

	for i :=1;i<11;i++ {
		go production2(ch,i)
	}
	for i :=1;i<11;i++ {
		go consumption2(ch,i)
	}
	for  {
		;
	}
}

// 实际使用中，应当遵循先channel后锁的顺序，即channel如果能满足需求，则不要用锁，如果场景比较复杂，channel无法满足，再加上锁来控制。
//	因为channel本身就是先进先出，等同于消息队列
func main888() {
	ch := make(chan int)
	v := 0

	// Consumer
	go func() {
		for {
			fmt.Printf("Consumer: %d\n", <-ch)
		}
	}()

	// Producer
	for {
		v++
		fmt.Printf("Producer: %d\n", v)
		ch <- v
		time.Sleep(time.Second)
	}
}

func main() {
	cond := sync.NewCond(new(sync.Mutex))
	condition := 0

	// Consumer
	go func() {
		for {
			cond.L.Lock()
			for condition == 0 {
				cond.Wait()
			}
			fmt.Printf("Consumer: %d\n", condition)
			condition--
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// Producer
	for {
		time.Sleep(time.Second)
		cond.L.Lock()
		for condition == 3 {
			cond.Wait()
		}
		condition++
		fmt.Printf("Producer: %d\n", condition)
		cond.Signal()
		cond.L.Unlock()
	}
}
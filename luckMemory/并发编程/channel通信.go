/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  channel通信
 * @Date: 2020/12/5 19:00
 */
package main

import (
	"fmt"
	"time"
)

/*
	channel: 是一种数据类型, 类似与队列,先进先出
			 主要解决goroutine同步问题以及协程直接数据共享(数据传递)的问题
	写入端与读端必须同时满足条件才能在管道中进行数据流通,否则就会阻塞
		写入端: chan <-
		读端: <- chan

	无缓冲通道:
		通道容量为0,长度也为0,不能存储数据,所以需要一个读,一个写的go程,它具备同步的能力. 读、写同步、
		应用与两个go程以上
	有缓冲通道:
		通道容量大于0,应用与两个go程以上,一个读,一个写,它的长度(len)表示缓冲区剩余容量
		缓冲区可以进行数据存储,容量上限才会阻塞,具备异步的能力,所以它可以不是同时操作缓冲区

	关闭channel: 写入端主动关闭后,读端就可以判断是否全部读取完成
			     确认数据不在发送了,关闭channel
				 已经关闭的channel不能在写入数据,否则就会panic
				 已经关闭的无缓存channel,继续读取不会报错,读到的是 通道数据类型的零值， 如 int 对应的0
		  		 已经关闭的有缓存channel,如果缓冲区有数据,先读数据,读完数据后,可以继续读取不会报错,读到的也是通道数据类型的零值
		当写的次数不固定时,可以使用关闭channel,用来告诉读端不用在等了。

	单向channel: 默认channel是双向的。
		双向通道: ch := make(chan int)
		单向写通道: var sendCh chan <- int     sendCh = make(chan <- int)
		单向读通道: var recvCh <- chan int     recvCh = make(<- chan int)
		注意: 双向通道可以隐士转换成任意一种单向通道 如  sendCh = ch
			  而单向通道不能转换成双向通道
	channel作为函数参数时,传递的是引用

*/

// 定义一个全局无缓冲通道.让协程之前阻塞操作,从而实现同步,输出时数据不会错乱
var channel = make(chan int)
func printer (s string)  {
	for _,str := range s{
		fmt.Printf("%c",str)
		time.Sleep(time.Second)
	}
}
func person1()  {
	printer("hello")
	channel <- 1
	fmt.Println("写入完成")
}
func person2()  {
	<- channel   // 阻塞
	printer("world")
}
func main4() {
	go person1()
	go person2()
	for  {
		;
	}
}

// 无缓冲通道,同步通信
func main5()  {
	ch := make(chan int)
	go func() {
		for i := 0; i<5;i++{
			ch <- i
		}
	}()
	time.Sleep(time.Second*2)  // 睡眠的时候写就会阻塞
	for i := 0; i<5;i++{
		n := <- ch
		fmt.Println("读完 n=",n)
	}
}

// 有缓冲通道  异步通信
func main6() {
	ch := make(chan int,3) // 存满3个之前,不会阻塞
	fmt.Println(len(ch),cap(ch))

	go func() {
		for i := 0;i < 5;i++{
			ch <- i
			fmt.Println("子go程写入:",i)
		}
	}()

	time.Sleep(time.Second*5) // 睡眠的时候写只要不超过缓存值就不会阻塞
	for i := 0;i < 5;i++{
		num := <- ch
		fmt.Println("主go程读取:",num)
	}
}

// 关闭channel
func main7() {
	ch := make(chan int)
	go func() {
		for i:=1;i<6;i++{
			ch <- i
		}
		close(ch) // 写入端写完数据后,主动关闭channel
		fmt.Println("数据写完")
	}()

	// 主go程读取数据并判断channel是否已经关闭 方式一：
	/*
	for  {
		if c,ok := <- ch;ok{
			fmt.Println("读取到的数据:",c)
		}else {
			break
		}
	}
	*/
	// 主go程读取数据并判断channel是否已经关闭 方式二：使用for range ,不需要判断, 并且ch 不能写成  <- ch
	for data := range ch{
		fmt.Println("读取到的数据:",data)
	}
}

// 单向channel
func main8() {
	var sendCh chan <- int
	sendCh = make(chan <- int)
	sendCh <- 100

	// 通过隐士转换
	ch := make(chan int)
	var recvCh <- chan int
	recvCh = ch
	<- recvCh
}
func send(out chan <- int)  {
	out <- 100
	close(out)
}
func recv(in <- chan int)  {
	for data := range in{
		fmt.Println("recv data := ",data)
	}
}
func main() {
	ch := make(chan int)
	go func() {
		send(ch) // 隐士转换,双向channel转单向写channel
	}()
	recv(ch)
}
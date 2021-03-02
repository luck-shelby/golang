/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  生产者与消费者
 * @Date: 2020/12/6 16:08
 */
package main

import "fmt"

/*
	单向channel最典型的应用: 生产者与消费者模型
		生产者与消费者之间的缓冲区(channel)起到的作用: 有缓冲channel与无缓冲channel都可以作为缓冲区,一个需要时时同步,另一个异步
			1. 解耦
			2. 并发(生产者与消费者处理数据可以不对等,能保持正常通信)
			3. 缓存(生产者与消费者数据处理速度不一致时,暂存数据)

*/
func producer(out chan <- int)  {
	for i := 1; i < 11; i++{
		out <- i * i
	}
	close(out)
}
func consumer(in <- chan int)  {
	for data := range in{
		fmt.Println("消费者读取data: ",data)
	}
}
func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
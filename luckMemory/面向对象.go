/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  接口
 * @Date: 2020/12/16 21:10
 */
package main

import "fmt"

/*
	接口是一组行为规范的定义,结构体把接口所有的方法都重写后,那么改结构体就属于接口类型
	多态: 最常见的就是把接口当作方法的参数
		结构体实现了接口的全部方法后,就认为改结构体实现了接口,这样就可以把结构体赋值给接口变量
	接口断言: 如果 interface{}作为方法的参数就可以接收任意类型,然后就可以判断参数到底是什么类型,就可以使用接口断言
		它具有判断与转换的功能
*/
type Live interface {
	Run()
}
type People struct {
	name string
}
type Animal struct {
	name string
}
func (receiver *People) Run()  {
	// 只要在方法中有* 在结构体赋值给接口变量必须是指针
	fmt.Println(receiver.name,"is running")
}
func (receiver *Animal)Run()  {
	fmt.Println(receiver.name,"is running")
}
func main50() {
	var live Live = &People{"张三"}
	live.Run()
}

// 多态表现
func SomethingRun(live Live)  {
	live.Run()
}
func main60() {
	person := &People{"张三"}
	animal := &Animal{"老虎"}
	SomethingRun(person)
	SomethingRun(animal)
}

// 接口断言
func main() {
	var i interface{} = 456
	val,ok := i.(int)
	if ok{
		fmt.Println(val)
	}
}

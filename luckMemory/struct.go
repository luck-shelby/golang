/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  struct
 * @Date: 2020/12/2 15:17
 */
package main

import "fmt"

// 结构体是一个数据类型,自定义类型（定义后等价与 int,string,bool...）
type Person struct {
	name string
	sex byte
	age int
}

func main3() {
	// 顺序初始化: 需要将结构体所有成员初始化
	var xiaoMing Person = Person{"小明",'m',18}
	fmt.Printf("%v \n",xiaoMing)  // {小明 109 18}

	// 指定成员初始化, 未被初始化的成员是该数据类型的默认值, 如int默认为0...
	xiaoHong := Person{
		name: "小红",
		sex: 'w',
	}
	fmt.Printf("%v \n",xiaoHong)  // {小红 119 0}

	// 结构体比较, 只能使用 == 或者 ！=
	println(xiaoMing == xiaoHong,"or",xiaoMing != xiaoHong)

	// 当结构体中字段类型都一致时可以进行赋值
	var xiaoHua Person
	xiaoHua = xiaoHong
	fmt.Println(xiaoHua)  // {小红 119 0}

	// 结构体变量的地址 等于结构体首个元素地址
	fmt.Printf("xiaoHua 的地址是: %p \n",&xiaoHua)  // 0xc0000044e0
	fmt.Printf("xiaoHua结构体中的第一个元素的地址是:%p",&xiaoHua.name)  // 0xc0000044e0
}

func NewPerson()*Person  {
	return &Person{
		name: "spark",
		sex: 'm',
		age: 18,
	}
}
func main() {
	// 把结构体当作参数传给形参, 默认是值拷贝, 在实际中几乎都不会使用值拷贝方式, 而是使用结构体指针方式
	var xiaoMing *Person = &Person{"小明",'m',18}
	fmt.Printf("%v \n",xiaoMing)  // &{小明 109 18}

	// 方式二
	xiaoHone := new(Person)
	xiaoHone.name = "小红"
	xiaoHone.age = 18
	fmt.Printf("%v \n",xiaoHone)  // &{小红 0 18}

	// 指针结构体地址: 结构体指针变量地址
	fmt.Printf("xiaoHone 的地址是: %p \n",xiaoHone)  // 0xc0000044e0
	fmt.Printf("xiaoHone结构体中的第一个元素的地址是:%p, 值是:%v",&xiaoHone.name,xiaoHone.name)  // 0xc0000044e0

	// 指针做函数的返回值
	spark := NewPerson()
	fmt.Printf(" spark : %v \n",spark)  // spark : &{spark 109 18}
}
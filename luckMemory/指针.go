package main

import "fmt"
func main() {
	/*
		指针就是地址, 指针变量就是存储地址的变量
	*/
	// 指针声明而没有赋值，默认为nil，即该指针没有任何指向。当指针没有指向的时候，不能对(*point)进行操作包括读取，否则会报空指针异常。
	a := 100
	var b *int
	b = &a
	fmt.Printf("b: %p %#v %#v \n", &b, b, *b)

	*b = 1000  // *b 称为解引用,或者间接引用
	fmt.Printf("b: %p %#v %#v \n", &b, b, *b)

	// 内存
	/*
		变量存储在内存的栈(栈是一种数据结构，它按照后进先出的原则存储数据)中,而函数运行时在栈中维护一个独立的栈帧(stack frame)给函数运行的内存空间,
			函数调用接受后就会释放栈帧
			栈帧存储: 1. 局部变量 2. 形参 3. 描述信息
	*/

	// 空指针: 未被初始化的指针  如 var b *int
	/*
	var ab *int
	fmt.Println(*ab) // 未被初始化报错: invalid memory address or nil pointer dereference
	*/
	var ab *int = &a
	fmt.Println(ab,*ab)  // 0xc000010090 1000

	// 野指针: 被无效的地址空间初始化

	// 变量: 在等号左边是代表 该变量指向的内存空间, 在右边代表变量内存空间存储的数据值

	// new() : 在堆(heap)中申请一片内存空间
	var abc *string
	abc =  new(string)  // 这样就不是一个空指针,默认时类型的默认值，返回该内存空间地址
	fmt.Printf("abc = %q \n",*abc)  // abc = ""

	//  在堆(heap)中内存可以无限申请使用, 当使用完可以把指针置为空.让GC垃圾回收机制处理

	// 指针作为函数的参数或者返回值
	/*
		传引用: 将地址的值作为函数的参数，返回值
		传值: 将值拷贝一份给形参
	*/
	q := 10
	w := 100
	test(&q,&w)
	fmt.Printf("main  q = %d ,w = %d ,q=%p\n",q,w,&q)
}

func test(q,w *int)  {  //  函数传参: 不管传的是什么类型,都是值拷贝
	// q,w = w,q   只是交换了指针地址
	*q,*w = *w,*q  // 传引用时 如 在A 栈帧中去修改B栈帧的数据
	fmt.Printf("test  q = %d ,w = %d w=%q\n",*q,*w,w)
}
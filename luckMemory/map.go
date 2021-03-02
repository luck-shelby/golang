/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  map
 * @Date: 2020/12/1 21:44
 */
package main

import (
	"fmt"
	"strings"
)

// 字典 ：它是无序的并且唯一  key值不能是引用类型, map是引用类型
func main1() {
	var m map[int]string  // 声明map,没有创建空间不能使用 报错:  assignment to entry in nil map
	if m == nil{
		fmt.Println("assignment to entry in nil map")
	}else {
		m[1] = "spark"
	}

	// 方式二
	m2 := map[int]string{}
	m2[1] = "spark"
	fmt.Println(m2)  // map[1:spark]

	// 方式三
	m3 := make(map[int]string)
	m3[1] = "double"
	fmt.Println(m3)  // map[1:double]

	// map不能使用cap
	m4 := make(map[int]string,5) // 5 代表这个map能存储多少键值对,可以自动扩容
	m4[1] = "哈哈"
	fmt.Println(m4,len(m4))  // map[1:哈哈] 1
}

func MapDelete(data map[int]string,key int) {
	_,ok :=  data[key]
	if ok{
		delete(data,key)
	}else {
		fmt.Println("删除对象不存在!")
	}
}
func main2() {
	// 赋值过程中,如果出现key一致就会替换掉之前的key的值

	// 判断key是否存在
	m1 := make(map[int]string,5)
	m1[1] = "哈哈"
	m1[2] = "哈哈哈"
	m1[3] = "哈哈哈哈"
	m1[4] = "哈哈哈哈哈"
	m1[5] = "哈哈哈哈哈哈"
	v,ok := m1[1]
	if ok{
		fmt.Println("v",v)
	}

	// 删除key
	delete(m1,3)
	fmt.Println("m1",m1)

	// map当前函数的参数  传的是引用与切片一致
	MapDelete(m1,5)
	fmt.Println("m1",m1)
}

func totalWords(data string)(result map[string] int) {
	s := strings.Fields(data)  // 按照字符串空格拆分成字符串切片  [i love my work and i love  my family too]
	result = make(map[string]int)
	for i := 0;i <len(s);i++{
		if _,ok := result[s[i]];ok{
			result[s[i]] = result[s[i]]+1
		}else {
			result[s[i]] = 1
		}
	}
	return
}
func main()  {
	// exercise  统计字符串单词格式
	str := "i love my work and i love love love my family too"
	result := totalWords(str)
	fmt.Println(result)
}
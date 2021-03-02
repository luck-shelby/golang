/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  切片
 * @Date: 2020/11/30 14:35
 */
package main

import "fmt"

/*
	数组缺点: (1) 数组的容量固定,不能自动扩容。(2) 数组是值传递,数组作为函数的参数时,就会将数组值拷贝一份给形参
	切片:
		slice是引用类型,总是指向一个底层的array
	切片使用:
		切片名称【low:high:max】
		low: 起始下标位置 high: 结束下标位置(len=high-low)  max: 容量(max-low)
*/

func main01() {
	arr := [6]int{1,2,3,4,5,6}
	fmt.Println("原数组:",arr,len(arr),cap(arr))
	s := arr[2:4:6]  // len(4-1=3) cap(6-1=5)
	fmt.Println(s,len(s),cap(s))

	// 当切片没有指定容量, 它的容量是跟原数组/切片一致,原数组容量为6,没有指定容量是从起始位置开始计算: 原数组的容量-low就是当前切片的容量
	s1 := arr[3:5]
	fmt.Println(s1,len(s1),cap(s1)) // 4 5
}

func main02() {
	// 容量注意事项(没有指定容量)
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	s := arr[2:5]  // [3 4 5]
	fmt.Println(s,len(s),cap(s)) // 3 8 切片没有指定容量, 它的容量是跟原数组/切片一致 原(10)-low = 8

	s1 := s[2:5] // [5 6 7]
	fmt.Println(s1,len(s1),cap(s1)) // 3 6
}

func main03() {
	// 容量注意事项(指定容量)
	arr := [10]int{1,2,3,4,5,6,7,8,9,10}
	s := arr[2:5:5]  // [3 4 5]
	fmt.Println(s,len(s),cap(s))  // 3 3  max: 容量(max-low)

	// 报错: 示例
	s1 := s[2:5]
	/*
	报错: runtime error: slice bounds out of range [:5] with capacity 3
	因为s 容量为3,指定的范围超出 容量
	*/
	fmt.Println(s1,len(s1),cap(s1))
}

func Test1(data []string)[]string  {
	color := make([]string,0)
	for _,str := range data{
		if str != ""{
			color = append(color,str)
		}
	}
	return color
}
// 切片截取 不使用append方式
func Test2(data []string)[]string  {
	index := 0
	for _,str := range data{
		if str != ""{
			data[index] = str
			index ++
		}
	}
	return data[:index]
}

// 去重
func Test3(data []string)[]string  {
	color := data[:1]
	for _,val := range data{
		flg := true
		for i:=0; i<len(color);i++{
			if val == color[i] {
				flg = false
				break
			}
		}
		if flg{
			color = append(color,val)
		}
	}
	return color
}

func Test4(data []int, num int)  []int{
	n := -1
	for i,val := range data{
		if val == num{
			n = i
			break
		}
	}
	if n < 0 {
		return nil
	}
	data = append(data[:n], data[n+1:]...)
	return data
}
func main() {
	// 使用make 创建后没有指定容量. 默认容量等于长度  make只能创建slice,map,channel
	s := make([]int,5)  // [0 0 0 0 0] 5 5  使用make创建后切片是有零值的
	fmt.Println(s,len(s),cap(s))

	// append 向切片尾部追加

	// exercise
	color1 := []string{"red","black","","yellow"}
	newColor := Test1(color1)
	fmt.Println(newColor)

	color2 := []string{"red","black","","yellow"}
	newColor2 := Test2(color2)
	fmt.Println("newColor2",newColor2)

	// 字符串去重
	color3 := []string{"red","black","red","yellow","black","blue","blue","blue"}
	newColor3 := Test3(color3)
	fmt.Println("newColor3",newColor3)

	// copy(目标位置切片, 源切片) 对应位置拷贝  复制长度以 len 小的为准。
	data := [...]int{1,2,3,4,5,6,7,8,9,10}
	s1 := data[8:] 		// [9 10]
	s2 := data[:5] 		// [1 2 3 4 5]
	copy(s2,s1)        // 将s1复制到s2里
	fmt.Println("s2 = ",s2)  // [9 10 3 4 5]

	// 删除切片元素
	data1 := []int{5,6,7,8,9}
	result := Test4(data1,19)
	if result == nil{
		fmt.Println("删除元素不存在!")
	}else {
		fmt.Println("result",result)
	}
}
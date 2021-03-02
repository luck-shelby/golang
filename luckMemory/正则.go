/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  正则
 * @Date: 2020/12/16 15:53
 */
package main

import (
	"fmt"
	"regexp"
)

/*
	(): 把一部分正则括起来组成一个单元,然后就可以对这个单元进行匹配限制（限定符）
	在提取网页数据只要指定起始位置后,中间的内容都会被提取下来: (?s:(.*?))  如 <div>(?s:(.*?))</div>
*/
func main() {
	str := "You want to see a miracle, son? Be the miracle I guess I kinda liked the way you numbed all the pain I was " +
		"getting kinda used to being someone you loved"
	// 编译正则表达式
	reg := regexp.MustCompile(`t[o]+`)
	ret1 := reg.FindAllStringSubmatch(str,-1)

	// 匹配小数
	fractionalPart := "3.14 and 所以13.76是 75.13792864928…% 将显示为 75.13792%"
	digit := regexp.MustCompile(`[\d]+\.[\d]*`)
	ret2 := digit.FindAllStringSubmatch(fractionalPart,-1)
	fmt.Println("ret1 = ",ret1, "ret2 = ",ret2)
}

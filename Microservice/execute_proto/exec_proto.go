/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File: execute_proto
 * @Date: 2020/12/25 22:06
 */
package main

import (
	"Lucks/Microservice"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	text := &microServer.PersonRequest{
		Name: "double",
		Age: 18,
		Hobby: []string{"bar","ktv"},
	}
	fmt.Println(text)
	// 编码成二进制
	data,err := proto.Marshal(text)
	if err != nil {
		fmt.Println("proto.Marshal err:",err)
		return
	}
	fmt.Println(data)

	// 解码
	newText := &microServer.PersonRequest{}
	err = proto.Unmarshal(data,newText)
	if err != nil {
		fmt.Println("proto.Unmarshal err:",err)
		return
	}
	fmt.Println(newText)
	fmt.Println(newText.GetAge(),"或者,",newText.Age)
}
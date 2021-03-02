/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/15 17:28
 */
package main

import (
	"fmt"
	"net"
	"os"
)

func ErrHandle(err error,info string)  {
	if err != nil {
		fmt.Println(info,err)
		os.Exit(1) // 将当前进程结束
	}
}
func main() {
	conn,err := net.Dial("tcp","127.0.0.1:8809")
	ErrHandle(err,"net.Dial: ")
	defer conn.Close()

	httpRequest := "GET /memberList HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"
	_, _ = conn.Write([]byte(httpRequest))

	buf := make([]byte,4096)
	for  {
		n,err := conn.Read(buf)
		ErrHandle(err,"net.Dial: ")
		if n ==0{
			return
		}
		fmt.Println("-----",string(buf[:n]))
	}
}

/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/11 15:23
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	conn,err := net.Dial("udp","127.0.0.1:8018")
	if err != nil {
		fmt.Println("net.Dial ERR:",err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for  {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if inputInfo == "exit"{
			fmt.Println( "客户端断开连接!!!")
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := make([]byte,4096)
		n, err := conn.Read(buf[:])
		if n == 0{
			fmt.Println("检测到服务器已经关闭!!!,程序退出")
			return
		}
		if err == io.EOF{
			fmt.Println("Exit success")
			return
		}
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("服务器返回信息：",string(buf[:n]))
	}
}

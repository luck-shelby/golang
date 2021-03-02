/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/11 17:03
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func send(conn net.Conn,path string)  {
	// （3）逐行读取 可以带有缓冲区(用户缓区)的文件读取
	file,err := os.OpenFile(path,os.O_RDWR,0666)
	if err != nil {
		fmt.Println("打开文件失败:",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for  {
		content,err := reader.ReadBytes('\n')
		// content,err := reader.ReadString('\n')
		if err == io.EOF || err != nil{
			break
		}
		_, _ = conn.Write(content)
	}
}
func main() {
	argList  := os.Args   // 获取命令行参数
	if len(argList) != 2{
		fmt.Println("参数格式错误! ")
		return
	}
	filePath := argList[1]
	fileInfo ,err := os.Stat(filePath)
	if err != nil {
		fmt.Println("获取文件属性错误! ")
		return
	}
	fileName := fileInfo.Name()

	conn,err := net.Dial("tcp","127.0.0.1:8009")
	if err != nil {
		fmt.Println("net.Dial ERR:",err)
		return
	}
	defer conn.Close()

	_,err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("conn.Write ERR:",err)
		return
	}

	buf := make([]byte,1096)
	n,err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read ERR:",err)
		return
	}
	if string(buf[:n]) == "ok"{
		// 发送文件内容
		send(conn,argList[1])
	}
	fmt.Println("服务端返回信息:",string(buf[:n]))

}

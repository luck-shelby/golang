/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  server
 * @Date: 2020/12/11 19:28
 */
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"runtime"
	"strings"
)
func GetCurrentDirectory() string {
	_, fileStr, _, _ := runtime.Caller(1)
	str := path.Dir(fileStr)
	return strings.Replace(str, "\\", "/", -1) //将\替换成/
}
func recv(conn net.Conn,fileName string)  {
	str := GetCurrentDirectory()
	f,err := os.OpenFile(str + "/" +fileName, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败:",err)
		return
	}
	defer f.Close()
	// 字节方式写入
	buf := make([]byte,4096)
	for  {
		n,err := conn.Read(buf)
		if n == 0{
			fmt.Println("文件接收完成")
			return
		}
		_, err = f.Write(buf[:n])
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func main() {
	listen,err := net.Listen("tcp","127.0.0.1:8009")
	if err != nil {
		return
	}
	defer listen.Close()
	fmt.Println("等待客户端请求...")
	conn,err := listen.Accept()
	if err != nil {
		return
	}
	buf := make([]byte,1024)
	n,err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read ERR :",err)
		return
	}
	_, _ = conn.Write([]byte("ok"))

	// 获取文件内容
	fileName := string(buf[:n])
	fmt.Println("文件名称:",fileName)
	recv(conn,fileName)
}

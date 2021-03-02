/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  socket
 * @Date: 2020/12/8 16:20
 */
package main

import (
	"fmt"
	"net"
)

/*
	socket: 网络通信过程中, socket一定是成对出现的 双向全双工,服务器中有两个套接字
*/
func HandleConn(conn net.Conn)  {
	fmt.Println("客户端建立连接成功，等待客户端发送数据: ",conn.RemoteAddr().String())
	defer conn.Close()
	// 循环读取客户端数据
	buf := make([]byte,4096)
	for  {
		n,err:=conn.Read(buf)
		fmt.Println("字节: ",buf[:n])
		if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit"{
			fmt.Println(conn.RemoteAddr().String() +"客户端请求断开连接!!!")
			return
		}
		if n == 0{
			fmt.Println(conn.RemoteAddr().String() +"客户端断开连接!!!")
			return
		}
		if err != nil {
			fmt.Println("conn.Read ERR: ",err)
			return
		}
		fmt.Println("读到数据:",string(buf[:n]))
		// 读多少写多少
		_, _ = conn.Write(buf[:n])
	}
}
func main() {
	// 创建一个监听的套接字
	socket,err := net.Listen("tcp","127.0.0.1:8009")
	if err != nil {
		fmt.Println("net.Listen ERR: ",err)
		return
	}
	defer socket.Close()
	// 阻塞监听客户端连接请求
	for  {
		fmt.Println("等待客户端建立连接... ")
		conn,err := socket.Accept()
		if err != nil {
			fmt.Println("socket.Accept ERR: ",err)
			return
		}
		// 启用go程
		go HandleConn(conn)
	}
}
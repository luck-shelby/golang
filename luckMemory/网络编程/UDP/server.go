/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  server
 * @Date: 2020/12/10 22:23
 */
package main

import (
	"fmt"
	"net"
	"time"
)

/*
	无连接,不可靠的报文传输,使用场景: 对数据实时传输要求较高的场合下,如 视频直播,电话会议 游戏等
*/

func main() {
	addr,err := net.ResolveUDPAddr("udp","127.0.0.1:8018")
	if err != nil {
		fmt.Println("net.ResolveIPAddr ERR: ",err)
		return
	}
	// 创建用于通信的socket
	conn,err := net.ListenUDP("udp",addr)
	if err != nil {
		fmt.Println("net.ListenUDP ERR: ",err)
		return
	}
	defer conn.Close()
	for  {
		fmt.Println("服务器socket创建完成,等到客户端发送消息...")
		buf := make([]byte,4096)
		n,cltAddr,err := conn.ReadFromUDP(buf) // 读取的字节数, 客户端信息
		if  string(buf[:n]) == "exit"{
			fmt.Println( cltAddr,"客户端请求断开连接!!!")
			continue
		}
		if n == 0{
			fmt.Println(conn.RemoteAddr().String() +"客户端断开连接!!!")
			return
		}
		if err != nil {
			fmt.Println("conn.ReadFromUDP ERR: ",err)
			return
		}
		fmt.Printf("服务器读到来自 %v数据 %s \n",cltAddr,string(buf[:n]))
		dayTime := time.Now().String()
		_, _ = conn.WriteToUDP([]byte(dayTime), cltAddr)
	}
}
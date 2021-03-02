/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  chatRomm
 * @Date: 2020/12/12 18:09
 */
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 创建用户结构体类型
type Client struct {
	C chan string
	Name string
	Address string
}
// 创建全局map存储在线用户
var onlineMap map[string] Client

// 创建全局的channel 传递用户消息
var message = make(chan string)

func MakeMessage(clt Client,msg string)(message string)  {
	message = "[" + clt.Address + "]" + clt.Name +" : " + msg
	return
}

func handlerConnect(conn net.Conn)  {
	defer conn.Close()
	// 创建一个channel,判断用户是否退出
	quit := make(chan bool)
	// 创建一个channel,判断用户在规定时间内是否活跃
	active := make(chan bool)
	// 创建新连接用户信息，默认都是IP
	netAddr := conn.RemoteAddr().String()
	clt := Client{
		C: make(chan string),
		Name: netAddr,
		Address: netAddr,
	}
	// 将用户信息添加到map
	onlineMap[netAddr] = clt
	// 创建给当前登录用户读取消息的go程
	go WriteMsgToClient(clt,conn)
	// 发送用户上线信息到全局channel中
	message <- MakeMessage(clt,"just went online!")
	// 处理用户发送的消息
	buf := make([]byte,4096)
	go func() {
		for  {
			n,err:=conn.Read(buf)
			if n == 0{
				quit <- true
				fmt.Printf("%v 客户端退出 \n",clt.Name)
				return
			}
			if err != nil {
				fmt.Println("读取客户端消息错误: ",err)
				return
			}
			msg := string(buf[:n])
			inputInfo := strings.Trim(msg, "\r\n")
			if strings.ToUpper(inputInfo) == "WHO"{
				_, _ = conn.Write([]byte("Current online users: \n"))
				for _,user := range onlineMap{
					userInfo := user.Address + ":" + user.Name + "\n"
					_, _ = conn.Write([]byte(userInfo))
				}
			}else if len(strings.Split(inputInfo,"|"))== 2 &&  strings.ToUpper(strings.Split(inputInfo,"|")[0]) == "RENAME" {
				// 替换
				clt.Name = strings.Split(inputInfo,"|")[1]
				onlineMap[netAddr] = clt
				_, _ = conn.Write([]byte("Amendments to the success of information! \n"))
			}else {
				message <- MakeMessage(clt,msg)
			}
			// 只要有数据传递就会执行下面的
			active <- true
		}
	}()
	for  {
		select {
		// 监听channel
		case <-quit:
			// 将用户 从 map中移除
			close(clt.C)  // 当在其他函数中创建的go程中,函数退出不会导致它对应的子go程退出,所以手动把它关闭
			delete(onlineMap,clt.Address)
			message <- MakeMessage(clt,"logout")
			return
		case <- active: // 只要执行到这个case就又会重新执行for,就不会往下执行,那么下面的case就会重新计时
		case <-time.After(time.Second*10):
			close(clt.C)
			delete(onlineMap,clt.Address)
			message <- MakeMessage(clt,"logout")
			return
		}
	}
}

func Manage()  {
	// 初始化onlineMap
	onlineMap = make(map[string]Client)
	// 循环监听全局channel是否有数据，存储
	for  {
		msg := <- message
		for _,clt :=range onlineMap{
			// 循环发送消息给在线用户
			clt.C <- msg
		}
	}
}

func WriteMsgToClient(clt Client,conn net.Conn)  {
	for msg := range clt.C{
		// 监听用户自带的channel是否有数据,把读到的数据通过socket发给终端
		_, _ = conn.Write([]byte(msg + "\n"))
	}
}

func main() {
	listen,err := net.Listen("tcp","172.21.165.61:8809")
	if err != nil {
		fmt.Println("Listen ERR: ",err)
		return
	}
	defer listen.Close()
	// 用来管理map和全局channel
	go Manage()
	// 监听客户端连接请求
	for  {
		conn,err := listen.Accept()
		fmt.Printf("客户端%v 已经建立连接 \n",conn.RemoteAddr().String())
		if err != nil {
			fmt.Println("Accept ERR: ",err)
			return
		}
		// 启动go程处理客户端请求
		go handlerConnect(conn)
	}
}

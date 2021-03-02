/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  server
 * @Date: 2020/12/26 13:33
 */
package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A,B int
}
// 服务器端需要注册结构体对象，然后通过对象所属的方法暴露给调用者
type MathService struct {}

func (self *MathService) Multiply(args *Args,reply *int)error  {
	*reply = args.A * args.B
	return nil
}
func (self *MathService) Divide(args *Args,reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	*reply = args.A / args.B
	return nil
}
func main() {
	// 初始化指针数据类型
	math := new(MathService)
	// 将服务对象进行注册
	// _ = rpc.Register(math) 注册方式一
	err := rpc.RegisterName("MathService",math)  // 注册方式二
	if err != nil {
		panic(err.Error())
	}
	// 将MathService中提供的服务注册到HTTP协议上,方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()
	// 在特定的端口进行监听
	listen,err := net.Listen("tcp",":8089")
	if err != nil {
		fmt.Println("启动服务监听失败:",err.Error())
		return
	}
	defer listen.Close()
	err = http.Serve(listen,nil)
	if err != nil {
		fmt.Println("启动 HTTP 服务失败:",err.Error())
		return
	}
}


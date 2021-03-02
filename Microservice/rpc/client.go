/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/26 13:55
 */
package main

import (
	"fmt"
	"net/rpc"
)

type ab struct {
	A,B int
}

func main() {
	arg := ab{100,0}
	client, err := rpc.DialHTTP("tcp",  "127.0.0.1:8089")
	if err != nil {
		fmt.Println("建立与服务端连接失败:",err.Error())
		return
	}
	defer client.Close()
	var reply int
	// 远端方法调用 客户端成功连接服务端以后，就可以通过方法调用调用服务端的方法
	err = client.Call("MathService.Divide",arg,&reply)  // 同步的调用
	/*
		client.Call:第一个参数表示要调用的远端服务的方法名，第二个参数是调用时要传入的参数，第三个参数是调用要接收的返回值
	*/
	if err != nil {
		fmt.Println("调用远程方法 MathService.Multiply 失败:",err.Error())
		return
	}
	// 异步调用
	/*
	syncCall := client.Go("MathService.Divide", arg,&reply,nil)
	replayDone := <-syncCall.Done
	fmt.Println(replayDone)
	fmt.Println(reply)
	*/
	fmt.Printf("%d*%d=%d\n", arg.A, arg.B, reply)

}

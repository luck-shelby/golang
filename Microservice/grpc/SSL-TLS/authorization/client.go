/*
@Author : YanLeJun
@Description: 我相信一切都是最好的安排
@File : client
@Date: 2020/12/31 20:55
@Version: 1.5.3
*/
package main

import (
	pd "Lucks/Microservice/grpc/SSL-TLS/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)
// 在go1.5 需要在运行前前添加配置: GODEBUG=x509ignoreCN=0
func main() {
	//从输入的证书文件中为客户端构造TLS凭证
	cert, err := credentials.NewClientTLSFromFile("Microservice/grpc/SSL-TLS/keys/server.pem","luck")
	if err != nil {
		fmt.Println("加载证书失败:",err.Error())
		return
	}
	conn,err := grpc.Dial("localhost:8809",[]grpc.DialOption{grpc.WithTransportCredentials(cert)}...)
	if err != nil {
		fmt.Println("连接失败:",err.Error())
		return
	}
	defer conn.Close()
	// 创建gRPC客户端
	client := pd.NewMathServiceClient(conn)
	add := pd.RequestArgs{Arg1: 1,Arg2: 2}
	// 超时
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()
	response,err := client.AddMethod(ctx,&add)
	if err != nil {
		fmt.Println("调用远程方式失败:",err.Error())
		return
	}
	fmt.Println(response.GetCode(),response.GetMsg())
}

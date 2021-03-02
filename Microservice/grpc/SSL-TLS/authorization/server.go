/*
@Author : YanLeJun
@Description: 我相信一切都是最好的安排
@File : server
@Date: 2020/12/31 20:54
@Version: 1.5.3
*/
package main

import (
	pd "Lucks/Microservice/grpc/SSL-TLS/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

/*
	SSL全称是Secure Sockets Layer，又被称之为安全套接字层，是一种标准安全协议
	是通过非对称加密的方式来实现的,所谓非对称加密方式又称之为公钥加密，密钥对由公钥和私钥两种密钥组成。私钥和公钥成对存在，
	先生成私钥，通过私钥生成对应的公钥。
	加密过程:
		客户端想要向服务器发起链接，首先会先向服务端请求要加密的公钥。获取到公钥后客户端使用公钥将信息进行加密，服务端接收到加密信息，
		使用私钥对信息进行解密并进行其他后续处理，完成整个信道加密并实现数据传输的过程
*/
type MathManage struct {}

func (self *MathManage)AddMethod(ctx context.Context, request *pd.RequestArgs) (response *pd.ResponseArgs, err error)  {
	fmt.Println(" 服务端 Add方法 ")
	result := request.Arg1 + request.Arg2
	fmt.Println(" 计算结果是：", result)
	response = new(pd.ResponseArgs)
	response.Code = 1
	response.Msg = "执行成功"
	return response, nil
}

func main() {
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	cert, err := credentials.NewServerTLSFromFile("Microservice/grpc/SSL-TLS/keys/server.pem", "Microservice/grpc/SSL-TLS/keys/server.key")
	if err != nil {
		fmt.Println("加载证书失败:",err.Error())
		return
	}
	//实例化gRPC server, 开启TLS认证
	server := grpc.NewServer(grpc.Creds(cert))
	// 在gRPC服务器注册我们的服务
	pd.RegisterMathServiceServer(server,new(MathManage))
	listen,err := net.Listen("tcp",":8809")
	if err != nil {
		fmt.Println("listen tcp err:",err.Error())
		return
	}
	_ = server.Serve(listen)
}
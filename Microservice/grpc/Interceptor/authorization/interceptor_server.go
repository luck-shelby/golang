/*
@Author : YanLeJun
@Description: 我相信一切都是最好的安排
@File : token_server
@Date: 2021/1/1 20:43
@Version: 1.5.3
*/
package main

import(
	"Lucks/Microservice/grpc/Interceptor"
	pd "Lucks/Microservice/grpc/Interceptor/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

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
	//实例化gRPC server, 开启TLS认证,与拦截器认证
	server := grpc.NewServer(grpc.Creds(cert),grpc.UnaryInterceptor(Interceptor.TokenInterceptor))
	// 在gRPC服务器注册我们的服务,
	pd.RegisterMathServiceServer(server,new(MathManage))
	listen,err := net.Listen("tcp",":8809")
	if err != nil {
		fmt.Println("listen tcp err:",err.Error())
		return
	}
	_ = server.Serve(listen)
}

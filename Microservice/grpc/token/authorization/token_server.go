/*
@Author : YanLeJun
@Description: 我相信一切都是最好的安排
@File : token_server
@Date: 2021/1/1 20:43
@Version: 1.5.3
*/
package main

import (
	pd "Lucks/Microservice/grpc/token/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
)
/*
	Token认证:基于Token的身份验证是无状态，不需要将用户信息服务存在服务器或者session中。
	客户端在发送请求前，首先向服务器发起请求，服务器返回一个生成的token给客户端。客户端将token保存下来，用于后续每次请求时，携带着token参数。
		服务端在进行处理请求之前，会首先对token进行验证，只有token验证成功了，才会处理并返回相关的数据。
*/
type MathManage struct {}

func (self *MathManage)AddMethod(ctx context.Context, request *pd.RequestArgs) (response *pd.ResponseArgs, err error)  {
	// 解析metada中的信息并验证
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appKey string
		appSecret string
	)
	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}
	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}
	if appKey != "hello" || appSecret != "20190812" {
		return nil,  grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appKey, appSecret)
	}

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

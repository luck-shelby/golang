/*
@Author : YanLeJun
@Description: 我相信一切都是最好的安排
@File : token_client
@Date: 2021/1/1 20:43
@Version: 1.5.3
*/
package main

import (
	pd "Lucks/Microservice/grpc/token/message"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)
// 允许开发者自定义自己的认证规则,需要实现grpc.WithPerRPCCredentials()接口
type TokenAuthentication struct {
	AppKey    string
	AppSecret string
}
//组织token认证的metadata信息
func (self *TokenAuthentication)GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)  {
	return map[string]string{
		// appid和appkey字段均需要保持小写，不能大写
		"appid":self.AppKey,
		"appkey":self.AppSecret,
	},nil
}
//	是否基于TLS认证进行安全传输
func (self *TokenAuthentication)RequireTransportSecurity() bool  {
	return true
}
func main() {
	//从输入的证书文件中为客户端构造TLS凭证
	cert, err := credentials.NewClientTLSFromFile("Microservice/grpc/SSL-TLS/keys/server.pem","luck")
	if err != nil {
		fmt.Println("加载证书失败:",err.Error())
		return
	}
	// 自定义的token认证信息作为参数进行传入
	auth := TokenAuthentication{
		AppKey:    "hello",
		AppSecret: "20190812",
	}
	conn,err := grpc.Dial("localhost:8809",grpc.WithTransportCredentials(cert),grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		fmt.Println("连接失败:",err.Error())
		return
	}
	defer conn.Close()
	client := pd.NewMathServiceClient(conn)
	add := pd.RequestArgs{Arg1: 1,Arg2: 2}
	response,err := client.AddMethod(context.Background(),&add)
	if err != nil {
		fmt.Println("调用远程方式失败:",err.Error())
		return
	}
	fmt.Println(response.GetCode(),response.GetMsg())
}

/*
@Author : YanLeJun
@Description: 我相信一切都是最好的安排
@File : interceptor
@Date: 2021/1/1 21:40
@Version: 1.15.0
*/
package Interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/*
	拦截器(Interceptor):
		在服务端的方法中，每个方法都要进行token的判断。程序效率太低，可以优化一下处理逻辑，在调用服务端的具体方法之前，先进行拦截，
			并进行token验证判断，这种方式称之为拦截器处理。除了此处的token验证判断处理以外，还可以进行日志处理
		在NewSever时添加拦截器设置，grpc框架中可以通过grpc.UnaryInterceptor()方法设置自定义的拦截器，并返回ServerOption
*/
// 自定义拦截器
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	//通过metadata获取token认证信息
	md,exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var appKey string
	var appSecret string
	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}
	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}
	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
	}

	// 如果通过token验证，继续处理请求,后续继续处理可以由grpc.UnaryHandler进行处理
	// grpc.UnaryHandler同样是一个方法,具体的实现就是自己 定义实现的服务方法
	return handler(ctx, req)
}


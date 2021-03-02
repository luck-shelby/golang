/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/26 18:24
 */
package main

import (
	orderServer "Lucks/Microservice/grpc/clientway-RPC/order_service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"time"
)

func main() {
	conn,err := grpc.Dial("localhost:8809", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed creates a client connection")
		return
	}
	defer conn.Close()
	orderMap := map[string]orderServer.OrderRequest{
		"201907300001": orderServer.OrderRequest{OrderId: "201907300001", TimeStamp: time.Now().Unix()},
		"201907310001": orderServer.OrderRequest{OrderId: "201907310001", TimeStamp: time.Now().Unix()},
		"201907310002": orderServer.OrderRequest{OrderId: "201907310002", TimeStamp: time.Now().Unix()},
	}
	client := orderServer.NewOrderServiceClient(conn)
	// 调用服务方法
	addOrderListClient, err := client.AddOrderList(context.Background())
	if err != nil {
		fmt.Println("调用服务失败,:",err.Error())
		return
	}
	for _,info := range orderMap{
		err = addOrderListClient.Send(&info)
		if err != nil {
			fmt.Println("发送数据错误,:",err.Error())
			return
		}
	}
	for {
		response, err := addOrderListClient.CloseAndRecv()
		if err == io.EOF {
			fmt.Println(" 读取数据结束了 ")
			return
		}
		if err != nil {
			fmt.Println("接收服务端错误:",err.Error())
			return
		}
		fmt.Println("服务端:",response.GetOrderStatus())
	}
}
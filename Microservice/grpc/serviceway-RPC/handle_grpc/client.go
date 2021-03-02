/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/26 18:24
 */
package main

import (
	orderServer "Lucks/Microservice/grpc/serviceway-RPC/order_service"
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
	client := orderServer.NewOrderServiceClient(conn)
	request := orderServer.OrderRequest{TimeStamp: time.Now().Unix()}
	orderInfosClient, err := client.GetOrderInfo(context.TODO(),&request)
	for  {
		orderInfo, err := orderInfosClient.Recv()
		if err == io.EOF{
			fmt.Println("读取结束")
			return
		}
		if err != nil {
			fmt.Println("读取结束")
			return
		}
		fmt.Println("读取到的信息：", orderInfo)
	}
}
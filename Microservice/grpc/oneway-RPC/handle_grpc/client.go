/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/26 18:24
 */
package main

import (
	orderServer "Lucks/Microservice/grpc/oneway-RPC/order_service"
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	orderRequest := orderServer.OrderRequest{OrderId: "201907310002",TimeStamp: time.Now().Unix()}
	info,err := client.GetOrderInfo(context.Background(),&orderRequest)
	if err != nil {
		fmt.Println("failed get order info,err:",err.Error())
		return
	}
	fmt.Println(info.GetOrderName(),info.GetOrderId(),info.GetOrderStatus())
}
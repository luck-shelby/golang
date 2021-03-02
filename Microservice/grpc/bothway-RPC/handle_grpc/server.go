/*
@Author : YanLeJun
@Description: 我相信一切都是最好的安排
@File : server
@Date: 2020/12/31 17:47
@Version: 1.5.3
*/
package main

import (
	orderServer "Lucks/Microservice/grpc/bothway-RPC/order_service"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
)
type OrderServiceImpl struct {}

func(o *OrderServiceImpl)GetOrderInfo(stream orderServer.OrderService_GetOrderInfoServer) error{
	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(" 数据读取结束 ")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}
		fmt.Println(orderRequest.GetOrderId())
		orderMap := map[string]orderServer.OrderResponse{
			"201907300001": orderServer.OrderResponse{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
			"201907310001": orderServer.OrderResponse{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
			"201907310002": orderServer.OrderResponse{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
		}
		result := orderMap[orderRequest.GetOrderId()]
		//发送数据
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println("数据发送完成: err:",err.Error())
			return err
		}
		if err != nil {
			fmt.Println("数据发送错误: err",err.Error())
			return err
		}
	}
	return nil
}
func main() {
	server := grpc.NewServer()
	orderServer.RegisterOrderServiceServer(server,new(OrderServiceImpl))
	listen,err := net.Listen("tcp",":8809")
	if err != nil {
		fmt.Println("listen server err:",err.Error())
		return
	}
	defer listen.Close()
	_ = server.Serve(listen)
}
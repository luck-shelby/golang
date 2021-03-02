/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  server
 * @Date: 2020/12/26 18:23
 */
package main

import (
	orderServer "Lucks/Microservice/grpc/serviceway-RPC/order_service"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

// 在.proto定义好服务接口并生成对应的go语言文件后，需要对服务接口做具体的实现
type OrderServiceImpl struct {}

func(o *OrderServiceImpl)GetOrderInfo(request *orderServer.OrderRequest, stream orderServer.OrderService_GetOrderInfoServer) error{
	fmt.Println(" 服务端流 RPC 模式")
	orderMap := map[string]orderServer.OrderResponse{
		"201907300001": orderServer.OrderResponse{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": orderServer.OrderResponse{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": orderServer.OrderResponse{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	for id,info := range orderMap{
		if time.Now().Unix() >= request.TimeStamp{
			fmt.Println("订单序列号ID：", id)
			fmt.Println("订单详情：", &info)
			_ = stream.Send(&info)
		}
	}
	return nil
}
func main() {
	server := grpc.NewServer()
	orderServer.RegisterOrderServiceServer(server,new(OrderServiceImpl))
	lis, err := net.Listen("tcp", ":8809")
	if err != nil {
		fmt.Println("failed to listen tcp, err:",err.Error())
		return
	}
	defer lis.Close()
	_ = server.Serve(lis)
}
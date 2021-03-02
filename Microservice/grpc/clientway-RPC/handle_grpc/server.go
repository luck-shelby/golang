/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  server
 * @Date: 2020/12/26 18:23
 */
package main

import (
	orderServer "Lucks/Microservice/grpc/clientway-RPC/order_service"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
)

// 在.proto定义好服务接口并生成对应的go语言文件后，需要对服务接口做具体的实现
type OrderServiceImpl struct {}

func(o *OrderServiceImpl)AddOrderList(stream orderServer.OrderService_AddOrderListServer) error{
	fmt.Println(" 客户端流 RPC 模式")
	for  {
		//从流中读取数据信息
		orderRequest, err := stream.Recv()
		if err == io.EOF{
			fmt.Println(" 读取数据结束 ")
			result := orderServer.OrderResponse{OrderStatus: "数据读取完成"}
			return stream.SendAndClose(&result)
		}
		if err != nil {
			fmt.Println(" 读取数据错误 ")
			return err
		}
		//打印接收到的数据
		fmt.Println(orderRequest)
	}
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
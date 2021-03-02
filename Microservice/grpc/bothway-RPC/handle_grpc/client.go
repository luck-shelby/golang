/*
@Author : YanLeJun
@Description: 我相信一切都是最好的安排
@File : client
@Date: 2020/12/31 17:47
@Version: 1.5.3
*/
package main
import (
	orderServer "Lucks/Microservice/grpc/bothway-RPC/order_service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
)

func main() {
	conn, err := grpc.Dial("localhost:8809", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := orderServer.NewOrderServiceClient(conn)

	fmt.Println("客户端请求RPC调用：双向流模式")
	orderIDs := []string{"201907300001", "201907310001", "201907310002"}

	orderInfoClient, err := orderServiceClient.GetOrderInfo(context.Background())
	for _, orderID := range orderIDs {
		orderRequest := orderServer.OrderRequest{OrderId: orderID}
		err := orderInfoClient.Send(&orderRequest)
		if err != nil {
			panic(err.Error())
		}
	}
	//关闭
	_ = orderInfoClient.CloseSend()
	for {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}
		if err != nil {
			return
		}
		fmt.Println("读取到的信息：", orderInfo)
	}
}

/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  server
 * @Date: 2020/12/29 17:48
 */
package main

import (
	message "Lucks/Microservice/rpc/RPC_Protobuf"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type OrderServer struct {}

func (receiver *OrderServer) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	}else {
		result := orderMap[request.OrderId]
		if result.OrderId != ""{
			fmt.Println( orderMap[request.OrderId])
			*response = orderMap[request.OrderId]
		}else {
			return errors.New("server error")
		}
	}
	return nil
}

func main() {
	orderServer := new(OrderServer)
	err := rpc.Register(orderServer)
	if err != nil {
		panic(err.Error())
	}
	rpc.HandleHTTP()
	listen,err := net.Listen("tcp",":8089")
	if err != nil {
		return
	}
	_ = http.Serve(listen, nil)
}

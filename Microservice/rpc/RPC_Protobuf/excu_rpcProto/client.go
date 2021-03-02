/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  client
 * @Date: 2020/12/29 17:48
 */
package main

import (
	message "Lucks/Microservice/rpc/RPC_Protobuf"
	"fmt"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8089")
	if err != nil {
		panic(err.Error())
	}
	timeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "201907300001", TimeStamp: timeStamp}
	var response *message.OrderInfo
	err = client.Call("OrderServer.GetOrderInfo", request, &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response)
}

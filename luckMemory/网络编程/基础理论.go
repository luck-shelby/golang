/**
 * @Author: YanLeJun
 * @Description: 我相信一切都是最好的安排
 * @File:  基础理论
 * @Date: 2020/12/8 12:45
 */
package main

/*
	1. 协议:  一组规则,要求使用协议双方必须严格尊重协议内容
	2. 典型协议
		传输层: 常见的是 TCP/UDP协议
		应用层: 常见的是Http,FTP协议
		网络层: 常见的是IP,ICMP,IGMP协议
		网络接口层: 常见的是ARP,RARP协议
			TCP: 传输控制协议,是一种面向连接的,可靠的,基于字节流的传输层通信协议
			UDP: 用户数据报协议,是一种无连接的传输协议,提供面向事务的简单不可靠信息传递服务
			APR: 通过已知的IP,寻找对应主机的MAC地址
			RARP: 是反向地址转换协议,根据MAC地址确认IP地址
	3. 网络分层架构
		基于TCP/IP: 四层  应用层--传输层--网络层--链路层
		链路层: 借助与ARP协议寻找网络中另一台主机的MAC地址
		网络层：确认网络中另一台主机的IP地址,其中的IP协议就是标识网络中唯一一台主机
		传输层：确认端口号
		应用层：可以自定义协议

	4. TCP状态转换: 数据传递期间:  ESTABLISEHED
		主动发起连接请求端:
			CLOSED --> 完成三次握手 --> ESTABLISEHED(数据通信状态) --> Dial()函数返回
		被动发起连接请求端:
			CLOSED --> 调用Accept()函数 --> LISTEN --> 完成三次握手 --> ESTABLISEHED(数据通信状态) --> Accept()函数返回

		主动关闭连接请求端: FIN_WAIT2(半关闭)，TIME_WAIT，2MSL 这些只会在主动关闭连接请求端
			 ESTABLISEHED --> FIN_WAIT2(半关闭) - TIME_WAIT --> 2MSL(确认最好一个ack被对端成功接收) --> CLOSED
		被动关闭连接请求端:
			ESTABLISEHED --> CLOSED
*/

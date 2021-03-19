package Sprocess

import (
	"net"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/util"
)

func Reading(conn net.Conn) {
	for {
		// 首先创建一个Transfer实例
		tf := util.NewTransfer(conn)
		// 接收来自客户端的数据包
		mes, err := tf.ReadPkg()
		if err != nil {
			return
		}
		// 根据不同类型的数据包进行分配给不同的方法
		switch mes.Type {
		case message.UserToDriverMesType:
			//todo 乘客向司机发送沟通消息
			NewSSmsProcess(conn).SendMesToAnother(mes)
		case message.ExitLoginMesType:
			//todo 处理退出信息
			NewSuerProcess(conn).ExitLogin(mes)
		case message.CreateOrderMesType:
			//todo 处理乘客下单信息
			NewSOrderProcess(conn).CreateOrder(mes)
		case message.CancelOrderMesType:
			//todo 处理乘客取消订单信息
			NewSOrderProcess(conn).CancelOrder(mes)
		default:
			// 处理无效信息
		}
	}
}

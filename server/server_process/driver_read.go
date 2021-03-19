package Sprocess

import (
	"net"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/util"
)

func DriverReading(conn net.Conn) {
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
		//todo 将司机的接单消息推送给乘客
		case message.EndOrderMesType:
			NewSOrderProcess(conn).EndOrder(mes)
		//todo 将司机的接单消息推送给乘客
		case message.DriverIsOrderMesType:
			NewSDriverSmsProcess(conn).SendMesIsOrder(mes)
		//todo 司机向乘客发送消息
		case message.DriverToUserMesType:
			NewSDriverSmsProcess(conn).SendMesToAnother(mes)
		case message.DriverExitLoginMesType:
			// 处理退出信息
			NewSdriverProcess(conn).ExitLogin(mes)
		default:
			// 处理无效信息
		}
	}
}

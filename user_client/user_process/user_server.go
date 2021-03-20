package Cprocess

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/redis"
	"trail_didi_3/pkg/util"
)

type Cserver struct {
	// Conn net.Conn
}

// func NewCserver(conn net.Conn) *Cserver {
// 	return &Cserver{
// 		Conn: conn,
// 	}
// }

func ShowLoginInterface() {
	var k = 1
	for {
		var key int
		if k == 1 {
			fmt.Println("----------------登录成功-------------")
		} else {
			fmt.Println("----------------继续下单-------------")
		}
		fmt.Println("\t\t\t 1 下打车订单")
		fmt.Println("\t\t\t 2 退出登录")
		fmt.Println("\t\t\t 请选择(1-2):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("-----下打车定单-----")
			Corder.CreateOrder()
		case 2:
			fmt.Println(" 2 退出登录")
			cp := NewCuserProcess()
			cp.Conn = Csms.Conn
			cp.ExitLogin(int(CurUser.Id))
			os.Exit(0)
		default:
			fmt.Println("无此功能")
			break
		}
		k++
	}

}

func ServerProcessMes(conn net.Conn) {
	//Start = make(chan bool)
	//select {
	//case Start <- true:
	//	//fmt.Println("解除阻塞执行协程里面for循环，监听服务端是否有消息推送过来")
	//}
	// 创建transfer实例，不断的读取服务端发来的数据包
	tf := util.NewTransfer(conn)
	for {
		resMes, err := tf.ReadPkg()
		if err != nil {
			// fmt.Println("服务器出错啦111")
			return
		}
		//解析数据包
		DataMes, err := tf.ParsePacket(resMes)
		if err != nil {
			// fmt.Println("服务器出错啦")
			return
		}
		// 读取消息类型
		switch resMes.Type {
		//todo 将司机结束订单的消息推送给乘客
		case message.ToUserEndOrderMesType:
			// 将数据转换成CurUserMes
			dialogMes, ok := DataMes.(message.ToUserEndOrder)
			if !ok {
				return
			}
			info := fmt.Sprintf("订单号为:\t%s\t的订单已经结束;结束司机名称:[%s]",
				dialogMes.OrderSn, dialogMes.DriverName)
			fmt.Println(info)
		case message.ResCreateOrderMesType:
			//todo 获取乘客创建订单的返回信息
			//将司机接单返回信息存进数据库
			rdConn := redis.GetInstance()
			defer rdConn.Close()
			redis.AddResultInfo(rdConn, message.ResCreateOrderStatus, strconv.Itoa(int(CurUser.Id)), resMes.Data)
		case message.ResCancelOrderMesType:
			//todo 获取乘客取消订单的返回信息
			dialogOtherMes, ok := DataMes.(message.ResCancelOrder)
			if !ok {
				return
			}
			//将司机接单返回信息存进数据库
			rdConn := redis.GetInstance()
			defer rdConn.Close()
			var cancelOrderKey = dialogOtherMes.OrderSn + "" + strconv.Itoa(int(CurUser.Id))
			redis.AddResultInfo(rdConn, message.ResCancelOrderStatus, cancelOrderKey, resMes.Data)
		case message.DriverPushUserIsOrderMesType:
			// 将数据转换成CurUserMes
			dialogMes, ok := DataMes.(message.DriverPushUserIsOrder)
			if !ok {
				return
			}
			//todo 将司机接单成功的消息推送给乘客
			info := fmt.Sprintf("订单号为:\t%s\t接单成功;司机名称:[%s];司机ID为:[%d]",
				dialogMes.OrderSn, dialogMes.DriverName, dialogMes.Order.DriverId)
			fmt.Println(info)
		case message.DriverToUserMesType:
			//todo 司机向用户发送消息
			dialogOtherMes, ok := DataMes.(message.DriverToUserMes)
			if !ok {
				return
			}
			// 将数据输出在控制台
			info := fmt.Sprintf("司机%s[%d]对你%s[%d]说:%s",
				dialogOtherMes.DriverName, dialogOtherMes.Id, dialogOtherMes.OtherUserName, dialogOtherMes.OtherUserId, dialogOtherMes.Dialog)
			fmt.Println(info)
		default:
			break
		}
	}
}

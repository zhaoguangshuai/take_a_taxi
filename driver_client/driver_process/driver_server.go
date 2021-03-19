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

func ShowLoginInterface() {
	var l = 1
	for {
		var key int
		var order_sn string
		if l == 1 {
			fmt.Println("----------------登录成功-------------")
		} else {
			fmt.Println("----------------继续接单-------------")
		}
		fmt.Println("\t\t\t 1 接单")
		fmt.Println("\t\t\t 2 退出登录")
		fmt.Println("\t\t\t 请选择(1-2):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("----------------司机正在接单-------------")
			fmt.Println("请输入要接单的订单号")
			fmt.Scanf("%s\n", &order_sn)
			//todo 编辑司机接单的业务逻辑
			DriverCsms.SendDriverIsOrder(order_sn)
		case 2:
			fmt.Println(" 2 退出登录")
			cp := NewCuserProcess()
			cp.Conn = DriverCsms.Conn
			cp.ExitLogin(CurDriver.Id)
			os.Exit(0)
		default:
			fmt.Println("无此功能")
			break
		}
		l++
	}
}

func ServerProcessMes(conn net.Conn) {
	// 创建transfer实例，不断的读取服务端发来的数据包
	tf := util.NewTransfer(conn)
	// tf := util.NewTransfer(this.Conn)
	for {
		resMes, err := tf.ReadPkg()
		// fmt.Println("resMes=", resMes.Type)
		// fmt.Println("resMes=", resMes.Data)
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
		//todo 获取乘客想我推送取消订单的信息
		case message.ToDriverCancelOrderMesType:
			// 将数据转换成CurUserMes
			orderMes, ok := DataMes.(message.ToDriverCancelOrder)
			if !ok {
				return
			}
			// 将数据输出在控制台
			info := fmt.Sprintf("订单号为:\t%s\t的订单，乘客已取消",
				orderMes.OrderSn)
			fmt.Println(info)
		case message.OrderPushMesType:
			// 将数据转换成CurUserMes
			orderMes, ok := DataMes.(message.Order)
			if !ok {
				return
			}
			// 将数据输出在控制台
			info := fmt.Sprintf("订单号为:\t%s\t待接单，起点=》终点:%s=》%s;该乘客ID:[%d]",
				orderMes.OrderSn, orderMes.StartAddress, orderMes.EndAddress, orderMes.UserId)
			fmt.Println(info)
		case message.UserToDriverMesType:
			// 将数据转换成CurUserMes
			dialogOtherMes, ok := DataMes.(message.UserToDriverMes)
			if !ok {
				return
			}
			// 将数据输出在控制台
			info := fmt.Sprintf("乘客%s[%d]对你%s[%d]说:%s",
				dialogOtherMes.UserName, dialogOtherMes.UserId, dialogOtherMes.OtherDriverName, dialogOtherMes.OtherDriverId, dialogOtherMes.Dialog)
			fmt.Println(info)
		case message.ResDriverIsOrderMesType:
			//todo 获取司机接单成功的返回信息
			dialogOtherMes, ok := DataMes.(message.ResOrderMes)
			if !ok {
				return
			}
			//将司机接单返回信息存进数据库
			rdConn := redis.GetInstance()
			defer rdConn.Close()
			var status_key = dialogOtherMes.OrderSn + "" + strconv.Itoa(CurDriver.Id)
			redis.AddResultInfo(rdConn, message.ResDriverIsOrderStatus, status_key, resMes.Data)
		case message.ResEndOrderMesType:
			//todo 获取司结束成功的返回信息
			dialogOtherMes, ok := DataMes.(message.ResEndOrder)
			if !ok {
				return
			}
			//将司机接单返回信息存进数据库
			rdConn := redis.GetInstance()
			defer rdConn.Close()
			var status_key = dialogOtherMes.OrderSn + "" + strconv.Itoa(CurDriver.Id)
			redis.AddResultInfo(rdConn, message.ResDriverEndOrderStatus, status_key, resMes.Data)
		default:
			break
		}
	}
}

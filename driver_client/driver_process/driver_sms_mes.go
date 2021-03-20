package Cprocess

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/redis"
	"trail_didi_3/pkg/util"
)

type driverCsmsMes struct {
	Conn net.Conn
}

var DriverCsms *driverCsmsMes

func NewDriverCsmsMes(conn net.Conn) *driverCsmsMes {
	return &driverCsmsMes{
		Conn: conn,
	}
}

//todo 司机给乘客发送接单消息
func (this *driverCsmsMes) SendDriverIsOrder(order_sn string) {
	var driverIsOrder message.DriverIsOrder

	driverIsOrder.Order_sn = order_sn
	driverIsOrder.Driver = CurDriver

	// 创建一个Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 将dialogOtherMes 封装成 Message
	mes, err := tf.EncapsulationPacket(message.DriverIsOrderMesType, "driver", driverIsOrder)
	if err != nil {
		return
	}
	// 发送数据包给服务端
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}

	//todo 等待司机服务端返回接单成功的消息，做下一步处理
	time.Sleep(1 * time.Second)
	rdConn := redis.GetInstance()
	defer rdConn.Close()
	var status_key = order_sn + "" + strconv.Itoa(int(CurDriver.Id))
	data, ok := redis.SelectResultInfo(rdConn, message.ResDriverIsOrderStatus, status_key)
	if !ok {
		fmt.Println("从数据库获取接单同步返回信息失败")
		return
	}

	switch data.Code {
	case message.CodeIsOrderStatusOne:
		fmt.Println("该订单号的订单不存在，请重新接单")
		return
	case message.CodeIsOrderStatusTwo:
		fmt.Println("该订单已经被取消，请重新接单")
		return
	case message.CodeIsOrderStatusThree:
		fmt.Println("该订单已被其他司机接单，请重新接单")
		return
	case message.CodeIsOrderStatusFour:
		fmt.Println("该订单已经完成，请重新接单")
		return
	case message.CodeIsOrderStatusFive:
		this.isOrderSuccess()
	}
}

func (this *driverCsmsMes) isOrderSuccess() {
	var i = 1
	for {
		var key int
		if i == 1 {
			fmt.Println("----------------司机接单成功-------------")
		} else {
			fmt.Println("----------------与乘客沟通中-------------")
		}
		fmt.Println("\t\t\t 1 与乘客沟通")
		fmt.Println("\t\t\t 2 退出登录")
		fmt.Println("\t\t\t 3 订单完成")
		fmt.Println("\t\t\t 4 继续接单")
		fmt.Println("\t\t\t 请选择(1-4):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			this.SendDialogToAnother()
			break
		case 2:
			fmt.Println(" 2 退出登录")
			cp := NewCuserProcess()
			cp.Conn = DriverCsms.Conn
			cp.ExitLogin(int(CurDriver.Id))
			os.Exit(0)
		case 3:
			this.EndOrder()
			return
		case 4:
			return
		default:
			fmt.Println("无此功能")
			break
		}
		i++
	}
}

//todo 司机完成订单
func (this *driverCsmsMes) EndOrder() {
	fmt.Println("----------------司机正在结束订单-------------")
	var order_sn string
	fmt.Println("请输入要结束的订单的订单号")
	fmt.Scanf("%s\n", &order_sn)

	var endOrderInfo message.EndOrder
	endOrderInfo.OrderSn = order_sn
	endOrderInfo.DriverId = CurDriver.Id

	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 将dialogMes封装
	mes, err := tf.EncapsulationPacket(message.EndOrderMesType, "driver", endOrderInfo)
	if err != nil {
		return
	}

	// 发送数据包给服务端
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}

	//todo 等待司机服务端返回结束订单成功的消息
	time.Sleep(1 * time.Second)
	rdConn := redis.GetInstance()
	defer rdConn.Close()
	var status_key = order_sn + "" + strconv.Itoa(int(CurDriver.Id))
	data, ok := redis.SelectResultInfo(rdConn, message.ResDriverEndOrderStatus, status_key)
	if !ok {
		fmt.Println("从数据库获取结束订单同步返回信息失败")
		return
	}

	if data.Code == message.CodeEndOrderSuccessful {
		fmt.Println("订单结束成功")
	} else {
		fmt.Println("订单结束失败，请联系客服")
	}
}

//todo 司机给乘客发送消息
func (this *driverCsmsMes) SendDialogToAnother() {
	var otherUserId uint64
	var dialog string
	var DriverToUserMes message.DriverToUserMes
	fmt.Println("请输入你想沟通的乘客Id:")
	fmt.Scanf("%d\n", &otherUserId)
	// 获取乘客id 的名字,并判断对方是否在线
	//otherUser, ok := Cdrivers.SearchOnlineUser(otherUserId)
	//if !ok {
	//	fmt.Println("该用户不在线")
	//}

	fmt.Println("请输入你想对该乘客说的话")
	fmt.Scanf("%s\n", &dialog)

	// 将数据添加到dialogOtherUserMes 中
	DriverToUserMes.Dialog = dialog
	DriverToUserMes.OtherUserId = otherUserId
	DriverToUserMes.Driver = CurDriver

	// 创建一个Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 将dialogOtherMes 封装成 Message
	mes, err := tf.EncapsulationPacket(message.DriverToUserMesType, "driver", DriverToUserMes)
	if err != nil {
		return
	}

	// 发送数据包给服务端
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}
}

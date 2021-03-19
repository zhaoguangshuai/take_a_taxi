package Sprocess

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/redis"
	"trail_didi_3/pkg/util"
)

var rwMutex *sync.RWMutex //定义读写锁的全局变量
type SdriverSmsProcess struct {
	Conn   net.Conn
	Driver message.Driver
}

func NewSDriverSmsProcess(conn net.Conn) *SdriverSmsProcess {
	return &SdriverSmsProcess{
		Conn: conn,
	}
}

//todo 将司机接单的消息发送给乘客
func (this *SdriverSmsProcess) SendMesIsOrder(smsMes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 解析私聊数据包
	dataMes, err := tf.ParsePacket(smsMes)
	if err != nil {
		return
	}
	// 进行类型断言
	driverIsOrder, ok := dataMes.(message.DriverIsOrder)

	//定义返回信息
	var driverResMes message.ResOrderMes

	//获取redis服务连接信息
	rdConn := redis.GetInstance()
	defer rdConn.Close()

	rwMutex = new(sync.RWMutex) //获取该包的对象
	//开始上锁
	rwMutex.Lock()
	//todo 获取订单信息
	orderInfo, ok := redis.SelectOrderInfo(rdConn, message.OrderInfoKey, driverIsOrder.Order_sn)
	if !ok {
		fmt.Println("该订单不存在")
		driverResMes.Code = message.CodeIsOrderStatusOne
	}
	orderStatus := orderInfo.OrderStatus
	if orderStatus == message.OrderStatusOne { //代表司机接单成功
		//修改订单状态，修改订单接单司机
		orderInfo.OrderStatus = message.OrderStatusTwo
		orderInfo.DriverId = driverIsOrder.Driver.Id
		data, _ := json.Marshal(orderInfo)
		redis.AddOrder(rdConn, message.OrderInfoKey, orderInfo.OrderSn, string(data))

		//todo 司机向该乘客发送接单成功消息
		// 获取接收方的连接数据
		sp, ok := SMUSER.OnlineUsers[orderInfo.UserId]
		if !ok {
			return
		}
		var driverPushUserIsOrder message.DriverPushUserIsOrder
		driverPushUserIsOrder.Driver = driverIsOrder.Driver
		driverPushUserIsOrder.Order = orderInfo
		driverPushUserIsOrder.User = sp.User

		// 创建接收方的Transfer实例
		tfSp := util.NewTransfer(sp.Conn)
		// 将数据封装起来
		resMes, err := tfSp.EncapsulationPacket(message.DriverPushUserIsOrderMesType, "driver", driverPushUserIsOrder)
		if err != nil {
			return
		}
		//todo 司机向该乘客发送接单成功消息
		err = tfSp.WritePkg(resMes)
		if err != nil {
			return
		}

		driverResMes.Code = message.CodeIsOrderStatusFive
	}
	//结束解锁
	rwMutex.Unlock()

	switch orderStatus {
	case message.OrderStatusTwo:
		driverResMes.Code = message.CodeIsOrderStatusThree
	case message.OrderStatusThree:
		driverResMes.Code = message.CodeIsOrderStatusFour
	case message.OrderStatusFour:
		driverResMes.Code = message.CodeIsOrderStatusTwo
	}
	driverResMes.Order = orderInfo
	//todo 向司机客户端返回接单成功
	ResMes, err := tf.EncapsulationPacket(message.ResDriverIsOrderMesType, "driver", driverResMes)
	if err != nil {
		return
	}
	//todo 将数据包发送回司机
	err = tf.WritePkg(ResMes)
	if err != nil {
		return
	}

}

func (this *SdriverSmsProcess) SendMesToAnother(smsMes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 解析私聊数据包
	dataMes, err := tf.ParsePacket(smsMes)
	if err != nil {
		return
	}
	// fmt.Println("this.COnn", this.Conn)
	// fmt.Println("this.User=", this.User)
	// 进行类型断言
	dialogOtherUserMes, ok := dataMes.(message.DriverToUserMes)
	if !ok {
		return
	}
	//todo 获取乘客（接受信息者）的连接数据
	sp, ok := SMUSER.OnlineUsers[dialogOtherUserMes.OtherUserId]
	if !ok {
		return
	}
	//将乘客名称加入到信息中
	dialogOtherUserMes.OtherUserName = sp.UserName
	// fmt.Printf("用户%s[%d]对你%s[%d]说:%s",)
	// 创建接收方的Transfer实例
	tfSp := util.NewTransfer(sp.Conn)
	// fmt.Println("other Conn", sp.Conn)
	// fmt.Println("Other COnn", sp.User)
	// 将数据封装起来
	resMes, err := tfSp.EncapsulationPacket(message.DriverToUserMesType, "driver", dialogOtherUserMes)
	if err != nil {
		return
	}
	// 向客户端发送数据
	err = tfSp.WritePkg(resMes)
	if err != nil {
		return
	}
}

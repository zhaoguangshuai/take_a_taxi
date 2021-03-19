package Sprocess

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
	"trail_didi_3/pkg/helper"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/redis"
	"trail_didi_3/pkg/util"
)

type SOrderProcess struct {
	Conn net.Conn
	User message.User
}

func NewSOrderProcess(conn net.Conn) *SOrderProcess {
	return &SOrderProcess{
		Conn: conn,
	}
}

func (this *SOrderProcess) CreateOrder(smsMes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 解析私聊数据包
	dataMes, err := tf.ParsePacket(smsMes)
	if err != nil {
		return
	}
	//todo 进行类型断言,获取司机传过来的订单信息
	dialogOtherUserMes, ok := dataMes.(message.DialogOrderInfoMes)
	if !ok {
		return
	}
	//todo 正式创建订单
	var orderInfo message.Order
	var orderSn = helper.GetOrderSn()
	orderInfo.OrderSn = orderSn
	orderInfo.UserId = dialogOtherUserMes.Order.UserId
	orderInfo.StartAddress = dialogOtherUserMes.StartAddress
	orderInfo.EndAddress = dialogOtherUserMes.EndAddress
	orderInfo.OrderStatus = message.OrderStatusOne
	orderInfo.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	// 获取数据库的连接,并读取
	rdConn := redis.GetInstance()
	defer rdConn.Close()

	data, err := json.Marshal(orderInfo)
	if err != nil {
		fmt.Println("订单信息序列化失败111")
	}
	ok = redis.AddOrder(rdConn, message.OrderInfoKey, orderSn, string(data))

	// 定义返回消息类型
	var resOrderMes message.ResOrderMes

	if ok { //todo 订单创建成功
		//todo 向司机推送订单信息，好让司机抢单
		NewSSmsProcess(this.Conn).SendMesToAllDriver(orderInfo)
		resOrderMes.Order = orderInfo
		resOrderMes.Code = message.CodeOrderCreateSuccess
	} else { //todo 订单创建失败
		resOrderMes.Code = message.CodeOrderCreateFail
	}

	// 封装resRegisterMes
	resMes, err := tf.EncapsulationPacket(message.ResCreateOrderMesType, "user", resOrderMes)
	if err != nil {
		fmt.Println(err)
	}

	//发送数据包
	err = tf.WritePkg(resMes)
	if err != nil {
		fmt.Println("server err", err)
		return
	}

}

//todo 乘客取消订单
func (this *SOrderProcess) CancelOrder(smsMes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 解析私聊数据包
	dataMes, err := tf.ParsePacket(smsMes)
	if err != nil {
		return
	}
	//todo 进行类型断言
	requestInfo, ok := dataMes.(message.CancelOrder)
	if !ok {
		return
	}
	//todo 正式取消订单
	// 获取数据库的连接,并读取
	rdConn := redis.GetInstance()
	defer rdConn.Close()

	orderInfo, _ := redis.SelectOrderInfo(rdConn, message.OrderInfoKey, requestInfo.OrderSn)

	// 定义返回消息类型
	var resOrderMes message.ResCancelOrder
	resOrderMes.OrderSn = requestInfo.OrderSn
	//todo 属于已接单状态，并且是自己的订单 才可以取消
	if orderInfo.OrderStatus == message.OrderStatusTwo && orderInfo.UserId == requestInfo.UserId {
		//修改订单状态为取消
		orderInfo.OrderStatus = message.OrderStatusFour
		data, _ := json.Marshal(orderInfo)
		redis.AddOrder(rdConn, message.OrderInfoKey, orderInfo.OrderSn, string(data))

		//todo 乘客向该司机推送取消订单的消息
		// 获取接收方的连接数据
		sp, ok := SMDRIVER.OnlineDrivers[orderInfo.DriverId]
		if !ok {
			return
		}
		var driverPushUserIsOrder message.ToDriverCancelOrder
		driverPushUserIsOrder.Order = orderInfo
		driverPushUserIsOrder.User = this.User

		// 创建接收方的Transfer实例
		tfSp := util.NewTransfer(sp.Conn)
		// 将数据封装起来
		resMes, err := tfSp.EncapsulationPacket(message.ToDriverCancelOrderMesType, "user", driverPushUserIsOrder)
		if err != nil {
			return
		}
		//todo 乘客向该司机推送取消订单的消息
		err = tfSp.WritePkg(resMes)
		if err != nil {
			return
		}
		//取消成功
		resOrderMes.Code = message.CodeCancelOrderSuccessful
		//todo 属于待接单状态，并且是自己的订单，才可以取消
	} else if orderInfo.OrderStatus == message.OrderStatusOne && orderInfo.UserId == requestInfo.UserId {
		//修改订单状态为取消
		orderInfo.OrderStatus = message.OrderStatusFour
		data, _ := json.Marshal(orderInfo)
		redis.AddOrder(rdConn, message.OrderInfoKey, orderInfo.OrderSn, string(data))
		//取消成功
		resOrderMes.Code = message.CodeCancelOrderSuccessful
	} else { //其他情况取消失败
		resOrderMes.Code = message.CodeCancelOrderFailure
	}

	// 封装resRegisterMes
	resMes, err := tf.EncapsulationPacket(message.ResCancelOrderMesType, "user", resOrderMes)
	if err != nil {
		fmt.Println(err)
	}

	//发送数据包
	err = tf.WritePkg(resMes)
	if err != nil {
		fmt.Println("server err", err)
		return
	}

}

//todo 司机结束订单
func (this *SOrderProcess) EndOrder(smsMes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 解析私聊数据包
	dataMes, err := tf.ParsePacket(smsMes)
	if err != nil {
		return
	}
	//todo 进行类型断言
	requestInfo, ok := dataMes.(message.EndOrder)
	if !ok {
		return
	}
	//todo 正式结束订单
	// 获取数据库的连接,并读取
	rdConn := redis.GetInstance()
	defer rdConn.Close()

	orderInfo, _ := redis.SelectOrderInfo(rdConn, message.OrderInfoKey, requestInfo.OrderSn)

	// 定义返回消息类型
	var resOrderMes message.ResEndOrder
	resOrderMes.OrderSn = requestInfo.OrderSn
	//todo 属于已接单状态，并且是自己接的订单 才可以改为已完成
	if orderInfo.OrderStatus == message.OrderStatusTwo && orderInfo.DriverId == requestInfo.DriverId {
		//修改订单状态为取消
		orderInfo.OrderStatus = message.OrderStatusThree
		data, _ := json.Marshal(orderInfo)
		redis.AddOrder(rdConn, message.OrderInfoKey, orderInfo.OrderSn, string(data))

		//todo 司机向该乘客推送结束订单的消息
		// 获取接收方的连接数据
		sp, ok := SMUSER.OnlineUsers[orderInfo.UserId]
		if !ok {
			return
		}
		var driverPushUserEndOrder message.ToUserEndOrder
		driverPushUserEndOrder.Order = orderInfo
		driverPushUserEndOrder.Driver = SMDRIVER.OnlineDrivers[requestInfo.DriverId].Driver

		// 创建接收方的Transfer实例
		tfSp := util.NewTransfer(sp.Conn)
		// 将数据封装起来
		resMes, err := tfSp.EncapsulationPacket(message.ToUserEndOrderMesType, "driver", driverPushUserEndOrder)
		if err != nil {
			return
		}
		//todo 司机向该乘客推送结束订单的消息
		err = tfSp.WritePkg(resMes)
		if err != nil {
			return
		}
		//结束成功
		resOrderMes.Code = message.CodeEndOrderSuccessful
	} else { //其他情况结束失败
		resOrderMes.Code = message.CodeEndOrderFailure
	}
	resOrderMes.OrderSn = orderInfo.OrderSn

	// 封装resRegisterMes
	resMes, err := tf.EncapsulationPacket(message.ResEndOrderMesType, "driver", resOrderMes)
	if err != nil {
		fmt.Println(err)
	}

	//发送数据包
	err = tf.WritePkg(resMes)
	if err != nil {
		fmt.Println("server err", err)
		return
	}

}

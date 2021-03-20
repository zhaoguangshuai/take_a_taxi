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

type CuserOrder struct {
	Conn net.Conn
}

var Corder *CuserOrder

func NewCuserOrder(conn net.Conn) *CuserOrder {
	return &CuserOrder{
		Conn: conn,
	}
}

func (this *CuserOrder) CreateOrder() {
	fmt.Println("请输入起点：")
	var startAddress string
	fmt.Scanf("%s\n", &startAddress)

	fmt.Println("请输入终点：")
	var endAddress string
	fmt.Scanf("%s\n", &endAddress)

	// 定义变量，将数据发送给服务器
	var orderMes message.DialogOrderInfoMes
	//todo 将订单信息封装到orderMes中
	orderMes.User = CurUser
	orderMes.Order.UserId = CurUser.Id
	orderMes.StartAddress = startAddress
	orderMes.EndAddress = endAddress

	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 将dialogMes封装
	mes, err := tf.EncapsulationPacket(message.CreateOrderMesType, "user", orderMes)
	if err != nil {
		return
	}

	// 发送创建订单的数据包
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}

	//todo 等待司机服务端返回接单成功的消息，做下一步处理
	time.Sleep(1 * time.Second)
	rdConn := redis.GetInstance()
	defer rdConn.Close()
	data, ok := redis.SelectResultInfo(rdConn, message.ResCreateOrderStatus, strconv.Itoa(int(CurUser.Id)))
	if !ok {
		fmt.Println("从数据库获取创建订单同步返回信息失败")
		return
	}

	if data.Code == message.CodeOrderCreateSuccess {
		var key int
		fmt.Printf("-------订单创建成功,订单号为=》%s\t等待司机接单--------\n", data.OrderSn)
		fmt.Println("\t\t\t 1 取消订单")
		fmt.Println("\t\t\t 2 退出登录")
		fmt.Println("\t\t\t 3 等待司机接单后与司机沟通")
		fmt.Println("\t\t\t 4 继续下单")
		fmt.Println("\t\t\t 请选择(1-4):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			this.CancelOrder()
			return
		case 2:
			fmt.Println(" 2 退出登录")
			cp := NewCuserProcess()
			cp.Conn = Csms.Conn
			cp.ExitLogin(int(CurUser.Id))
			os.Exit(0)
		case 3:
			Csms.SendDialogToAnother()
		case 4:
			return
		default:
			fmt.Println("无此功能")
			break
		}
	} else {
		fmt.Println("订单创建失败,请重新创建订单")
	}
}

func (this *CuserOrder) CancelOrder() {
	fmt.Println("----------------乘客正在取消订单-------------")
	var order_sn string
	fmt.Println("请输入要取消的订单的订单号")
	fmt.Scanf("%s\n", &order_sn)

	var cancelOrderInfo message.CancelOrder
	cancelOrderInfo.OrderSn = order_sn
	cancelOrderInfo.UserId = CurUser.Id

	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 将dialogMes封装
	mes, err := tf.EncapsulationPacket(message.CancelOrderMesType, "user", cancelOrderInfo)
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
	var status_key = order_sn + "" + strconv.Itoa(int(CurUser.Id))
	data, ok := redis.SelectResultInfo(rdConn, message.ResCancelOrderStatus, status_key)
	if !ok {
		fmt.Println("从数据库获取取消订单同步返回信息失败")
		return
	}

	if data.Code == message.CodeCancelOrderSuccessful {
		fmt.Println("订单取消成功")
	} else {
		fmt.Println("订单取消失败,请联系客服")
	}
}

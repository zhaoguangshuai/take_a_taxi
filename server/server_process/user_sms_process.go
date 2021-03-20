package Sprocess

import (
	"fmt"
	"net"
	"trail_didi_3/models/chat_message"
	"trail_didi_3/models/order"
	"trail_didi_3/models/user"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/orm"
	"trail_didi_3/pkg/util"
)

type SSmsProcess struct {
	Conn net.Conn
	User user.User
}

func NewSSmsProcess(conn net.Conn) *SSmsProcess {
	return &SSmsProcess{
		Conn: conn,
	}
}

//todo 乘客向司机发送消息沟通
func (this *SSmsProcess) SendMesToAnother(smsMes message.Message) {
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
	dialogOtherUserMes, ok := dataMes.(message.UserToDriverMes)
	if !ok {
		return
	}
	//todo 获取司机（接受信息者）的连接数据
	sp, ok := SMDRIVER.OnlineDrivers[int(dialogOtherUserMes.OtherDriverId)]
	fmt.Println("乘客发送的消息信息", dialogOtherUserMes)
	if !ok {
		return
	}
	//将乘客名称加入到信息中
	dialogOtherUserMes.OtherDriverName = sp.DriverName
	// 创建接收方的Transfer实例
	tfSp := util.NewTransfer(sp.Conn)
	// fmt.Println("other Conn", sp.Conn)
	// fmt.Println("Other COnn", sp.User)
	// 将数据封装起来
	resMes, err := tfSp.EncapsulationPacket(message.UserToDriverMesType, "user", dialogOtherUserMes)
	if err != nil {
		return
	}
	// 向客户端发送数据
	err = tfSp.WritePkg(resMes)
	if err != nil {
		return
	}
	//todo 将发送的消息存储到数据库
	ChatMessage := chat_message.ChatMessage{}
	ChatMessage.DriveId = dialogOtherUserMes.OtherDriverId
	ChatMessage.UserId = dialogOtherUserMes.User.Id
	ChatMessage.Content = dialogOtherUserMes.Dialog
	orm.GetInstance().Create(&ChatMessage)
}

/*
用户下单后向所有司机推送消息
*/
func (this *SSmsProcess) SendMesToAllDriver(orderInfo order.Order) {
	//todo 向所有在线司机发送订单信息
	drives := SMDRIVER.GetAllOnlineUser()
	for _, sp := range drives {
		sTf := util.NewTransfer(sp.Conn)
		// 封装消息数据
		resMes, err := sTf.EncapsulationPacket(message.OrderPushMesType, "user", orderInfo)
		if err != nil {
			return
		}
		// 发送消息数据包给所有在线用户
		err = sTf.WritePkg(resMes)
		if err != nil {
			return
		}
	}

}

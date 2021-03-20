package Cprocess

import (
	"fmt"
	"net"
	"os"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/util"
)

type CsmsMes struct {
	Conn net.Conn
}

var Csms *CsmsMes

func NewCsmsMes(conn net.Conn) *CsmsMes {
	return &CsmsMes{
		Conn: conn,
	}
}

//todo 乘客与司机进行沟通
func (this *CsmsMes) SendDialogToAnother() {
	var otherDriverId uint64
	var dialog string
	var dialogUserToDriverMes message.UserToDriverMes
	fmt.Println("请输入你想沟通的司机Id:")
	fmt.Scanf("%d\n", &otherDriverId)
	// 获取司机id 的名字,并判断对方是否在线
	//otherUser, ok := Cusers.SearchOnlineUser(otherUserId)
	//if !ok {
	//	fmt.Println("该用户不在线")
	//}

	fmt.Println("请输入你想对司机说的话")
	fmt.Scanf("%s\n", &dialog)

	// 将数据添加到dialogOtherUserMes 中
	dialogUserToDriverMes.Dialog = dialog               //说的内容
	dialogUserToDriverMes.OtherDriverId = otherDriverId //需要发给那个司机
	dialogUserToDriverMes.User = CurUser                //当前说话的乘客信息

	// 创建一个Transfer实例
	tf := util.NewTransfer(this.Conn)

	// 将dialogOtherMes 封装成 Message
	mes, err := tf.EncapsulationPacket(message.UserToDriverMesType, "user", dialogUserToDriverMes)
	if err != nil {
		return
	}
	// 发送数据包给服务端
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}

	for {
		var key int
		fmt.Println("----------------与司机沟通中-------------")
		fmt.Println("\t\t\t 1 取消订单")
		fmt.Println("\t\t\t 2 退出登录")
		fmt.Println("\t\t\t 3 继续与司机沟通")
		fmt.Println("\t\t\t 4 继续下单")
		fmt.Println("\t\t\t 请选择(1-4):")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			Corder.CancelOrder()
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
	}
}

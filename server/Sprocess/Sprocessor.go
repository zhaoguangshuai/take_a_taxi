package Sprocess

import (
	"fmt"
	"net"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/util"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) CreatProcess() {
	// 首先创建一个Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 接收来自客户端的数据包
	mes, err := tf.ReadPkg()
	if err != nil {
		return
	}

	if mes.UserType == "user" { //todo 乘客的登陆注册
		switch mes.Type {
		case message.LoginMesType:
			// 处理登录信息
			fmt.Println("处理登录信息")
			NewSuerProcess(this.Conn).LoginCheck(mes)

		case message.RegisterMesType:
			// 处理注册信息
			fmt.Println("处理注册信息")
			NewSuerProcess(this.Conn).Register(mes)
		//case message.DialogMesType:
		//	// 处理注册信息
		//	fmt.Println("处理用户的消息")
		//	NewSSmsProcess(this.Conn).SendMesToAll(mes)
		default:
			// 处理无效信息
		}
	} else { //todo 司机的登陆注册
		switch mes.Type {
		case message.DriverLoginMesType:
			// 处理司机登录信息
			fmt.Println("处理登录信息")
			NewSdriverProcess(this.Conn).DriverLoginCheck(mes)

		case message.DriverRegisterMesType:
			// 处理司机注册信息
			fmt.Println("处理司机注册信息")
			NewSdriverProcess(this.Conn).DriverRegister(mes)
		//case message.DialogMesType:
		//	// 处理注册信息
		//	fmt.Println("处理用户的消息")
		//	NewSSmsProcess(this.Conn).SendMesToAll(mes)
		default:
			// 处理无效信息
		}
	}

}

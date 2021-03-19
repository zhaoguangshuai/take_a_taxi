package Cprocess

import (
	"fmt"
	"net"
	config2 "trail_didi_3/pkg/config"
	"trail_didi_3/pkg/helper"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/util"
)

//var Start chan bool

type CuserProcess struct {
	Conn net.Conn
}

func NewCuserProcess() *CuserProcess {
	return &CuserProcess{}
}

func (this *CuserProcess) Login(userId int, userPwd string) (err error) {
	// 连接到服务器
	this.Conn, err = net.Dial("tcp", config2.GetString("app.url"))
	if err != nil {
		return
	}
	// 延时关闭连接
	defer this.Conn.Close()
	// 创建登录消息实例
	loginMes := message.LoginMes{
		User: message.User{
			UserId:  userId,
			UserPwd: userPwd,
		},
	}
	//  创键Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 对数据进行封装
	mes, err := tf.EncapsulationPacket(message.LoginMesType, "user", loginMes)
	if err != nil {
		return
	}
	// 向服务端发送数据包
	tf.WritePkg(mes)
	// 接收服务端的响应
	resMes, err := tf.ReadPkg()
	if err != nil {
		return
	}
	// 解析服务端发回的数据包
	//fmt.Println(resMes)
	data, err := tf.ParsePacket(resMes)
	if err != nil {
		return
	}
	// 对data 进行类型转换
	resLoginMes, ok := data.(message.ResLoginMes)
	if !ok {
		return
	}
	if resLoginMes.Code == message.CodeLoginSuccessful {
		fmt.Println("登录成功")
		CurUser = resLoginMes.User

		// 创建消息实例
		Csms = NewCsmsMes(this.Conn)

		//创建订单实例
		Corder = NewCuserOrder(this.Conn)

		// 如果登录成功,则显示登录界面,并启动协程来接收服务端发来的数据
		go ServerProcessMes(this.Conn)

		ShowLoginInterface()

	} else if resLoginMes.Code == message.CodeLoginFailure {
		fmt.Println("登录失败")
	} else if resLoginMes.Code == message.CodeHaveNotRegister {
		fmt.Println("未注册用户")
	}

	return
}

/**
乘客注册
*/
func (this *CuserProcess) Register(userId int, userPwd, userName string) (err error) {
	fmt.Println("用户注册")
	//todo 连接服务器
	this.Conn, err = net.Dial("tcp", config2.GetString("app.url"))
	if err != nil {
		return
	}
	//todo 延时关闭连接
	defer this.Conn.Close()

	//todo 创建RegisterMes实例
	registerMes := message.RegisterMes{
		User: message.User{
			UserId:   userId,
			UserPwd:  helper.Md5V2(userPwd),
			UserName: userName,
		},
	}
	//todo 创建Transfer实例
	tf := util.NewTransfer(this.Conn)
	//todo 将registerMes 封装
	mes, err := tf.EncapsulationPacket(message.RegisterMesType, "user", registerMes)
	if err != nil {
		return
	}
	// fmt.Println("Register    90 mes.type=", mes.Type)
	//todo 发送数据包
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}

	//  接收服务端返回的数据包
	resMes, err := tf.ReadPkg()
	if err != nil {
		return
	}
	// 解析数据包
	data, err := tf.ParsePacket(resMes)
	if err != nil {
		return
	}
	resRisterMes, ok := data.(message.ResRegisterMes)
	if !ok {
		return
	}
	if resRisterMes.Code == message.CodeRegisterSuccessful {
		fmt.Println("注册成功")
	} else {
		fmt.Println("注册失败")

	}
	return

}
func (this *CuserProcess) ExitLogin(userId int) {
	// 创建Transfer 实例
	tf := util.NewTransfer(this.Conn)
	var ExitLoginMes message.ExitLoginMes
	ExitLoginMes.User = CurUser

	// 将退出登录信息封装到mes
	mes, err := tf.EncapsulationPacket(message.ExitLoginMesType, "user", ExitLoginMes)
	if err != nil {
		return
	}
	// 将数据包发送给服务端
	err = tf.WritePkg(mes)
	if err != nil {
		return
	}
	this.Conn.Close()
}

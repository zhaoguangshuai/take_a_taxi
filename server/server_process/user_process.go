package Sprocess

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net"
	"trail_didi_3/models/user"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/orm"
	"trail_didi_3/pkg/util"
)

type SuserProcess struct {
	Conn net.Conn
	user.User
}

func NewSuerProcess(conn net.Conn) *SuserProcess {
	return &SuserProcess{
		Conn: conn,
	}
}

//todo 乘客登陆
func (this *SuserProcess) LoginCheck(mes message.Message) {
	//defer this.Conn.Close()
	// 解析数据包
	tf := util.NewTransfer(this.Conn)
	data, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	// 转化数据
	loginMes, ok := data.(message.LoginMes)
	if !ok {
		return
	}

	// 定义变量, 用于返回结果给客户端
	var resLoginMes message.ResLoginMes

	//todo 从数据库中获取登陆用户信息
	where := make(map[string]interface{})
	where["user_account"] = loginMes.UserAccount
	User := user.User{}
	orm.GetInstance().Where(where).One(&User)
	if User.Id > 0 {
		// 验证数据是否合法
		if loginMes.UserAccount == User.UserAccount && bcrypt.CompareHashAndPassword([]byte(User.UserPwd), []byte(loginMes.UserPwd)) == nil {
			resLoginMes.Code = message.CodeLoginSuccessful
			// 用户登录成功，则将数据存到在线用户管理中
			this.Id = User.Id
			this.UserName = User.UserName
			this.UserStatus = true
			//todo 添登陆的乘客添加到一个map容器中，保存在线的乘客列表
			SMUSER.AddOnlineUser(this)
			resLoginMes.User = user.User{
				Id:         User.Id,
				UserName:   User.UserName,
				UserStatus: User.UserStatus,
			}
		} else {
			resLoginMes.Code = message.CodeLoginFailure
		}
	} else {
		resLoginMes.Code = message.CodeHaveNotRegister
	}

	fmt.Println("登陆乘客集合11", SMUSER.GetAllOnlineUser())

	// 封装数据包成mes
	ResMes, err := tf.EncapsulationPacket(message.ResLoginMesType, "user", resLoginMes)
	if err != nil {
		return
	}
	//todo 将数据包发送回乘客
	err = tf.WritePkg(ResMes)
	if err != nil {
		return
	}

	//if this.UserStatus {
	// 给客户端发送已经在线的乘客信息
	//	SM.NotifyOthersUser()
	//}
	//todo 服务端开启一个协程时刻监听该登陆乘客发送到服务端的信息
	go Reading(this.Conn)

}
func (this *SuserProcess) Register(mes message.Message) {
	defer this.Conn.Close()
	tf := util.NewTransfer(this.Conn)
	// 解析数据包
	data, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	registerMes, ok := data.(message.RegisterMes)
	if !ok {
		return
	}
	// 定义返回消息类型
	var resRegisterMes message.ResRegisterMes

	// 获取数据库的连接,并读取
	where := make(map[string]interface{})
	where["user_account"] = registerMes.UserAccount
	User := user.User{}
	orm.GetInstance().Where(where).One(&User)
	if User.Id > 0 {
		resRegisterMes.Code = message.CodeRegisterFailure
	} else {
		User.UserAccount = registerMes.UserAccount
		pwd, _ := bcrypt.GenerateFromPassword([]byte(registerMes.UserPwd), bcrypt.DefaultCost)
		User.UserPwd = string(pwd)
		User.UserName = registerMes.UserName
		orm.GetInstance().Create(&User)
		if User.Id > 0 {
			resRegisterMes.Code = message.CodeRegisterSuccessful
		} else {
			resRegisterMes.Code = message.CodeRegisterFailure
		}
	}

	// 封装resRegisterMes
	resMes, err := tf.EncapsulationPacket(message.ResRegisterMesType, "user", resRegisterMes)
	if err != nil {
		return
	}
	//发送数据包
	err = tf.WritePkg(resMes)
	if err != nil {
		fmt.Println("server err", err)
		return
	}
	// fmt.Println("返回状态码成功")
}
func (this *SuserProcess) ExitLogin(mes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 解析退出登录的数据包
	dataMes, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	exitLoginMes, ok := dataMes.(message.ExitLoginMes)
	if !ok {
		return
	}
	// 将要退出登录的用户从在线列表中删除
	delete(SMUSER.OnlineUsers, int(exitLoginMes.Id))
	this.Conn.Close()
	fmt.Println("登陆乘客集合22", SMUSER.GetAllOnlineUser())

}

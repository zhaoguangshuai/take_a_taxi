package Sprocess

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net"
	"trail_didi_3/models/driver"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/orm"
	"trail_didi_3/pkg/util"
)

type SdriverProcess struct {
	Conn net.Conn
	driver.Driver
}

func NewSdriverProcess(conn net.Conn) *SdriverProcess {
	return &SdriverProcess{
		Conn: conn,
	}
}

//todo 司机登陆
func (this *SdriverProcess) DriverLoginCheck(mes message.Message) {
	// defer this.Conn.Close()
	// 解析数据包
	tf := util.NewTransfer(this.Conn)
	data, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	// 转化数据
	loginMes, ok := data.(message.DriverLoginMes)
	if !ok {
		return
	}

	// 定义变量, 用于返回结果给客户端
	var resLoginMes message.DriverResLoginMes

	// 从数据中获取数据
	where := make(map[string]interface{})
	where["driver_account"] = loginMes.DriverAccount
	Driver := driver.Driver{}
	orm.GetInstance().Where(where).One(&Driver)

	if Driver.Id > 0 {
		// 验证数据是否合法
		if loginMes.DriverAccount == Driver.DriverAccount && bcrypt.CompareHashAndPassword([]byte(Driver.DriverPwd), []byte(loginMes.DriverPwd)) == nil {
			resLoginMes.Code = message.CodeLoginSuccessful
			// 用户登录成功，则将数据存到在线用户管理中
			this.Id = Driver.Id
			this.DriverName = Driver.DriverName
			this.DriverStatus = true
			//todo 添登陆的司机添加到一个map容器中，保存在线的司机列表
			SMDRIVER.AddOnlineUser(this)
			resLoginMes.Driver = driver.Driver{
				Id:           Driver.Id,
				DriverName:   Driver.DriverName,
				DriverStatus: Driver.DriverStatus,
			}
		} else {
			resLoginMes.Code = message.CodeLoginFailure
		}
	} else {
		resLoginMes.Code = message.CodeHaveNotRegister
	}

	fmt.Println("登陆司机集合11", SMDRIVER.GetAllOnlineUser())

	// 封装数据包成mes
	ResMes, err := tf.EncapsulationPacket(message.DriverResLoginMesType, "driver", resLoginMes)
	if err != nil {
		return
	}
	//todo 将数据包发送回司机
	err = tf.WritePkg(ResMes)
	if err != nil {
		return
	}

	//todo 服务端开启一个协程时刻监听该司机发送到服务端的信息
	go DriverReading(this.Conn)

}
func (this *SdriverProcess) DriverRegister(mes message.Message) {
	defer this.Conn.Close()
	tf := util.NewTransfer(this.Conn)
	// 解析数据包
	data, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	registerMes, ok := data.(message.DriverRegisterMes)
	if !ok {
		return
	}
	// 定义返回消息类型
	var resRegisterMes message.DriverResRegisterMes

	// 从数据中获取数据
	where := make(map[string]interface{})
	where["driver_account"] = registerMes.DriverAccount
	Driver := driver.Driver{}
	orm.GetInstance().Where(where).One(&Driver)
	if Driver.Id > 0 {
		resRegisterMes.Code = message.CodeRegisterFailure
	} else {
		Driver.DriverAccount = registerMes.DriverAccount
		pwd, _ := bcrypt.GenerateFromPassword([]byte(registerMes.DriverPwd), bcrypt.DefaultCost)
		Driver.DriverPwd = string(pwd)
		Driver.DriverName = registerMes.DriverName
		orm.GetInstance().Create(&Driver)
		if Driver.Id > 0 {
			resRegisterMes.Code = message.CodeRegisterSuccessful
		} else {
			resRegisterMes.Code = message.CodeRegisterFailure
		}
	}
	// 封装resRegisterMes
	resMes, err := tf.EncapsulationPacket(message.DriverResRegisterMesType, "driver", resRegisterMes)
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
func (this *SdriverProcess) ExitLogin(mes message.Message) {
	// 创建Transfer实例
	tf := util.NewTransfer(this.Conn)
	// 解析退出登录的数据包
	dataMes, err := tf.ParsePacket(mes)
	if err != nil {
		return
	}
	exitLoginMes, ok := dataMes.(message.DriverExitLoginMes)
	if !ok {
		return
	}
	// 将要退出登录的用户从在线列表中删除
	delete(SMDRIVER.OnlineDrivers, int(exitLoginMes.Id))
	this.Conn.Close()
	fmt.Println(SMDRIVER.OnlineDrivers)
	fmt.Println("登陆司机集合22", SMDRIVER.GetAllOnlineUser())

}

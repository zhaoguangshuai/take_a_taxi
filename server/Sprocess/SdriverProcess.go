package Sprocess

import (
	"fmt"
	"net"
	"trail_didi_3/pkg/helper"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/redis"
	"trail_didi_3/pkg/util"
)

type SdriverProcess struct {
	Conn net.Conn
	message.Driver
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
	rdConn := redis.GetInstance()
	defer rdConn.Close()
	//todo 从数据库中获取登陆司机信息
	driver, ok := redis.SelectDriverInfo(rdConn, message.DriverInfoKey, loginMes.Id)
	if ok {
		// 验证数据是否合法
		if loginMes.Id == driver.Id && helper.Md5V2(loginMes.DriverPwd) == driver.DriverPwd {
			resLoginMes.Code = message.CodeLoginSuccessful
			// 用户登录成功，则将数据存到在线用户管理中
			this.Id = driver.Id
			this.DriverName = driver.DriverName
			this.DriverStatus = true
			//todo 添登陆的司机添加到一个map容器中，保存在线的司机列表
			SMDRIVER.AddOnlineUser(this)
			resLoginMes.Driver = message.Driver{
				Id:           driver.Id,
				DriverName:   driver.DriverName,
				DriverStatus: driver.DriverStatus,
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

	// 获取数据库的连接,并读取
	rdConn := redis.GetInstance()
	defer rdConn.Close()
	//todo 将乘客的注册信息添加到数据库
	_, ok = redis.SelectDriverInfo(rdConn, message.DriverInfoKey, registerMes.Id)

	if ok {
		resRegisterMes.Code = message.CodeRegisterFailure
	} else {
		ok := redis.AddDriver(rdConn, message.DriverInfoKey, registerMes.Id, mes.Data)
		if ok {
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
	delete(SMDRIVER.OnlineDrivers, exitLoginMes.Id)
	this.Conn.Close()
	fmt.Println(SMDRIVER.OnlineDrivers)
	fmt.Println("登陆司机集合22", SMDRIVER.GetAllOnlineUser())

}

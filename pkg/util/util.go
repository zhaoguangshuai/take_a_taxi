package util

import (
	"trail_didi_3/pkg/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type transfer struct {
	Conn net.Conn
	buf  [4096]byte
}

func NewTransfer(conn net.Conn) *transfer {
	return &transfer{
		Conn: conn,
	}
}

// 封装数据包
func (this *transfer) EncapsulationPacket(Type message.MesType,UserType string, Data interface{}) (mes message.Message, err error) {
	mes.UserType = UserType
	switch Type {
	case message.ResEndOrderMesType:
		mes.Type = message.ResEndOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.ResEndOrder
		loginMes, ok := Data.(message.ResEndOrder)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ToUserEndOrderMesType:
		mes.Type = message.ToUserEndOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.ToUserEndOrder
		loginMes, ok := Data.(message.ToUserEndOrder)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.EndOrderMesType:
		mes.Type = message.EndOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.EndOrder
		loginMes, ok := Data.(message.EndOrder)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ToDriverCancelOrderMesType:
		mes.Type = message.ToDriverCancelOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.ToDriverCancelOrder
		loginMes, ok := Data.(message.ToDriverCancelOrder)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ResCancelOrderMesType:
		mes.Type = message.ResCancelOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.ResCancelOrder
		loginMes, ok := Data.(message.ResCancelOrder)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.CancelOrderMesType:
		mes.Type = message.CancelOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.CancelOrder
		loginMes, ok := Data.(message.CancelOrder)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.UserToDriverMesType:
		mes.Type = message.UserToDriverMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.UserToDriverMes
		loginMes, ok := Data.(message.UserToDriverMes)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.DriverToUserMesType:
		mes.Type = message.DriverToUserMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.DriverToUserMes
		loginMes, ok := Data.(message.DriverToUserMes)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ResDriverIsOrderMesType:
		mes.Type = message.ResDriverIsOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.ResOrderMes
		loginMes, ok := Data.(message.ResOrderMes)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.DriverPushUserIsOrderMesType:
		mes.Type = message.DriverPushUserIsOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.DriverPushUserIsOrder
		loginMes, ok := Data.(message.DriverPushUserIsOrder)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.DriverIsOrderMesType:
		mes.Type = message.DriverIsOrderMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.DriverIsOrder
		loginMes, ok := Data.(message.DriverIsOrder)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.LoginMesType:
		mes.Type = message.LoginMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.LoginMes
		loginMes, ok := Data.(message.LoginMes)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ResLoginMesType:
		mes.Type = message.ResLoginMesType
		// 登录验证码
		var resLoginMes message.ResLoginMes
		resLoginMes, ok := Data.(message.ResLoginMes)
		if !ok {
			return
		}
		// resLoginMes 序列化
		data, err := json.Marshal(resLoginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.RegisterMesType:
		mes.Type = message.RegisterMesType
		// 注册消息
		var registerMes message.RegisterMes
		registerMes, ok := Data.(message.RegisterMes)
		if !ok {
			return
		}

		// registerMes 序列化
		data, err := json.Marshal(registerMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ResRegisterMesType:
		mes.Type = message.ResRegisterMesType
		// 注册验证码
		var resRegisterMes message.ResRegisterMes
		resRegisterMes, ok := Data.(message.ResRegisterMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(resRegisterMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ResStatusMesType:
		mes.Type = message.ResStatusMesType
		// 注册验证码
		var resStatusMes message.ResStatusMes
		resStatusMes, ok := Data.(message.ResStatusMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(resStatusMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成11", mes)
		return mes, err
	case message.DialogMesType:
		mes.Type = message.DialogMesType
		// 注册验证码
		var dialogMes message.DialogMes
		dialogMes, ok := Data.(message.DialogMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(dialogMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("dialogMes封装数据包完成", mes)
		return mes, err
	case message.DialogOtherUserMesType:
		mes.Type = message.DialogOtherUserMesType
		var dialogOtherUserMes message.DialogOtherUserMes
		dialogOtherUserMes, ok := Data.(message.DialogOtherUserMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(dialogOtherUserMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ExitLoginMesType:
		mes.Type = message.ExitLoginMesType
		// 注册验证码
		var exitLoginMes message.ExitLoginMes
		exitLoginMes, ok := Data.(message.ExitLoginMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(exitLoginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.DriverExitLoginMesType:
		mes.Type = message.DriverExitLoginMesType
		// 注册验证码
		var exitLoginMes message.DriverExitLoginMes
		exitLoginMes, ok := Data.(message.DriverExitLoginMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(exitLoginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.DriverLoginMesType:
		mes.Type = message.DriverLoginMesType
		// 登录消息
		// 数据包类型转换
		var loginMes message.DriverLoginMes
		loginMes, ok := Data.(message.DriverLoginMes)
		if !ok {
			return
		}
		// 将loginMes 序列化
		data, err := json.Marshal(loginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.DriverResLoginMesType:
		mes.Type = message.DriverResLoginMesType
		// 登录验证码
		var resLoginMes message.DriverResLoginMes
		resLoginMes, ok := Data.(message.DriverResLoginMes)
		if !ok {
			return
		}
		// resLoginMes 序列化
		data, err := json.Marshal(resLoginMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.DriverRegisterMesType:
		mes.Type = message.DriverRegisterMesType
		// 注册消息
		var registerMes message.DriverRegisterMes
		registerMes, ok := Data.(message.DriverRegisterMes)
		if !ok {
			return
		}

		// registerMes 序列化
		data, err := json.Marshal(registerMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.DriverResRegisterMesType:
		mes.Type = message.ResRegisterMesType
		// 注册验证码
		var resRegisterMes message.DriverResRegisterMes
		resRegisterMes, ok := Data.(message.DriverResRegisterMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(resRegisterMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.CreateOrderMesType:
		mes.Type = message.CreateOrderMesType
		var orderMes message.DialogOrderInfoMes
		orderMes, ok := Data.(message.DialogOrderInfoMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(orderMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.OrderPushMesType:
		mes.Type = message.OrderPushMesType
		var orderMes message.Order
		orderMes, ok := Data.(message.Order)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(orderMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	case message.ResCreateOrderMesType:
		mes.Type = message.ResCreateOrderMesType
		var orderMes message.ResOrderMes
		orderMes, ok := Data.(message.ResOrderMes)
		if !ok {
			return
		}
		// resRegisterMes 序列化
		data, err := json.Marshal(orderMes)
		if err != nil {
			return mes, err
		}
		mes.Data = string(data)
		// fmt.Println("封装数据包完成", mes)
		return mes, err
	default:
		return
	}
}

// 解析数据包
func (this *transfer) ParsePacket(mes message.Message) (DataMes interface{}, err error) {
	// fmt.Println("数据包解析", mes.Type)
	switch mes.Type {
	case message.ToUserEndOrderMesType:
		// 登录消息
		var loginMes message.ToUserEndOrder
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.ResEndOrderMesType:
		// 登录消息
		var loginMes message.ResEndOrder
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.EndOrderMesType:
		// 登录消息
		var loginMes message.EndOrder
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.ToDriverCancelOrderMesType:
		// 登录消息
		var loginMes message.ToDriverCancelOrder
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.ResCancelOrderMesType:
		// 登录消息
		var loginMes message.ResCancelOrder
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.CancelOrderMesType:
		// 登录消息
		var loginMes message.CancelOrder
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.UserToDriverMesType:
		// 登录消息
		var loginMes message.UserToDriverMes
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.DriverToUserMesType:
		// 登录消息
		var loginMes message.DriverToUserMes
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.ResDriverIsOrderMesType:
		// 登录消息
		var loginMes message.ResOrderMes
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.DriverPushUserIsOrderMesType:
		// 登录消息
		var loginMes message.DriverPushUserIsOrder
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.DriverIsOrderMesType:
		// 登录消息
		var loginMes message.DriverIsOrder
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.LoginMesType:
		// 登录消息
		var loginMes message.LoginMes
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.ResLoginMesType:
		// 登录验证码
		var resLoginMes message.ResLoginMes
		// resLoginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &resLoginMes)
		if err != nil {
			return resLoginMes, err
		}
		// fmt.Println("数据包解析成功")
		return resLoginMes, err
	case message.RegisterMesType:
		// 注册消息
		var registerMes message.RegisterMes
		// registerMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &registerMes)
		if err != nil {
			return registerMes, err
		}
		// fmt.Println("数据包解析成功")
		return registerMes, err
	case message.ResRegisterMesType:
		// 注册验证码
		var resRegisterMes message.ResRegisterMes
		// resRegisterMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &resRegisterMes)
		if err != nil {
			return resRegisterMes, err
		}
		// fmt.Println("数据包解析成功")
		return resRegisterMes, err
	case message.DialogMesType:
		// 注册验证码
		var dialogMes message.DialogMes
		// dialogMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &dialogMes)
		if err != nil {
			return dialogMes, err
		}
		// fmt.Println("DialogMes数据包解析成功")
		return dialogMes, err
	case message.ResStatusMesType:
		// 注册验证码
		var resStatusMes message.ResStatusMes
		// resStatusMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &resStatusMes)
		if err != nil {
			return resStatusMes, err
		}
		// fmt.Println("数据包解析成功")
		return resStatusMes, err
	case message.DialogOtherUserMesType:
		// 注册验证码
		var dialogOtherUserMes message.DialogOtherUserMes
		// DialogOtherUserMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &dialogOtherUserMes)
		if err != nil {
			return dialogOtherUserMes, err
		}
		// fmt.Println("数据包解析成功")
		return dialogOtherUserMes, err
	case message.ExitLoginMesType:
		// 注册验证码
		var exitLoginMes message.ExitLoginMes
		// DialogOtherUserMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &exitLoginMes)
		if err != nil {
			return exitLoginMes, err
		}
		// fmt.Println("数据包解析成功")
		return exitLoginMes, err
	case message.DriverExitLoginMesType:
		// 注册验证码
		var exitLoginMes message.DriverExitLoginMes
		// DialogOtherUserMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &exitLoginMes)
		if err != nil {
			return exitLoginMes, err
		}
		// fmt.Println("数据包解析成功")
		return exitLoginMes, err
	case message.DriverLoginMesType:
		// 登录消息
		var loginMes message.DriverLoginMes
		// 将loginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &loginMes)
		if err != nil {
			return loginMes, err
		}
		// fmt.Println("数据包解析成功")
		return loginMes, err
	case message.DriverResLoginMesType:
		// 登录验证码
		var resLoginMes message.DriverResLoginMes
		// resLoginMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &resLoginMes)
		if err != nil {
			return resLoginMes, err
		}
		// fmt.Println("数据包解析成功")
		return resLoginMes, err
	case message.DriverRegisterMesType:
		// 注册消息
		var registerMes message.DriverRegisterMes
		// registerMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &registerMes)
		if err != nil {
			return registerMes, err
		}
		// fmt.Println("数据包解析成功")
		return registerMes, err
	case message.DriverResRegisterMesType:
		// 注册验证码
		var resRegisterMes message.DriverResRegisterMes
		// resRegisterMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &resRegisterMes)
		if err != nil {
			return resRegisterMes, err
		}
		// fmt.Println("数据包解析成功")
		return resRegisterMes, err
	case message.CreateOrderMesType:
		//todo 解析乘客下单信息数据
		var orderMes message.DialogOrderInfoMes
		// DialogOtherUserMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &orderMes)
		if err != nil {
			return orderMes, err
		}
		// fmt.Println("数据包解析成功")
		return orderMes, err
	case message.OrderPushMesType:
		//todo 解析乘客下单信息数据
		var orderMes message.Order
		// DialogOtherUserMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &orderMes)
		if err != nil {
			return orderMes, err
		}
		// fmt.Println("数据包解析成功")
		return orderMes, err
	case message.ResCreateOrderMesType:
		//todo 解析乘客下单信息数据
		var orderMes message.ResOrderMes
		// DialogOtherUserMes 反序列化
		err := json.Unmarshal([]byte(mes.Data), &orderMes)
		if err != nil {
			return orderMes, err
		}
		// fmt.Println("数据包解析成功")
		return orderMes, err
	default:
		fmt.Println("dialogOtherUserMes数据包解析失败")
		return
	}
}

// 读取数据包
func (this *transfer) ReadPkg() (resMes message.Message, err error) {
	// fmt.Println("读取数据包00", resMes.Type)
	// fmt.Println("读取数据包00", resMes.Data)
	// 先接收数据包的长度(字节)
	_, err = this.Conn.Read(this.buf[:4])
	if err != nil {
		return
	}
	pkgLen := this.getPkgLen(this.buf[:4])
	// fmt.Println("数据包长度", pkgLen)
	// 读取数据包
	n, err := this.Conn.Read(this.buf[:pkgLen])
	if n != pkgLen {
		// return resMes, fmt.Errorf("数据包长度不一致")
	}
	// 将数据包进行反序列化
	err = json.Unmarshal(this.buf[:pkgLen], &resMes)
	if err != nil {
		return
	}
	//fmt.Println("读取数据包成功")
	return
}

// 发送数据包
func (this *transfer) WritePkg(mes message.Message) (err error) {
	// fmt.Println("发送数据", mes.Data)
	// 先将mes 进行反序列化
	data, err := json.Marshal(mes)
	if err != nil {
		return
	}
	// 获取数据包的长度字节
	buf := this.setPkgLen(data)
	// 将数据包的大小发送给对方
	_, err = this.Conn.Write(buf)
	if err != nil {
		return
	}
	// 发送数据包给对方
	_, err = this.Conn.Write(data)
	if err != nil {
		return
	}
	// fmt.Println("发送数据包成功")
	return
}
func (this *transfer) setPkgLen(data []byte) []byte {
	// 先定义一个uint32变量
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// fmt.Println("数据包长度", pkgLen)
	binary.BigEndian.PutUint32(this.buf[:4], pkgLen)
	return this.buf[:4]

}
func (this *transfer) getPkgLen(data []byte) int {
	resPkgLen := binary.BigEndian.Uint32(this.buf[:4])
	// fmt.Println("resPkgLen", resPkgLen)
	return int(resPkgLen)
}

package message

const (
	LoginMesType             = "LoginMes"
	ResLoginMesType          = "LoginResMes"
	RegisterMesType          = "RegisterMes"
	ResRegisterMesType       = "ResRegisterMes"
	DialogMesType            = "DialogMes"
	ResStatusMesType         = "ResStatusMes"
	DialogOtherUserMesType   = "DialogOtherUserMes"
	ExitLoginMesType         = "ExitLoginMes"
	DriverLoginMesType       = "DriverLoginMes"
	DriverResLoginMesType    = "DriverLoginResMes"
	DriverRegisterMesType    = "DriverRegisterMes"
	DriverResRegisterMesType = "DriverResRegisterMes"
	DriverExitLoginMesType   = "DriverExitLoginMes"
	CreateOrderMesType       = "CreateOrderMes"
	ResCreateOrderMesType    = "ResCreateOrderMes"
	OrderPushMesType         = "OrderPushMes"
	DriverIsOrderMesType = "DriverIsOrderMes"//司机接单成功通知
	ResDriverIsOrderMesType = "ResDriverIsOrderMes"//服务端返回司机信息
	DriverPushUserIsOrderMesType = "DriverPushUserIsOrderMes"//司机向乘客推送接单成功信息
	DriverToUserMesType = "DriverToUserMes"//司机向乘客发消息
	UserToDriverMesType = "UserToDriverMes"//乘客向司机发消息
	CancelOrderMesType       = "CancelOrderMes"//乘客取消订单
	ResCancelOrderMesType    = "ResCancelOrderMes"//服务端返回信息
	ToDriverCancelOrderMesType = "ToDriverCancelOrderMes"//乘客向司机推送取消订单的消息
	EndOrderMesType       = "EndOrderMes"//司机结束订单
	ResEndOrderMesType    = "ResEndOrderMes"//服务端返回信息
	ToUserEndOrderMesType = "ToUserEndOrderMes"//司机向乘客推送结束订单消息
)

// 数据库key值
const (
	DatabaseKey   = "users"
	DriverInfoKey = "drivers"
	OrderInfoKey  = "order"
	ResDriverIsOrderStatus = "res_driver_is_order_status"
	ResDriverEndOrderStatus = "res_driver_end_order_status"
	ResCancelOrderStatus = "res_cancel_order_status"
	ResCreateOrderStatus = "res_create_order_status"
)

// 状态码协议
const (
	CodeLoginSuccessful    = 102
	CodeLoginFailure       = 105
	CodeRegisterSuccessful = 202
	CodeRegisterFailure    = 205
	CodeHaveNotRegister    = 400
	CodeOrderCreateSuccess     = 210
	CodeOrderCreateFail        = 501
	CodeIsOrderStatusOne = 410//订单不存在
	CodeIsOrderStatusTwo = 411//订单已取消
	CodeIsOrderStatusThree = 412//订单已经接单
	CodeIsOrderStatusFour = 413//订单已经完成
	CodeIsOrderStatusFive = 414//接单成功
	CodeCancelOrderSuccessful = 415//订单取消成功
	CodeCancelOrderFailure = 416//订单取消失败
	CodeEndOrderSuccessful = 417//订单结束成功
	CodeEndOrderFailure = 418//订单结束失败
)

// 订单状态码
const (
	OrderStatusOne   = 1 //待接单
	OrderStatusTwo   = 2 //已接单
	OrderStatusThree = 3 //已完成
	OrderStatusFour  = 4 //已取消
)

type EndOrder struct {
	OrderSn string `json:"order_sn"`
	DriverId int `json:"driver_id"`
}

type ResEndOrder struct {
	OrderSn string `json:"order_sn"`
	ResMes
}

type ToUserEndOrder struct { //司机向乘客推送结束订单消息
	Order
	Driver
}

type CancelOrder struct {
	OrderSn string `json:"order_sn"`
	UserId int `json:"user_id"`
}

type ResCancelOrder struct {
	OrderSn string `json:"order_sn"`
	ResMes
}

type ToDriverCancelOrder struct { //乘客向司机推送取消订单的信息
	Order
	User
}

type ResRegisterMes struct {
	ResMes
}

// 协议：100 登录成功 ,105 密码错误,200 注册成功,300 用户存在,400 用户不存在
type ResMes struct {
	Code int `json:"code"`
}

type DriverIsOrder struct {
	Driver
	Order_sn string
}

type DriverPushUserIsOrder struct { //司机向乘客推送接单成功信息
	Driver
	Order
	User
}

type User struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus bool   `json:"userStatus"` // 在线就是true
}
type MesType string

// 定义消息类型
type Message struct {
	Type     string `json:"type"`     // 消息类型
	Data     string `json:"data"`     //数据的内容
	UserType string `json:"userType"` //发送消息的角色 user(代表乘客) driver(司机)
}

// 登录
type LoginMes struct {
	User
}

// 状态
type ResStatusMes struct {
	User
}

type ResLoginMes struct {
	ResMes
	User
}

type ExitLoginMes struct {
	User
}

// 注册
type RegisterMes struct {
	User
}

// 消息
type DialogMes struct {
	User
	Dialog string `json:"dialog"`
}
type DialogOtherUserMes struct {
	User
	OtherUserId int
	Dialog      string `json:"dialog"`
}

//乘客下单信息
type DialogOrderInfoMes struct {
	User
	Order
}

type Order struct {
	OrderSn      string `json:"order_sn"`
	UserId       int    `json:"user_id"`
	DriverId     int    `json:"driver_id"`
	StartAddress string `json:"start_address"`
	EndAddress   string `json:"end_address"`
	OrderStatus  int    `json:"order_status"`
	CreatedAt    string `json:"created_at"`
}

type Driver struct {
	Id     int `json:"id"`
	DriverPwd    string `json:"driver_pwd"`
	DriverName   string `json:"driver_name"`
	DriverStatus bool   `json:"driver_status"` // 在线就是true
}

// 登录
type DriverLoginMes struct {
	Driver
}

type DriverResLoginMes struct {
	ResMes
	Driver
}

type DriverExitLoginMes struct {
	Driver
}

// 注册
type DriverRegisterMes struct {
	Driver
}
type DriverResRegisterMes struct {
	ResMes
}

// 状态
type DriverResStatusMes struct {
	Driver
}

type ResOrderMes struct {
	ResMes
	Order
}

//司机向乘客沟通
type DriverToUserMes struct {
	Driver
	OtherUserId int
	OtherUserName string
	Dialog      string `json:"dialog"`
}
//乘客向司机沟通
type UserToDriverMes struct {
	User
	OtherDriverId int
	OtherDriverName string
	Dialog      string `json:"dialog"`
}

//司机与用户互相发送消息的数据结构
//type DialogMes struct {
//
//	Dialog string `json:"dialog"`
//}
//type DialogOtherUserMes struct {
//	User
//	OtherUserId int
//	Dialog      string `json:"dialog"`
//}

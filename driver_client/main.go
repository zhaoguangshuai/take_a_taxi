package main

import (
	"fmt"
	"os"
	"trail_didi_3/config"
	Cprocess "trail_didi_3/driver_client/CProcess"
)

func init() {
	//初始化配置
	config.Initialize()
}
func main() {
	var DriverId int
	var DriverPwd string
	var DriverName string
	for {
		// 用于记录用户输入的选项
		var key int
		fmt.Println("----------------欢迎登陆博雅司机系统------------")
		fmt.Println("\t\t\t 1 登陆司机系统")
		fmt.Println("\t\t\t 2 注册司机账号")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("-----登录-----")
			fmt.Println("请输入账号：")
			fmt.Scanf("%d\n", &DriverId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &DriverPwd)
			// 调用司机登录方法
			err := Cprocess.NewCuserProcess().DriverLogin(DriverId, DriverPwd)
			if err != nil {
				fmt.Println("服务器未开启,请联系后台人员")
			}
		case 2:
			fmt.Println("司机注册")
			fmt.Println("请输入账号：")
			fmt.Scanf("%d\n", &DriverId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &DriverPwd)
			fmt.Println("请输入昵称：")
			fmt.Scanf("%s\n", &DriverName)
			// 调用司机注册方法
			err := Cprocess.NewCuserProcess().DriverRegister(DriverId, DriverPwd, DriverName)
			if err != nil {
				fmt.Println("服务器未开启,请联系后台人员")
			}

		case 3:
			fmt.Println("你退出了系统")
			os.Exit(0)
		default:
			fmt.Println("没有这个选项")
		}
	}

}

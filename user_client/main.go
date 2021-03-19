package main

import (
	"fmt"
	"os"
	"trail_didi_3/config"
	Cprocess "trail_didi_3/user_client/CProcess"
)

func init() {
	//初始化配置
	config.Initialize()
}

func main() {
	var userId int
	var userPwd string
	var userName string
	for {
		//todo 用于记录用户输入的选项
		var key int
		fmt.Println("----------------欢迎进入博雅打车系统------------")
		fmt.Println("\t\t\t 1 登陆打车系统")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("-----登录-----")
			fmt.Println("请输入账号：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			// 调用登录方法
			err := Cprocess.NewCuserProcess().Login(userId, userPwd)
			if err != nil {
				fmt.Println("服务器未开启,请联系后台人员")
			}
		case 2:
			fmt.Println("用户注册")
			fmt.Println("请输入账号：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入昵称：")
			fmt.Scanf("%s\n", &userName)
			//todo 调用乘客注册方法
			err := Cprocess.NewCuserProcess().Register(userId, userPwd, userName)
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

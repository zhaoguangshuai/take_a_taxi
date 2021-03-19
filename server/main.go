// server
package main

import (
	"fmt"
	"net"
	"trail_didi_3/config"
	config2 "trail_didi_3/pkg/config"
	"trail_didi_3/server/Sprocess"
)

func creatProcessor(conn net.Conn) {
	//todo 创建一个任务管家，专门用于分配任务
	processor := Sprocess.Processor{
		Conn: conn,
	}
	fmt.Println("客户端连接conn = ", conn)
	processor.CreatProcess()
}

func init() {
	//初始化配置
	config.Initialize()
	//todo 维护一个乘客管理集合
	Sprocess.NewSonlineUserMgr()
	//todo 维护一个司机管理集合
	Sprocess.NewSonlineDriverMgr()
}

func main() {
	//todo 监听端口 9999
	fmt.Println("服务器在端口 9999 开始监听")
	listen, err := net.Listen("tcp", config2.GetString("app.url"))
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}

	// 2 等待获取来自于客户端的连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		fmt.Println("address=", conn.LocalAddr())
		if err != nil {
			fmt.Println("listen.Accept err", err)
			return
		}
		// 每连接到一个客户端就启动一个协程为其服务
		go creatProcessor(conn)

	}

}

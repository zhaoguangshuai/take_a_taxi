package Cprocess

import (
	"fmt"
	"trail_didi_3/pkg/message"
)

var CurUser message.User

type CuserMgr struct {
	OnlineUsers map[int]message.User
}

var Cusers CuserMgr = NewCuserMgr()

func NewCuserMgr() CuserMgr {
	return CuserMgr{
		OnlineUsers: make(map[int]message.User, 1024),
	}
}

func (this *CuserMgr) AddOnlineUser(user message.User) {
	this.OnlineUsers[user.UserId] = user
}
func (this *CuserMgr) DelOnlineUser(userId int) {
	delete(this.OnlineUsers, userId)
}
func (this *CuserMgr) ModifyOnlineUser(user message.User) {
	this.AddOnlineUser(user)
}
func (this *CuserMgr) SearchOnlineUser(userId int) (user message.User, ok bool) {
	user, ok = this.OnlineUsers[userId]
	return
}
func (this *CuserMgr) OutputAllOnlineUser() {
	// 显示所有在线用户
	fmt.Println("用户ID\t用户名\t状态")
	for _, value := range this.OnlineUsers {
		var statusStr string
		if value.UserStatus {
			statusStr = "在线"
		} else {
			statusStr = "离线"
		}
		info := fmt.Sprintf("%d\t%s\t%s",
			value.UserId, value.UserName, statusStr)
		fmt.Println(info)
	}
}

package Cprocess

import (
	"trail_didi_3/models/user"
)

var CurUser user.User

type CuserMgr struct {
	OnlineUsers map[int]user.User
}

var Cusers CuserMgr = NewCuserMgr()

func NewCuserMgr() CuserMgr {
	return CuserMgr{
		OnlineUsers: make(map[int]user.User, 1024),
	}
}

func (this *CuserMgr) AddOnlineUser(user user.User) {
	this.OnlineUsers[int(user.Id)] = user
}
func (this *CuserMgr) DelOnlineUser(userId int) {
	delete(this.OnlineUsers, userId)
}
func (this *CuserMgr) ModifyOnlineUser(user user.User) {
	this.AddOnlineUser(user)
}
func (this *CuserMgr) SearchOnlineUser(userId int) (user user.User, ok bool) {
	user, ok = this.OnlineUsers[userId]
	return
}

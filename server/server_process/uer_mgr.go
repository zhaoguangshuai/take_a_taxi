package Sprocess

//todo 用于维护在线乘客用户
type SonlineUerMgr struct {
	OnlineUsers map[int]*SuserProcess
}

var SMUSER *SonlineUerMgr

func NewSonlineUserMgr() {
	SMUSER = &SonlineUerMgr{
		OnlineUsers: make(map[int]*SuserProcess, 1024),
	}
}

// 对在线用户进行增删改查操作
func (this *SonlineUerMgr) AddOnlineUser(sp *SuserProcess) {
	SMUSER.OnlineUsers[int(sp.Id)] = sp
}
func (this *SonlineUerMgr) DelOnlineUser(sp *SuserProcess) {
	delete(SMUSER.OnlineUsers, int(sp.Id))
}
func (this *SonlineUerMgr) ModifyOnlineUser(sp *SuserProcess) {
	this.AddOnlineUser(sp)
}
func (this *SonlineUerMgr) SeletcOnlineUser(sp *SuserProcess) (value *SuserProcess, ok bool) {
	value, ok = this.OnlineUsers[int(sp.Id)]
	return
}
func (this *SonlineUerMgr) GetAllOnlineUser() (users map[int]*SuserProcess) {
	return this.OnlineUsers
}

package Sprocess

//todo 用于维护在线乘客用户
type SonlineDriverMgr struct {
	OnlineDrivers map[int]*SdriverProcess
}

var SMDRIVER *SonlineDriverMgr

func NewSonlineDriverMgr() {
	SMDRIVER = &SonlineDriverMgr{
		OnlineDrivers: make(map[int]*SdriverProcess, 1024),
	}
}

// 对在线用户进行增删改查操作
func (this *SonlineDriverMgr) AddOnlineUser(sp *SdriverProcess) {
	SMDRIVER.OnlineDrivers[sp.Id] = sp
}
func (this *SonlineDriverMgr) DelOnlineUser(sp *SdriverProcess) {
	delete(SMDRIVER.OnlineDrivers, sp.Id)
}
func (this *SonlineDriverMgr) ModifyOnlineUser(sp *SdriverProcess) {
	this.AddOnlineUser(sp)
}
func (this *SonlineDriverMgr) SeletcOnlineUser(sp *SdriverProcess) (value *SdriverProcess, ok bool) {
	value, ok = this.OnlineDrivers[sp.Id]
	return
}
func (this *SonlineDriverMgr) GetAllOnlineUser() (users map[int]*SdriverProcess) {
	return this.OnlineDrivers
}

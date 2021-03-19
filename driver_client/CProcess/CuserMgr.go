package Cprocess

import (
	"trail_didi_3/pkg/message"
)

var CurDriver message.Driver

type CdriverMgr struct {
	OnlineDrivers map[int]message.Driver
}

var Cdrivers CdriverMgr = NewCdriverMgr()

func NewCdriverMgr() CdriverMgr {
	return CdriverMgr{
		OnlineDrivers: make(map[int]message.Driver, 1024),
	}
}

func (this *CdriverMgr) AddOnlineUser(driver message.Driver) {
	this.OnlineDrivers[driver.Id] = driver
}
func (this *CdriverMgr) DelOnlineUser(driverId int) {
	delete(this.OnlineDrivers, driverId)
}
func (this *CdriverMgr) ModifyOnlineUser(driver message.Driver) {
	this.AddOnlineUser(driver)
}
func (this *CdriverMgr) SearchOnlineUser(driverId int) (driver message.Driver, ok bool) {
	driver, ok = this.OnlineDrivers[driverId]
	return
}

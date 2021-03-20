package Cprocess

import (
	"trail_didi_3/models/driver"
)

var CurDriver driver.Driver

type CdriverMgr struct {
	OnlineDrivers map[int]driver.Driver
}

var Cdrivers CdriverMgr = NewCdriverMgr()

func NewCdriverMgr() CdriverMgr {
	return CdriverMgr{
		OnlineDrivers: make(map[int]driver.Driver, 1024),
	}
}

func (this *CdriverMgr) AddOnlineUser(driver driver.Driver) {
	this.OnlineDrivers[int(driver.Id)] = driver
}
func (this *CdriverMgr) DelOnlineUser(driverId int) {
	delete(this.OnlineDrivers, driverId)
}
func (this *CdriverMgr) ModifyOnlineUser(driver driver.Driver) {
	this.AddOnlineUser(driver)
}
func (this *CdriverMgr) SearchOnlineUser(driverId int) (driver driver.Driver, ok bool) {
	driver, ok = this.OnlineDrivers[driverId]
	return
}

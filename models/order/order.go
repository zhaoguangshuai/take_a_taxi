package order

import (
	"time"
	"trail_didi_3/models/driver"
	"trail_didi_3/models/user"
)

type Order struct {
	Id           uint64        `gorm:"column:id;primaryKey;autoIncrement;not null"json:"id"`
	OrderSn      string        `json:"order_sn"`
	UserId       uint64          `json:"user_id"`
	User         user.User     `gorm:"-"relationship:"belongTo"json:"user"`
	DriverId     uint64          `json:"driver_id"`
	driver       driver.Driver `gorm:"-"relationship:"belongTo"json:"driver"`
	StartAddress string        `json:"start_address"`
	EndAddress   string        `json:"end_address"`
	OrderStatus  uint8         `json:"order_status"`
	CreatedAt    time.Time     `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt    time.Time     `gorm:"column:updated_at"json:"updated_at"`
}

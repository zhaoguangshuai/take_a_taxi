package driver

import "time"

type Driver struct {
	Id            uint64    `json:"id"`
	DriverAccount string    `validate:"required,min:2,max:10"json:"driver_account"`
	DriverPwd     string    `validate:"required,min:2,max:10"json:"driver_pwd"`
	DriverName    string    `validate:"required"json:"driver_name"`
	DriverStatus  bool      `json:"driver_status"`
	CreatedAt     time.Time `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"json:"updated_at"`
}

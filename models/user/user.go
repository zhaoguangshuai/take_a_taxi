package user

import "time"

type User struct {
	Id          uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"json:"id"`
	UserAccount string    `validate:"required,min:2,max:10"json:"user_account"`
	UserPwd     string    `validate:"required,min:2,max:10"json:"user_pwd"`
	UserName    string    `validate:"required"json:"user_name"`
	UserStatus  bool      `json:"user_status"`
	CreatedAt   time.Time `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"json:"updated_at"`
}

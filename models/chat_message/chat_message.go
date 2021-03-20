package chat_message

import (
	"time"
	"trail_didi_3/models/driver"
	"trail_didi_3/models/user"
)

type ChatMessage struct {
	Id        uint64        `json:"id"gorm:"column:id;primaryKey;autoIncrement;not null"`
	UserId    uint64        `json:"user_id"`
	User      user.User     `gorm:"-"relationship:"belongTo"json:"user"`
	DriveId   uint64        `json:"drive_id"`
	Driver    driver.Driver `gorm:"-"relationship:"belongTo"json:"drive"`
	Content   string `json:"content"`
	CreatedAt time.Time `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"json:"updated_at"`
}

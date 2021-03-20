package bootstrap

import (
	"trail_didi_3/models/chat_message"
	"trail_didi_3/models/driver"
	"trail_didi_3/models/order"
	"trail_didi_3/models/user"
	"trail_didi_3/pkg/database"
)

var MigrateStruct map[string]interface{}

//初始化表结构体
func init()  {
	//初始化该全局变量
	MigrateStruct = make(map[string]interface{})
	MigrateStruct["driver"] = driver.Driver{}
	MigrateStruct["user"] = user.User{}
	MigrateStruct["order"] = order.Order{}
	MigrateStruct["chat_message"] = chat_message.ChatMessage{}

}

func AutoMigrate()  {
	database.SetMysqlDB()
	for _, v := range MigrateStruct{
		_ = database.DB.AutoMigrate(v)
	}
}

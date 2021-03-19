package redis

import (
	"github.com/go-redis/redis"
	"trail_didi_3/pkg/config"
	"trail_didi_3/pkg/types"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

//返回redis的连接实例
func GetInstance() *redis.Client {
	//todo 初始化线程池
	InitClient()
	return rdb
}

// 初始化连接
func InitClient()  {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.host"),
		//Password: config.GetString("redis.auth"), // no password set
		DB:       types.StringToInt(config.GetString("redis.db")),  // use default DB
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
}

package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"trail_didi_3/pkg/message"
	"trail_didi_3/pkg/types"
)

// 查询用户数据
func SelectDriverInfo(Redis *redis.Client, key string, filed int) (data message.Driver, flag bool) {
	res := Redis.HGet(key, types.IntToString(filed)).Val()
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		return
	}
	flag = true
	return data, flag
}

//添加用户数据
func AddDriver(Redis *redis.Client, key string, filed int, value string) bool {
	ok,_ := Redis.HSet(key, types.IntToString(filed), value).Result()
	return ok
}

// 查询用户数据
func SelectUserInfo(Redis *redis.Client, key string, filed int) (data message.User, flag bool) {
	res := Redis.HGet(key, types.IntToString(filed)).Val()
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		return
	}
	flag = true
	return data, flag
}

//添加用户数据
func AddUser(Redis *redis.Client, key string, filed int, value string) bool {
	ok,_ := Redis.HSet(key, types.IntToString(filed), value).Result()
	return ok
}

//添加订单
func AddOrder(Redis *redis.Client, key string, field string, value string) bool {
	ok,_ := Redis.HSet(key, field, value).Result()
	return ok
}

//查询订单数据
func SelectOrderInfo(Redis *redis.Client, key string, field string) (data message.Order, flag bool) {
	res := Redis.HGet(key, field).Val()
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		return
	}
	flag = true
	return data, flag
}
//添加客户端操作，服务端返回信息
func AddResultInfo(Redis *redis.Client, key string, field string, value string) bool {
	ok,_ := Redis.HSet(key, field, value).Result()
	return ok
}

//获取客户端操作，服务端返回信息
func SelectResultInfo(Redis *redis.Client, key string, field string) (data message.ResOrderMes, flag bool) {
	res := Redis.HGet(key, field).Val()
	err := json.Unmarshal([]byte(res), &data)
	if err != nil {
		return
	}
	flag = true
	return data, flag
}
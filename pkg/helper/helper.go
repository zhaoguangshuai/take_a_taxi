package helper

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func GetTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTimeTick32() int32{
	return int32(time.Now().Unix())
}

func GetFormatTime(time time.Time)string{
	return time.Format("20060102")
}

//todo 获取订单号
func GetOrderSn() string {
	date := GetFormatTime(time.Now())
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, GetTimeTick32(), r)
	return code
}

//todo 当前方法执行的包目录，及方法名称
func PrintMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
//todo 调用该方法所在的地方
func PrintCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func Md5V2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
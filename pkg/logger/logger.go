package logger

import (
	"fmt"
	"log"
)

// LogError 当存在错误时记录日志
func LogError(err error)  {
	if err != nil {
		log.Println(err)
	}
}

//系统错误日志，直接输出
func SystemError(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}

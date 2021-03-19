package config

import "trail_didi_3/pkg/config"

func init() {
	config.Add("redis", config.StrMap{
		"host": config.Env("REDIS_HOST", "127.0.0.1:6379"),
		"auth": config.Env("REDIS_AUTH", "123456"),
		"port": config.Env("REDIS_PORT", "6379"),
		"db":config.Env("REDIS_DB",0),
	})

}

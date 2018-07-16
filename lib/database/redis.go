package database

import (
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"github.com/weeq/meowth/lib"
	"fmt"
	"time"
	"github.com/kataras/iris"
)

func RedisInit() *redis.Database {
	RedisHost := lib.Getenv("REDIS_HOST", "127.0.0.1")
	RedisPort := lib.Getenv("REDIS_PORT", "6379")
	cfg := &service.Config{
		Network:     "tcp",
		Addr:        fmt.Sprintf("%v:%v", RedisHost, RedisPort),
		Password:    lib.Getenv("REDIS_PASSWORD", ""),
		Database:    lib.Getenv("REDIS_DB", ""),
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: time.Duration(5) * time.Minute,
		Prefix:      "",
	}

	db := redis.New(*cfg)

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() { db.Close() })
	return db
}

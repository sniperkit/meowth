package config

import (
	"time"

	"github.com/weeq/meowth/lib"
)

type SessionsConfig struct {
	Expires time.Duration
	Cookie  string
	Drive   string
}

// Set Default Configuration
func SetSessionsConfig() *SessionsConfig {
	expires := lib.ParseInt(lib.Getenv("SESSION_EXPIRES", "2"))
	return &SessionsConfig{
		Expires: time.Duration(expires) * time.Hour, // time.Minute
		Cookie:  lib.Getenv("APP_NAME", "ads"),
		Drive:   lib.Getenv("SESSION_DRIVER", "badger"), // redis || boltdb || badger
	}
}

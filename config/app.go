package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/weeq/meowth/lib"
)

type AppConfig struct {
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
	AppUrl       string
	Port         string
	Env          string
	Locale       string
}

// Set Default Configuration
func SetAppConfig() *AppConfig {
	env := godotenv.Load()

	if env != nil {
		log.Fatal("Error loading .env file")
	}

	return &AppConfig{
		AppName:      lib.Getenv("APP_NAME", "Ads"),
		AppUrl:       lib.Getenv("APP_URL", "http://localhost"),
		AppSpawnDate: time.Now(),
		Port:         lib.Getenv("PORT", "8888"),
		Env:          lib.Getenv("APP_ENV", "product"),
		AppOwner:     lib.Getenv("APP_OWNER", "ads@ads.co"),
		Locale:       "tr-TR",
	}
}

package bootstrap

import (
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/weeq/meowth/lib"

	"github.com/gorilla/securecookie"
	"github.com/weeq/meowth/config"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/badger"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"github.com/weeq/meowth/lib/database"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
	"fmt"
)

var (
	LocalePath   = path.Join("./", "resources", "locales")
	DatabasePath = path.Join("./", "database")
)

type Configurator func(*Bootstrapper)

type Config struct {
	App      *config.AppConfig
	View     *config.ViewConfig
	Sessions *config.SessionsConfig
	Redis    *service.Config
}

type Bootstrapper struct {
	*iris.Application
	Config   *Config
	Sessions *sessions.Sessions
	Db       *gorm.DB
	MGO      *mgo.Database
}

// Load App config
func LoadAppConfig() *config.AppConfig {
	return config.SetAppConfig()
}

// Load View config
func LoadViewConfig() *config.ViewConfig {
	return config.SetViewConfig()
}

// Load View config
func LoadSessionConfig() *config.SessionsConfig {
	return config.SetSessionsConfig()
}

// Load Redis
func LoadRedis() *redis.Database {
	return database.RedisInit()
}

// Load Badger
func LoadBadger() *badger.Database {
	return database.BadgerDbInit()
}

// Load BoltDb
func LoadBoltDb() *boltdb.Database {
	return database.BoltDbInit()
}

func Boot(cfgs ...Configurator) *Bootstrapper {
	conf := Config{
		App:      LoadAppConfig(),
		View:     LoadViewConfig(),
		Sessions: LoadSessionConfig(),
	}

	boot := &Bootstrapper{
		Config:      &conf,
		Application: iris.New(),
		Db:          config.DatabaseInit(),
		MGO:         database.MgoInit(),
	}

	for _, cfg := range cfgs {
		cfg(boot)
	}

	return boot
}

func (boot *Bootstrapper) LoadLocales() {

	files, err := ioutil.ReadDir(LocalePath)
	if err != nil {
		log.Fatal(err)
	}

	languages := make(map[string]string)
	for _, f := range files {
		ext := filepath.Ext(f.Name())
		name := strings.Replace(f.Name(), string(ext), "", -1)
		dir := path.Join(LocalePath, f.Name())
		languages[name] = fmt.Sprintf("./%s", dir)
	}

	globalLocale := i18n.New(i18n.Config{
		Default:      boot.Config.App.Locale,
		Languages:    languages,
	})

	boot.Use(globalLocale)
}

func (boot *Bootstrapper) SetupViews() {
	viewsDir := boot.Config.View.Path
	layout := boot.Config.View.Layout
	view := iris.HTML(viewsDir, boot.Config.View.Ext).Layout(layout).Reload(true)
	boot.RegisterView(view)
}

func (boot *Bootstrapper) setCookieSession() *sessions.Sessions {
	hashKey := lib.RandomHmac512(32)
	blockKey := lib.EncodeSha512(hashKey)
	return sessions.New(sessions.Config{
		Cookie:   boot.Config.Sessions.Cookie,
		Expires:  boot.Config.Sessions.Expires,
		Encoding: securecookie.New([]byte(hashKey), []byte(blockKey)),
	})
}

func (boot *Bootstrapper) setDBSession() *sessions.Sessions {
	session := sessions.New(sessions.Config{
		Cookie:  boot.Config.Sessions.Cookie,
		Expires: boot.Config.Sessions.Expires}, // <=0 means unlimited life. Defaults to 0.
	)
	return session
}

func (boot *Bootstrapper) SetupSessions() {
	session := boot.setCookieSession()
	switch boot.Config.Sessions.Drive {
	case "cookie":
		session = boot.setCookieSession()
		break
	case "redis":
		//session = boot.setDBSession()
		//session.UseDatabase(boot.Redis)
		break
	case "boltdb":
		session = boot.setDBSession()
		session.UseDatabase(LoadBoltDb())
		break
	case "badger":
		session = boot.setDBSession()
		session.UseDatabase(LoadBadger())
		break
	}

	boot.Sessions = session
}

func (boot *Bootstrapper) Configure(cfgs ...Configurator) {
	for _, cfg := range cfgs {
		cfg(boot)
	}
}

func (boot *Bootstrapper) NewApp() *Bootstrapper {
	boot.LoadLocales()
	boot.SetupSessions()
	boot.SetupViews()
	// middleware
	boot.Use(recover.New())
	boot.Use(logger.New())

	return boot
}

func (boot *Bootstrapper) Listen(cfgs ...iris.Configurator) {
	port := boot.Config.App.Port
	boot.Run(iris.Addr(port), cfgs...)
}

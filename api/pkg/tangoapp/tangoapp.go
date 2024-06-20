package tangoapp

import (
	"fmt"
	"log"

	"github.com/k23dev/dbman"
	"github.com/k23dev/tango/pkg/tango_helpers"
	"github.com/k23dev/tango/pkg/tango_log"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TangoApp struct {
	Config *AppConfig
	Log    *echo.Logger
	Server *echo.Echo
	DB     *dbman.DBMan
	DBAuth *gorm.DB
}

const appconfigFile = "app.toml"
const dbConfigFile = "db.toml"

func NewTangoApp(configPath string) *TangoApp {

	tango_log.Print("Loading APP config (" + configPath + appconfigFile + ")...")
	var appconfig AppConfig
	// var dbman dbman.DBMan
	tango_helpers.ReadAndParseToml(configPath+appconfigFile, &appconfig)
	// set default app url
	appconfig.Url = fmt.Sprintf("%s:%d", appconfig.ServerHost, appconfig.ServerPort)

	tapp := &TangoApp{
		Config: &appconfig,
	}
	tango_log.PrintOk("App config loaded")
	// tapp.InitDBAuth("")
	tango_log.Print("Server config initialized")
	tapp.InitServer()

	dbConfigPath := "./config/" + dbConfigFile
	tango_log.Print("Loading DB config (" + dbConfigPath + ")...")
	tapp.InitDBMan(dbConfigPath)

	return tapp
}

func (tapp *TangoApp) InitDBMan(dbConfigPath string) {
	tapp.DB = dbman.New()
	tapp.DB.LoadConfigToml(dbConfigPath)
}

func (tapp *TangoApp) InitServer() {
	server := echo.New()
	tapp.Server = server
}

func (tapp *TangoApp) InitDBAuth(dbConfigPath string) {
	// tapp.DBAuth = dbman.New()
}

func (tapp *TangoApp) PrintAppInfo() {
	log.Printf("Starting app: %s (v%s)\n", tapp.Config.Name, tapp.Config.Version)
}

func (tapp *TangoApp) GetAppUrl() string {
	return fmt.Sprintf("%s:%d", tapp.Config.ServerHost, tapp.Config.ServerPort)
}

func (tapp *TangoApp) GetPortAsStr() string {
	return fmt.Sprintf("%d", tapp.Config.ServerPort)
}

func (tapp *TangoApp) GetTitleAndVersion() string {
	return tapp.Config.Name + " (V." + tapp.Config.Version + ")"
}

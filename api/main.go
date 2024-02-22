package main

import (
	"net/http"

	"github.com/k23dev/go4it"
	"github.com/k23dev/tango/app"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/k23dev/tango/pkg/webcore_features"
	"github.com/labstack/echo/v4"
)

func main() {

	app_config := go4it.NewApp("./config/appconfig")

	tapp := webcore.TangoApp{
		App:    &app_config,
		Server: echo.New(),
	}

	// Database connections
	app_config.Connect2Db("local")
	app_config.DB.SetPrimaryDB(0)

	tapp.PrintAppInfo()

	// Middleware
	webcore.MiddlewareSetup(&tapp)

	//  Routes

	if tapp.App.Config.App_setup_enabled && tapp.App.Config.App_debug_mode {
		webcore_features.SetupRoutes(&tapp)
	}

	webcore.SetupStaticRoutes(tapp.Server)

	app.AppSetup(&tapp)

	// open app in default browser
	// interact.OpenInBrowser("http://" + tapp.GetAppUrl())

	// Start server
	if tapp.App.Config.App_debug_mode {
		tapp.Server.Logger.Fatal(tapp.Server.Start(":" + tapp.GetPortAsStr()))
	} else {
		// tapp.Server.Start(":" + tapp.GetPortAsStr())
		if err := tapp.Server.Start(":8080"); err != http.ErrServerClosed {
			tapp.Server.Logger.Fatal(err)
		}
	}

}

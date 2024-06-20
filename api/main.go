package main

import (
	"log"

	"github.com/k23dev/tango/app"
	"github.com/k23dev/tango/pkg/tango_helpers"
	"github.com/k23dev/tango/pkg/tango_log"
	"github.com/k23dev/tango/pkg/tango_middlewares"
	"github.com/k23dev/tango/pkg/tango_routes"
	"github.com/k23dev/tango/pkg/tangoapp"
)

const configPath = "./config/"

func init() {
	tango_log.Print("Starting up")
}

func main() {
	tapp := tangoapp.NewTangoApp(configPath)
	err := tapp.DB.Connect("local")
	if err != nil {
		log.Fatal(err)
	}

	tapp.PrintAppInfo()

	// Middleware
	tango_middlewares.Setup(tapp)

	//  Tango Routes
	if tapp.Config.SetupEnabled && tapp.Config.NotInProduction {
		tango_routes.SetupRoutes(tapp)
	}

	tango_routes.SetupStaticRoutes(tapp)

	// App routes
	app.AppSetup(tapp)

	// open app in default browser
	if tapp.Config.OpenInBrowser {
		tango_helpers.OpenInBrowser("http://" + tapp.GetAppUrl())
	}

	// Start server
	tapp.Server.Logger.Fatal(tapp.Server.Start(":" + tapp.GetPortAsStr()))

}

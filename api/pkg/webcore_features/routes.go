package webcore_features

import (
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(tapp *webcore.TangoApp) {

	// setup
	setup := tapp.Server.Group("/setup")

	setup.GET("/", func(c echo.Context) error {
		return Setup(c, tapp)
	})

	//status
	setup.GET("/status", func(c echo.Context) error {
		return Status(c)
	})

	// seeder
	if tapp.App.Config.App_setup_enabled {
		setup.GET("/seed", func(c echo.Context) error {
			return Seed(c, tapp.App.DB.Primary)
		})
		setup.GET("/seed/:table_name", func(c echo.Context) error {
			return Seed(c, tapp.App.DB.Primary)
		})
	}

}

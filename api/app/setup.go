package app

import (
	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/app/routes"
	"github.com/k23dev/tango/pkg/tango_auth"
	"github.com/k23dev/tango/pkg/webcore"
)

func AppSetup(tapp *webcore.TangoApp) {

	// features routes
	routes.SetupAppRoutes(tapp)

	// migrate tables
	if tapp.App.Config.App_debug_mode {
		tapp.App.DB.Primary.AutoMigrate(&models.Category{})
		// migrate auth
		tapp.App.DB.Auth.AutoMigrate(&tango_auth.User{}, &tango_auth.Auth{})
	}

}

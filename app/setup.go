package app

import (
	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/app/routes"
	"github.com/k23dev/tango/pkg/webcore"
)

func AppSetup(tapp *webcore.TangoApp) {

	// features routes
	routes.SetupAppRoutes(tapp)

	if tapp.App.Config.App_debug_mode {
		tapp.App.DB.Primary.AutoMigrate(&models.Category{})
	}

}

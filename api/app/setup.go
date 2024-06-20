package app

import (
	"github.com/k23dev/tango/app/routes"
	"github.com/k23dev/tango/pkg/tangoapp"
)

func AppSetup(tapp *tangoapp.TangoApp) {

	routes.SetupAppRoutes(tapp)
}

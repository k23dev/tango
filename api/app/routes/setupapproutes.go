package routes

import (
	"tango_pkg/tangoapp"
)

func SetupAppRoutes(tapp *tangoapp.TangoApp) {
	rootPath := tapp.Server.Group("api")

	// Auth WIP
	// tango_auth.AuthRoutes(tapp, rootPath)
	// tango_auth.UsersRoutes(tapp, rootPath)

	// categories
	categoriesRoutes(tapp, rootPath)
}

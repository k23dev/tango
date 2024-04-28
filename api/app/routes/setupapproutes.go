package routes

import (
	"github.com/k23dev/tango/pkg/tango_auth"
	"github.com/k23dev/tango/pkg/webcore"
)

func SetupAppRoutes(tapp *webcore.TangoApp) {
	// rootPath := tapp.Server.Group("/api")
	rootPath := tapp.Server.Group("")

	// Auth
	tango_auth.AuthRoutes(tapp, rootPath)
	tango_auth.UsersRoutes(tapp, rootPath)
	// categories
	categoriesRoutes(tapp, rootPath)
}

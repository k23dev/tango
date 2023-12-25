package routes

import (
	"github.com/k23dev/tango/pkg/webcore"
)

func SetupAppRoutes(tapp *webcore.TangoApp) {
	// rootPath := tapp.Server.Group("/api")
	rootPath := tapp.Server.Group("")

	IndexRoutes(tapp, rootPath)
	// categories
	categoriesRoutes(tapp, rootPath)
}

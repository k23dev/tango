package tango_routes

import (
	"github.com/k23dev/tango/pkg/tangoapp"
)

func SetupStaticRoutes(tapp *tangoapp.TangoApp) {

	tapp.Server.Static("/", tapp.Config.PublicPath)
	tapp.Server.Static("/public", tapp.Config.PublicPath)
	tapp.Server.Static("/assets", tapp.Config.PublicAssetsPath)
	tapp.Server.Static("/assets/js", tapp.Config.PublicAssetsPath+"/js")
	tapp.Server.Static("/assets/css", tapp.Config.PublicAssetsPath+"/css")
	tapp.Server.Static("/images", tapp.Config.PublicAssetsPath+"/images")

}

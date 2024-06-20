package tango_middlewares

import (
	"github.com/k23dev/tango/pkg/tangoapp"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(tapp *tangoapp.TangoApp) {

	// CORS
	tapp.Server.Use(middleware.CORS())

	//  Recover from error
	tapp.Server.Use(middleware.Recover())

	tapp.Server.Use(middleware.Logger())

}

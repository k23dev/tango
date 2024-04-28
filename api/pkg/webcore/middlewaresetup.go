package webcore

import (
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareSetup(tapp *TangoApp) {

	// CORS
	tapp.Server.Use(middleware.CORS())

	//  Recover from error
	tapp.Server.Use(middleware.Recover())

	tapp.Server.Use(middleware.Logger())

}

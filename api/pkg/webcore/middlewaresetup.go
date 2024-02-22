package webcore

import (
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareSetup(tapp *TangoApp) {

	// tapp.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{tapp.App.Config.App_CORS_origins},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	tapp.Server.Use(middleware.CORS())

	//  Recover from error
	tapp.Server.Use(middleware.Recover())

	tapp.Server.Use(middleware.Logger())

}

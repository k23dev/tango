package webcore

import (
	"github.com/labstack/echo/v4"
)

const publicPath = "./public"
const assetsPath = publicPath + "/assets"

func SetupStaticRoutes(server *echo.Echo) {

	server.Static("/public", publicPath)
	server.Static("/assets", assetsPath)
	server.Static("/assets/js", assetsPath+"/js")
	server.Static("/assets/css", assetsPath+"/css")
	server.Static("/images", publicPath+"/images")

}

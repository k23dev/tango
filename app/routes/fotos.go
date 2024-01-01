
package routes

import (
	"github.com/k23dev/tango/app/features"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func fotosRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	fotos := rootPath.Group("/fotos/")

	fotos.GET("", func(c echo.Context) error {
		return features.FindAllFotos(c, tapp)
	})

	fotos.GET(":id", func(c echo.Context) error {
		return features.FindOneFoto(c, tapp)
	})

	fotos.GET("new", func(c echo.Context) error {
		return features.ShowFormFoto(c, tapp, true)
	})

	fotos.GET("edit/:id", func(c echo.Context) error {
		return features.ShowFormFoto(c, tapp, false)
	})

	fotos.POST("create", func(c echo.Context) error {
		return features.CreateFoto(c, tapp)
	})

	fotos.POST("update/:id", func(c echo.Context) error {
		return features.UpdateFoto(c, tapp)
	})

	fotos.GET("delete/:id", func(c echo.Context) error {
		return features.DeleteFoto(c, tapp)
	})
}
	
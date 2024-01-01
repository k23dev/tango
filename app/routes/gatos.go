
package routes

import (
	"github.com/k23dev/tango/app/features"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func gatosRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	gatos := rootPath.Group("/gatos/")

	gatos.GET("", func(c echo.Context) error {
		return features.FindAllGatos(c, tapp)
	})

	gatos.GET(":id", func(c echo.Context) error {
		return features.FindOneGato(c, tapp)
	})

	gatos.GET("new", func(c echo.Context) error {
		return features.ShowFormGato(c, tapp, true)
	})

	gatos.GET("edit/:id", func(c echo.Context) error {
		return features.ShowFormGato(c, tapp, false)
	})

	gatos.POST("create", func(c echo.Context) error {
		return features.CreateGato(c, tapp)
	})

	gatos.POST("update/:id", func(c echo.Context) error {
		return features.UpdateGato(c, tapp)
	})

	gatos.GET("delete/:id", func(c echo.Context) error {
		return features.DeleteGato(c, tapp)
	})
}
	
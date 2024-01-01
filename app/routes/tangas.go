
package routes

import (
	"github.com/k23dev/tango/app/features"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func tangasRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	tangas := rootPath.Group("/tangas/")

	tangas.GET("", func(c echo.Context) error {
		return features.FindAllTangas(c, tapp)
	})

	tangas.GET(":id", func(c echo.Context) error {
		return features.FindOneTanga(c, tapp)
	})

	tangas.GET("new", func(c echo.Context) error {
		return features.ShowFormTanga(c, tapp, true)
	})

	tangas.GET("edit/:id", func(c echo.Context) error {
		return features.ShowFormTanga(c, tapp, false)
	})

	tangas.POST("create", func(c echo.Context) error {
		return features.CreateTanga(c, tapp)
	})

	tangas.POST("update/:id", func(c echo.Context) error {
		return features.UpdateTanga(c, tapp)
	})

	tangas.GET("delete/:id", func(c echo.Context) error {
		return features.DeleteTanga(c, tapp)
	})
}
	
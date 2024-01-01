
package routes

import (
	"github.com/k23dev/tango/app/features"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func padronsRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	padrons := rootPath.Group("/padrons/")

	padrons.GET("", func(c echo.Context) error {
		return features.FindAllPadrons(c, tapp)
	})

	padrons.GET(":id", func(c echo.Context) error {
		return features.FindOnePadron(c, tapp)
	})

	padrons.GET("new", func(c echo.Context) error {
		return features.ShowFormPadron(c, tapp, true)
	})

	padrons.GET("edit/:id", func(c echo.Context) error {
		return features.ShowFormPadron(c, tapp, false)
	})

	padrons.POST("create", func(c echo.Context) error {
		return features.CreatePadron(c, tapp)
	})

	padrons.POST("update/:id", func(c echo.Context) error {
		return features.UpdatePadron(c, tapp)
	})

	padrons.GET("delete/:id", func(c echo.Context) error {
		return features.DeletePadron(c, tapp)
	})
}
	
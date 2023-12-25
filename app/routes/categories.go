package routes

import (
	"github.com/k23dev/tango/app/features"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func categoriesRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	categories := rootPath.Group("/categories/")

	categories.GET("", func(c echo.Context) error {
		return features.FindAllCategories(c, tapp)
	})

	categories.GET(":id", func(c echo.Context) error {
		return features.FindOneCategory(c, tapp)
	})

	categories.GET("new", func(c echo.Context) error {
		return features.ShowFormCategory(c, tapp, true)
	})

	categories.GET("edit/:id", func(c echo.Context) error {
		return features.ShowFormCategory(c, tapp, false)
	})

	categories.POST("create", func(c echo.Context) error {
		return features.CreateCategory(c, tapp)
	})

	categories.POST("update/:id", func(c echo.Context) error {
		return features.UpdateCategory(c, tapp)
	})

	categories.GET("delete/:id", func(c echo.Context) error {
		return features.DeleteCategory(c, tapp)
	})
}

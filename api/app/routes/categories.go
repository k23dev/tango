package routes

import (
	"github.com/k23dev/tango/app/features"
	"github.com/k23dev/tango/pkg/tangoapp"
	"github.com/labstack/echo/v4"
)

func categoriesRoutes(tapp *tangoapp.TangoApp, rootPath *echo.Group) {
	categories := rootPath.Group("/categories/")

	categories.GET(":id", func(ctx echo.Context) error {
		return features.FindOneCategory(ctx, tapp)
	})

	categories.GET("", func(ctx echo.Context) error {
		return features.FindAllCategories(ctx, tapp)
	})

	categories.POST("create", func(ctx echo.Context) error {
		return features.CreateCategory(ctx, tapp)
	})

	categories.PUT("update/:id", func(ctx echo.Context) error {
		return features.UpdateCategory(ctx, tapp)
	})

	categories.DELETE("delete/:id", func(ctx echo.Context) error {
		return features.DeleteCategory(ctx, tapp)
	})
}

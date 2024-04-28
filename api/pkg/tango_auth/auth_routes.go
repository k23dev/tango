package tango_auth

import (
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	users := rootPath.Group("/auth/")

	users.POST("login", func(ctx echo.Context) error {
		return AuthLogin(ctx, tapp)
	})

	users.GET("logout", func(ctx echo.Context) error {
		return AuthLogout(ctx, tapp)
	})
}

package templates

func (t *Templates) RouteAPI() string {
	t.setReplacements()

	template := `
package routes

import (
	"github.com/k23dev/tango/app/features"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func $PL$Routes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	$PL$ := rootPath.Group("/$PL$/")
	feat:=features.New$PC$Feature(tapp)

	$PL$.GET(":id", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.FindOne()
	})

	$PL$.GET("", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.FindAll()
	})

	$PL$.POST("create", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.Create()
	})

	$PL$.PUT("update/:id", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.Update()
	})

	$PL$.DELETE("delete/:id", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.Delete()
	})
}
	`
	return t.Replacements.Replace(template)
}

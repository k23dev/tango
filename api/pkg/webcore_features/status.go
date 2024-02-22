package webcore_features

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Status(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

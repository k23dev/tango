package tango_routes

import (
	"fmt"
	"net/http"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/pkg/tangoapp"
	"github.com/labstack/echo/v4"
)

func Setup(c echo.Context, tapp *tangoapp.TangoApp) error {
	automigrateModels(tapp)
	return c.String(http.StatusOK, "Setup enabled. Models Migrated.")
}

func SetupOnStartup(tapp *tangoapp.TangoApp) {
	fmt.Println("\nDatabase automigration...")
	automigrateModels(tapp)
}

func automigrateModels(tapp *tangoapp.TangoApp) {
	tapp.DB.Primary.AutoMigrate(&models.Category{})
}

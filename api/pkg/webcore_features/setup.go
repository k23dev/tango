package webcore_features

import (
	"fmt"
	"net/http"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func Setup(c echo.Context, tapp *webcore.TangoApp) error {
	automigrateModels(tapp)
	return c.String(http.StatusOK, "Setup enabled. Models Migrated.")
}

func SetupOnStartup(tapp *webcore.TangoApp) {
	fmt.Println("\nDatabase automigration...")
	automigrateModels(tapp)
}

func automigrateModels(tapp *webcore.TangoApp) {
	tapp.App.DB.Primary.AutoMigrate(&models.Category{})
}

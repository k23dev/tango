package webcore_features

import (
	"log"
	"net/http"

	"github.com/k23dev/go4it/interact"
	"github.com/k23dev/tango/app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const seedDir = "./seeds/"

func Seed(c echo.Context, db *gorm.DB) error {
	seedCategories(db)
	return c.JSON(http.StatusOK, "OK")
}

func seedCategories(db *gorm.DB) {
	cat_list := []models.Category{}
	interact.ReadAndParseJson(seedDir+"categories.json", &cat_list)

	db.Save(&cat_list)
	log.Println("Categories seeded")
}

package migrations

import (
	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/pkg/tangoapp"
)

func Migrate() {

}

func Setup(tapp *tangoapp.TangoApp) {

	// is DebugMode == true
	// migrate tables
	if tapp.Config.NotInProduction {
		tapp.DB.Primary.AutoMigrate(&models.Category{})
		// migrateAuth(tapp.DBAuth)
	}

}

// func migrateAuth(dbAuth *gorm.DB) {
// 	// migrate auth
// 	dbAuth.AutoMigrate(&tango_auth.User{}, &tango_auth.Auth{})
// }

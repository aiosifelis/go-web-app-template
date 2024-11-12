package migrations

import (
	"github.com/aiosifelis/go-web-app-template/src/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var createUserTable = &gormigrate.Migration{
	ID: "202411121426_create_users_table",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&models.User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}

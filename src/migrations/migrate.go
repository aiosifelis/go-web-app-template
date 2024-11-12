package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		createUserTable,
	})

	if err := m.Migrate(); err != nil {
		log.Fatal("Could not migrate:", err)
	}
	log.Println("Migration did run successfully")
}

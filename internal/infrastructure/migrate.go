package infrastructure

import (
	"log"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.UserProfile{},
		&entities.Problem{},
	)

	if err != nil {
		log.Printf("Error migrating database: %v\n", err)
		return err
	}

	log.Println("✅ Database migrated successfully")
	return nil
}

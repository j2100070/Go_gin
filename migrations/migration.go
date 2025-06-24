package migrations

import (
	"Go_gin/infra"
	"Go_gin/models"
)

func Migrate() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Failed to migrate database")
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}

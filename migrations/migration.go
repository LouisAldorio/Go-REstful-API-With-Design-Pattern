package migrations

import (
	"myapp/config"
	"myapp/models"
)

func Migrate() {
	db := config.GetDB()
	db.AutoMigrate(&models.Todo{}, &models.User{})
}
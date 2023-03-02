package database

import "task-5-vix-btpns-SuburFirmansyah/models"

func MigrateDb() {
	db := ConnDb()
	db.AutoMigrate(&models.User{}, &models.Photo{})
}

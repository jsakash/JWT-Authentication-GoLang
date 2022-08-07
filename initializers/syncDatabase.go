package initializers

import "github.com/jsakash/jwt-check/models"

func SyncDatabase() {

	DB.AutoMigrate(&models.User{})

}

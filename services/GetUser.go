package services

import (
	"github.com/TestingGorm/models"
)

//GetUser returns user from database
func GetUser(u *models.User) {
	database.db.Set("gorm:auto_preload", true).Where("name = ?", u.Name).Find(u)
	//database.db.Preload("Membership").Preload("Products").Preload("Creditcards", "deleted_at IS NULL").Preload("AccessLevels").Find(u)
}

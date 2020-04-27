package services

import "github.com/TestingGorm/models"

//Query queries data from DB
func Query(table string, u *[]models.User) {
	database.db.Set("gorm:auto_preload", true).Find(u)
	//database.db.Table(table).Find(u)
}

package services

import "github.com/TestingGorm/models"

//GetUser returns user from database
func GetUser(u *models.User) {
	database.db.Where("name = ?", u.Name).Find(u)
	RelateUser(u)
}

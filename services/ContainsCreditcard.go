package services

import "github.com/TestingGorm/models"

//ContainsCreditcard checks if a creditcard is already attached to an user
func ContainsCreditcard(u *models.User, num string) bool {
	var result models.Creditcard
	database.db.Model(&result).Where("number = ? AND user_id = ?", num, u.ID).Find(&result)
	if result.ID == 0 {
		return false
	}
	return true
}

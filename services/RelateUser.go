package services

import (
	"github.com/TestingGorm/models"
)

//RelateUser gets the information related to user from different databases
func RelateUser(u *models.User) {
	database.db.Model(u).Related(&((*u).Membership))
	database.db.Model(u).Related(&((*u).AccessLevels))
	database.db.Model(u).Related(&((*u).Creditcards))
	//log.Print(database.db.Model(&u).Related(&u.Membership).Error)
}

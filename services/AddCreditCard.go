package services

import (
	"github.com/TestingGorm/models"
)

//AddCreditCard appeds creditcard info into user
func AddCreditCard(u *models.User, nro string) {
	database.db.Model(u).Association("Creditcards").Append(models.Creditcard{Number: nro})
}

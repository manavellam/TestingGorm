package middleware

import (
	"github/TestingGorm/models"
	"github/TestingGorm/services"
)

//IsUser checks if a pair of U&P is valid
func IsUser(u *models.User) bool {
	return services.ContainsUser(u)
}

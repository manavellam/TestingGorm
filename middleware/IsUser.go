package middleware

import (
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
)

//IsUser checks if a pair of U&P is valid
func IsUser(u *models.User) bool {
	return services.ContainsUser(u)
}

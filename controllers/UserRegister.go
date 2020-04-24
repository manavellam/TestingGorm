package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
	"github.com/TestingGorm/util"

	"github.com/gin-gonic/gin"
)

//UserRegister responds to a POST request at /register, and adds an user to the database if not exists.
func UserRegister(c *gin.Context) {
	var (
		loaduser []models.User
	)
	errs := make(map[string]string)
	c.Bind(&loaduser)
	for i, u := range loaduser {
		if len(strings.TrimSpace(u.Name)) != 0 && len(strings.TrimSpace(u.Password)) != 0 {
			u.Password = util.HashString(u.Password)
			if !services.ContainsUser(&u) {
				log.Print(u)
				services.Insert("users", &u)
			} else {
				errs[u.Name] = "Already exists."
			}
		} else {
			if u.Name == "" {
				errs[fmt.Sprintf("User #%v ", i)] = "Missing Info."
			} else {
				errs[u.Name] = "Missing Info."
			}
		}
	}
	log.Print("ERRS:", errs)
	if len(errs) > 0 {
		c.Writer.Write([]byte(fmt.Sprintf("Errors occur during insertion of data. Following users were not added: %v", errs)))
	} else {
		c.Writer.WriteHeader(200)
		c.Writer.WriteString("User(s) succesfuly registered into DB.")
	}
}

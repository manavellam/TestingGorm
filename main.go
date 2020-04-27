package main

import (
	"github.com/TestingGorm/app"
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/routes"
	"github.com/TestingGorm/services"
)

func main() {
	var conf models.Config
	app.ReadConfig(&conf)

	var u models.User
	u.Name = "Silvia"
	services.GetUser(&u)
	router := routes.Routes()
	router.Run(conf.ServerPort)
}

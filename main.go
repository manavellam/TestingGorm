package main

import (
	"github.com/TestingGorm/app"
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/routes"
)

func main() {
	var conf models.Config
	app.ReadConfig(&conf)

	router := routes.Routes()
	router.Run(conf.ServerPort)
}

package main

import (
	"github.com/TestingGorm/routes"
)

func main() {
	router := routes.Routes()
	router.Run(":8080")
}

package main

import (
	"github/TestingGorm/routes"
)

func main() {
	router := routes.Routes()
	router.Run(":8080")
}

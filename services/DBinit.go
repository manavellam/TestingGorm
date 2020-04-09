package services

import (
	"log"

	"github.com/jinzhu/gorm"
	//Mysqldriver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB Encapsulates database
type DB struct {
	db *gorm.DB
}

var database DB

func init() {
	var err error
	database.db, err = gorm.Open("mysql", "root:root@/Test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Print("ERROR")
		log.Println(err)
	}

}

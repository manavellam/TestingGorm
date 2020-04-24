package services

import (
	"fmt"
	"log"
	"strconv"

	"github.com/TestingGorm/app"
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/util"
	"github.com/jinzhu/gorm"

	//Mysqldriver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB Encapsulates database
type DB struct {
	db *gorm.DB
}

var database DB

//DBinit is...
func init() {
	var err error
	var config models.Config
	var user models.User

	app.ReadConfig(&config)
	mySQLStmt := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", config.MySQLUser, config.MySQLPass, config.MySQLDB)
	database.db, err = gorm.Open("mysql", mySQLStmt)
	if err != nil {
		log.Print("ERROR")
		log.Println(err)
	}

	//Creates table for testing purposes
	database.db.AutoMigrate(&user)
	user.Name = config.DBSampleUser
	user.Password = util.HashString(config.DBSamplePass)
	user.AccessLevelsID, _ = strconv.Atoi(config.DBSampleAccess)
	user.MembershipID, _ = strconv.Atoi(config.DBSampleMember)

	//database.db.Where("id = ?", &user.MembershipID).Find(&user.Membership)
	//database.db.Where("id = ?", &user.AccessLevelsID).Find(&user.AccessLevels)

	if !ContainsUser(&user) {
		Insert("users", &user)
	}

}

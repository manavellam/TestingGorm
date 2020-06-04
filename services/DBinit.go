package services

import (
	"fmt"
	"log"
	"strconv"
	"time"

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
	var memb models.Membership
	var cred models.Creditcard
	var prod models.Product
	var access models.AccessLevels

	app.ReadConfig(&config)
	//(docker.for.mac.localhost)

	//This statement works in compose, where address is not localhost but <db-service name>
	mySQLStmt := fmt.Sprintf("%v:%v@(db)/%v?charset=utf8&parseTime=True&loc=Local", config.MySQLUser, config.MySQLPass, config.MySQLDB)

	//This statement works with db in localhost
	//mySQLStmt := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True&loc=Local", config.MySQLUser, config.MySQLPass, config.MySQLDB)
	database.db, err = gorm.Open("mysql", mySQLStmt)
	for err != nil {
		log.Print("Trying to connect...")
		log.Println(err)
		time.Sleep(5 * time.Second)
		database.db, err = gorm.Open("mysql", mySQLStmt)
	}

	//Creates table for testing purposes
	database.db.AutoMigrate(&user)
	database.db.AutoMigrate(&memb)
	database.db.AutoMigrate(&prod)
	database.db.AutoMigrate(&cred)
	database.db.AutoMigrate(&access)
	user.Name = config.DBSampleUser
	user.Password = util.HashString(config.DBSamplePass)
	user.AccessLevelsID, _ = strconv.Atoi(config.DBSampleAccess)
	user.MembershipID, _ = strconv.Atoi(config.DBSampleMember)

	if !ContainsUser(&user) {
		Insert("users", &user)
	}

}

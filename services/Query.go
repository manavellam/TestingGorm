package services

import "log"

//Query queries data from DB
func Query(table string, u interface{}) {
	log.Print(u)
	database.db.Table(table).Find(u)
}

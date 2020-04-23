package services

import "log"

//Insert inserts data into an specified table in DB
func Insert(table string, u interface{}) {
	log.Printf("TABLE: >>%v<< , STRUCTURE: >>%+v<<", table, u)
	database.db.Table(table).Create(u)
}

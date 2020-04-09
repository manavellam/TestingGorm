package services

//Query queries data from DB
func Query(table string, u interface{}) {
	database.db.Table(table).Find(u)
}

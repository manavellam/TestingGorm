package services

//Insert inserts data into an specified table in DB
func Insert(table string, u interface{}) {
	database.db.Table(table).Create(u)
}

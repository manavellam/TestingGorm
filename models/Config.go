package models

//Config parses the config.json file
type Config struct {
	ServerPort     string
	MySQLUser      string
	MySQLPass      string
	MySQLDB        string
	DBSampleUser   string
	DBSamplePass   string
	DBSampleAccess string
	DBSampleMember string
}

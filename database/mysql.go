package database

var connection string

func init() {
	connection = "Mysql"
}

func GetDatabaseConnection() string {
	return connection
}

package db_config

import "os"

var DB_DRIVER = "mysql"
var DB_HOST = "localhost"
var DB_PORT = "3306"
var DB_NAME = "tries_rcabs"
var DB_USER = "root"
var DB_PASSWORD = "mysql"

func InitDatabaseConfig() {
	portDbUser := os.Getenv("DB_USER")
	portDbName := os.Getenv("DB_NAME")
	portDbPassword := os.Getenv("DB_PASSWORD")

	if portDbUser != "" {
		DB_USER = portDbUser
	}

	if portDbUser != "" {
		DB_NAME = portDbName
	}

	if portDbUser != "" {
		DB_PASSWORD = portDbPassword
	}
}

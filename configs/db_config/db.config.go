package db_config

import "os"

var DB_HOST = "127.0.0.1" // localhost = 127.0.0.1
var DB_PORT = "3306" 
var DB_NAME = "gin_gonic"
var DB_USER = "root"
var DB_PASSWORD = ""
var DB_DRIVER = "mysql" // support mysql, pgsql

func InitDatabaseConfig() {
	env_DB_HOST := os.Getenv("DB_HOST")
	if env_DB_HOST != "" {
		DB_HOST = env_DB_HOST
	}
	env_DB_PORT := os.Getenv("DB_PORT")
	if env_DB_PORT != "" {
		DB_PORT = env_DB_PORT
	}
	env_DB_NAME := os.Getenv("DB_NAME")
	if env_DB_NAME != "" {
		DB_NAME = env_DB_NAME
	}
	env_DB_USER := os.Getenv("DB_USER")
	if env_DB_USER != "" {
		DB_USER = env_DB_USER
	}
	env_DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if env_DB_PASSWORD != "" {
		DB_PASSWORD = env_DB_PASSWORD
	}
	env_DB_DRIVER := os.Getenv("DB_DRIVER")
	if env_DB_DRIVER != "" {
		DB_DRIVER = env_DB_DRIVER
	}
}

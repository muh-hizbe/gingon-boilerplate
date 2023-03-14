package database

import (
	"fmt"
	"gingon-boilerplate/configs/db_config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var errConnection error
	if db_config.DB_DRIVER == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)
		DB, errConnection = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if db_config.DB_DRIVER == "pgsql" {
		dsnPgSql := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=Asia/Jakarta", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)
		DB, errConnection = gorm.Open(postgres.Open(dsnPgSql), &gorm.Config{})
	}

	if errConnection != nil {
		panic("Can't connect to database")
	} else {
		log.Println("Connected to database")
	}
}

package configs

import (
	"gingon-boilerplate/configs/app_config"
	"gingon-boilerplate/configs/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}

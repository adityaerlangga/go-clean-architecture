package database

import (
	"fmt"
	"go-clean-architecture/app/config"
	"go-clean-architecture/domains"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	db.AutoMigrate(&domains.User{})
	if err != nil {
		panic(err)
	}
	return db
}

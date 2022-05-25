package models

import (
	"crud/helpers"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SetUpDB() {
	config, err := helpers.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	DB_USER := config.DBuser
	DB_PASSWORD := config.DBpassword
	DB_NAME := config.DBname
	DB_HOST := config.DBhost
	DB_PORT := config.DBport
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&User{})
	// defer db.Close() // close the connection when the function returns
	DB = db
}

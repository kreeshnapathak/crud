package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = "thereader"
	DB_NAME     = "mydatabase"
	DB_HOST     = "localhost"
	DB_PORT     = 3306
)

var DB *gorm.DB

func SetUpDB() {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Movie{})
	// defer db.Close() // close the connection when the function returns
	DB = db
}

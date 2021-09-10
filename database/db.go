package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "root"
	password = "password"
	dbname   = "my_db"
	dbhost   = "127.0.0.1"
	dbport   = "3306"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := username + ":" + password + "@tcp" + "(" + dbhost + ":" + dbport + ")/" + dbname + "?parseTime=true&loc=Local"
	// fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}

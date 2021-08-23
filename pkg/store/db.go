package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"uk-towns/pkg/entities"
)

func DataBaseConnect() (db *gorm.DB)  {
	// Open a new connection to our mysql database.
	dsn := "root:root@tcp(127.0.0.1:3306)/foobar?charset=utf8mb4&parseTime=True&loc=Local" // bad andrew!
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection to database failed")
	}
	// run migrations
	err = db.AutoMigrate(&entities.Journey{})
	if err != nil {
		panic("Journey migration failed.")
	}
	err = db.AutoMigrate(&entities.Town{})
	if err != nil {
		panic("Town migration failed.")
	}
	return
}
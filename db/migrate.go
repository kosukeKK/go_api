package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	mds "github.com/catman/go_api/models"
)


func main(){
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// go run db/migrate.go select * from products
	db.AutoMigrate(&mds.Item{})
	db.AutoMigrate(&mds.User{})
}

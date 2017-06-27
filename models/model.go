package models


import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }


type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Product struct {
	gorm.Model
	Code string
	Price uint
}

type Item struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}

type User struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Items []Item `json:"items"`
}

//func models() *DB {
//	db, err := gorm.Open("sqlite3", "test.db")
//	if err != nil {
//		panic("failed to connect database")
//	}
//	defer db.Close()
//	return db
//
//	// Migrate the schema
//	//db.AutoMigrate(&Product{})
//	//// Create
//	//db.Create(&Product{Code: "L1212", Price: 1000})
//	//
//	//// Read
//	//var product Product
//	//db.First(&product, 1) // find product with id 1
//	//db.First(&product, "code = ?", "L1212") // find product with code l1212
//	//
//	//// Update - update product's price to 2000
//	//db.Model(&product).Update("Price", 2000)
//	//
//	//// Delete - delete product
//	//db.Delete(&product)
//}

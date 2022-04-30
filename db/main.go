// hellogo/internal/databases/mysql.go
package db

import (
	"log"
	"playground/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect DB
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/my_first_db?charset=utf8&parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&models.User{}, &models.Roll{})

	return db
}

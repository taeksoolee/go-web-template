package models_test

import (
	"log"
	"playground/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func TestCreate(t *testing.T) {
	db, err := gorm.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/my_first_db?charset=utf8&parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.Roll{})

	db.Create(&models.Roll{
		RollName: "admin",
	})

	var r models.Roll
	db.First(&r)

	db.Create(&models.User{
		FirstName: "lee",
		LastName:  "taeksoo",
		Roll:      r,
	})
}

func TestSelectUser(t *testing.T) {
	db, err := gorm.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/my_first_db?charset=utf8&parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.Roll{})

	var u models.User
	db.Preload("Roll").Find(&u, "id=?", 1)

	log.Default().Println(u)
}

func TestSelectUsers(t *testing.T) {
	db, err := gorm.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/my_first_db?charset=utf8&parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.Roll{})

	var us []models.User
	db.Preload("Roll").Find(&us)

	log.Default().Println(us)
}

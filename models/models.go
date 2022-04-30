package models

import "github.com/jinzhu/gorm"

type Roll struct {
	gorm.Model
	RollName string
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	RollID    int
	Roll      Roll `gorm:"foreignKey:RollID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

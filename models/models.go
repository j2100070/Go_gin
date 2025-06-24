package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Description string
	SoldOut     bool `gorm:"not null;default:false"`
	UserID      *uint
	User        User `gorm:"foreignKey:UserID"`
}

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Email    string `gorm:"not null; unique"`
	Password string `gorm:"not null"`
}

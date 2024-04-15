package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `gorm:"default:user" json:"-" binding:"-`
}

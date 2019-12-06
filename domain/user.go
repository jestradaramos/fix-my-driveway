package domain

import (
	"github.com/jinzhu/gorm"
)

// User will hold the information of people using this applcation
type User struct {
	gorm.Model
	name     string 
	email    string `gorm:"unique"`
	password string 
}

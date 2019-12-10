package domain

import (
	"github.com/jinzhu/gorm"
)

// User will hold the information of people using this applcation
type User struct {
	gorm.Model
	Name     string 
	Email    string `gorm:"unique"`
	Password string 
}

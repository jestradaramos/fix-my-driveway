package domain

import (
	"github.com/jinzhu/gorm"
)

// User will hold the information of people using this applcation
type User struct {
	gorm.Model
	id       string `json:"id,omitempty"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}

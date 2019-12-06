package domain

import (
	"github.com/jinzhu/gorm"
)

// TimeEntry used for holding the information regarding peoples arrival and departure
type TimeEntry struct {
	gorm.Model
	userID    int    `json:"user_id,omitempty"`
	date      string `json:"date,omitempty"`
	entryType string `json:"entry_type,omitempty"`
}

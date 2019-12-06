package repo


import (
	"github.com/jinzhu/gorm"
	// We need the postgres driver
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

// Repository yup
type Repository struct {
	db *gorm.DB
}


// NewRepo will be used to create a database client
func NewRepo() (*Repository, error) {

	// Figure out the options
	db, err := gorm.Open("postgres")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	

	return &Repository{db}, nil
}

package repo
import (
	"errors"
	"github.com/jinzhu/gorm"
	"fix-my-driveway/domain"
)

// InitUserTable makes sure that the table exists on the DB we working on
func (r *Repository) InitUserTable() {
	if !r.db.HasTable(&domain.User{}) {
		r.db.CreateTable(&domain.User{})
	}
}
// AddUser adds a user to our pg db
func (r *Repository) AddUser(user *domain.User) error{
	u := &domain.User{}
	if result := r.db.FirstOrCreate(&u, user); result.Error != nil {
		return result.Error
	}
	return nil
}
// UpdateUser updates the user in our pg db
func (r *Repository) UpdateUser(user *domain.User) error {
	u := &domain.User{}
	if result := r.db.Model(u).Update(user); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByEmail finds a user by their email
func (r *Repository) GetUserByEmail(mail string) (*domain.User, error) {
	user := &domain.User{}
	result := r.db.Model(user).Find(mail)
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return nil, errors.New("User Not Found")
		}
		return nil, result.Error
	}
	return user, nil
}
package interfaces

import (
	"fmt"
	"github.com/adriancarayol/go-event/config/db"
	"github.com/adriancarayol/go-event/pkg/domain/model"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: db.GetDB(),
	}
}

func (r *userRepository) Save(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("error while saving user: %s", err)
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
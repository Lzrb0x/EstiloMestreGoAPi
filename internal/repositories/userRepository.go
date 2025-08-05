package repositories

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"gorm.io/gorm"
)


type UserRepositoryInterface interface {
	AddUser(user *models.User) error
	GetUserById(id uint) (*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) AddUser(user *models.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *UserRepository) GetUserById(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}



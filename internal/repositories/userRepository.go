package repositories

import (
	"github.com/Lzrb0x/estiloMestreGO/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	AddUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetByUserIdentifier(userIdentifier string) (*models.User, error)
	UpdateUser(user *models.User) error
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

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) GetByUserIdentifier(userIdentifier string) (*models.User, error) {
	var user models.User 

	result := r.db.Where("user_identifier = ?", userIdentifier).First(&user)
	if result.Error != nil {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	result := r.db.Save(user)
	return result.Error
}

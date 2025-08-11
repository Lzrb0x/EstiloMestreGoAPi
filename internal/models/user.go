package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model               //id, created_at, updated_at, deleted_at
	UserIdentifier uuid.UUID `gorm:"type:uuid;"` // if not set, will be genereated automatically
	Name           string    `gorm:"not null;type:varchar(255)"`
	Email          string    `gorm:"optional;type:varchar(255);unique"`
	Password       string    `gorm:"optional;type:varchar(512)"`
}

func NewUser(name, email, password string) (*User, error) {
	user := &User{
		Name:           name,
		Email:          email,
		Password:       password,
		UserIdentifier: uuid.New(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {

	var errorMessages []string

	if u.Name == "" {
		errorMessages = append(errorMessages, "Name is required")
	}
	if u.Email == "" {
		errorMessages = append(errorMessages, "Email is required")
	}
	if len(errorMessages) > 0 {
		return fmt.Errorf("%s", strings.Join(errorMessages, ", "))
	}

	return nil
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.Model.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.Model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.Model.UpdatedAt.Format(time.RFC3339),
	}
}
